// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package iscplotterycontract

import (
	"strconv"

	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
)

const ROUND_PERIOD = uint64(86400) // 1 day
const MAX_DIGIT = int(6)           // max random digit
const MAX_NUMBER = uint64(999999)  // the max random number

const MATCH_1_PRIZE_PERCENT = uint64(2)
const MATCH_2_PRIZE_PERCENT = uint64(4)
const MATCH_3_PRIZE_PERCENT = uint64(6)
const MATCH_4_PRIZE_PERCENT = uint64(12)
const MATCH_5_PRIZE_PERCENT = uint64(25)
const MATCH_6_PRIZE_PERCENT = uint64(50)

const FIXED_FEE = uint64(1) // the fee per buy to avoid spammer

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	if f.Params.Owner().Exists() {
		f.State.Owner().SetValue(f.Params.Owner().Value())
		return
	}
	f.State.Owner().SetValue(ctx.RequestSender())
}

func funcSetOwner(ctx wasmlib.ScFuncContext, f *SetOwnerContext) {
	f.State.Owner().SetValue(f.Params.Owner().Value())
}

func viewGetOwner(ctx wasmlib.ScViewContext, f *GetOwnerContext) {
	f.Results.Owner().SetValue(f.State.Owner().Value())
}

func funcCreateTicket(ctx wasmlib.ScFuncContext, f *CreateTicketContext) {
	caller := ctx.Caller()

	// Create ScBalances proxy to the allowance balances for this request.
	// Note that ScBalances wraps an ScImmutableMap of token color/amount combinations
	incoming := ctx.Allowance()
	incomingAmount := incoming.BaseTokens()

	// minus amount by fixed fee
	incomingAmount -= FIXED_FEE
	if incomingAmount <= 0 {
		ctx.Log("invalid incoming amount")
		return
	}

	// get params
	buyNumber := f.Params.Number().Value()
	if len(buyNumber) != MAX_DIGIT {
		ctx.Log("invalid number")
		return
	}

	// get state
	tickets := f.State.Tickets()
	ticketsLen := tickets.Length()
	currentRoundIdx := f.State.CurrentRoundIdx().Value()
	rounds := f.State.Rounds()
	currentRound := rounds.GetRound(currentRoundIdx).Value()

	ticket := &Ticket{
		Idx:         ticketsLen,
		Id:          ctx.RequestID(),
		Buyer:       caller,
		Amount:      incomingAmount,
		Number:      buyNumber,
		RoundIdx:    currentRoundIdx,
		CreatedAt:   ctx.Timestamp(),
		IsPaid:      false,
		PrizeAmount: 0,
	}
	tickets.GetTicket(ticketsLen).SetValue(ticket)

	currentRound.PrizePool += incomingAmount
	currentRound.TicketCount += 1
	rounds.GetRound(currentRoundIdx).SetValue(currentRound)

	f.Events.Buy(caller.Address(), incomingAmount, currentRoundIdx)
	f.Events.Round(
		currentRound.DrawAt,
		currentRound.Idx,
		currentRound.IsDrawn,
		currentRound.PrizePool,
		currentRound.TicketCount,
		currentRound.WinningNumber,
	)
}

func funcDraw(ctx wasmlib.ScFuncContext, f *DrawContext) {
	// get state
	tickets := f.State.Tickets()
	ticketsLen := tickets.Length()
	historyTickets := f.State.HistoryTickets()
	currentRoundIdx := f.State.CurrentRoundIdx().Value()
	rounds := f.State.Rounds()
	currentRound := rounds.GetRound(currentRoundIdx).Value()

	winningNumber := ctx.Random(MAX_NUMBER)
	winningNumberStr := strconv.FormatUint(winningNumber, 10)

	// pading zeroes to the left
	for i := len(winningNumberStr); i < MAX_DIGIT; i++ {
		winningNumberStr = "0" + winningNumberStr
	}

	currentRound.WinningNumber = winningNumberStr
	currentRound.IsDrawn = true

	// calculate prize pool
	prizePool := currentRound.PrizePool
	prizePoolsByDigit := make(map[int]uint64)
	prizePoolsByDigit[0] = 0
	prizePoolsByDigit[1] = prizePool * MATCH_1_PRIZE_PERCENT / 100
	prizePoolsByDigit[2] = prizePool * MATCH_2_PRIZE_PERCENT / 100
	prizePoolsByDigit[3] = prizePool * MATCH_3_PRIZE_PERCENT / 100
	prizePoolsByDigit[4] = prizePool * MATCH_4_PRIZE_PERCENT / 100
	prizePoolsByDigit[5] = prizePool * MATCH_5_PRIZE_PERCENT / 100
	prizePoolsByDigit[6] = prizePool * MATCH_6_PRIZE_PERCENT / 100

	// count winner by digit
	prizeWinnerByDigit := make(map[int]uint64)
	for i := uint32(0); i <= ticketsLen; i++ {
		ticket := tickets.GetTicket(i).Value()

		matchedCount := 0
		for i := 0; i <= MAX_DIGIT; i++ {
			if ticket.Number[i] == winningNumberStr[i] {
				matchedCount += 1
			}
		}
		prizeWinnerByDigit[matchedCount] += ticket.Amount

		ticket.MatchedDigit = uint8(matchedCount)
		tickets.GetTicket(i).SetValue(ticket)
	}

	// distribute prize pool
	prizePerTicketsByDigit := make(map[int]uint64)
	for i := 1; i <= MAX_DIGIT; i++ {
		prizePerTicketsByDigit[i] = prizePoolsByDigit[i] / prizeWinnerByDigit[i]
	}

	// pay tickets
	for i := uint32(0); i <= ticketsLen; i++ {
		ticket := tickets.GetTicket(i).Value()

		paidAmount := prizePerTicketsByDigit[int(ticket.MatchedDigit)]
		if paidAmount > 0 {
			// tranfers amount to the ticket buyer
			transfers := wasmlib.NewScTransferBaseTokens(paidAmount)
			ctx.Send(ticket.Buyer.Address(), transfers)
			ticket.IsPaid = true
		}

		ticket.PrizeAmount = paidAmount

		// add ticket to history tickets
		historyTickets.GetTicket(i).SetValue(ticket)

		// clear ticket on the current round
		tickets.GetTicket(i).Delete()
	}

	// new round
	newRound := &Round{
		Idx:           currentRoundIdx + 1,
		DrawAt:        ctx.Timestamp() + ROUND_PERIOD,
		IsDrawn:       false,
		PrizePool:     0,
		TicketCount:   0,
		WinningNumber: "",
	}
	rounds.GetRound(newRound.Idx).SetValue(newRound)
	f.State.CurrentRoundIdx().SetValue(newRound.Idx)
}

func funcGetMyTickets(ctx wasmlib.ScFuncContext, f *GetMyTicketsContext) {
	resultTickets := f.Results.Tickets()
	stateTickets := f.State.Tickets()
	stateTicketsLen := stateTickets.Length()
	var n uint32
	for i := uint32(0); i < stateTicketsLen; i++ {
		Ticket := stateTickets.GetTicket(i).Value()
		if Ticket.Buyer == ctx.Caller() {
			resultTickets.GetTicket(n).SetValue(Ticket)
			n++
		}
	}
}

func funcGetMyHistoryTickets(ctx wasmlib.ScFuncContext, f *GetMyHistoryTicketsContext) {
	resultTickets := f.Results.Tickets()
	stateHistoryTickets := f.State.HistoryTickets()
	stateHistoryTicketsLen := stateHistoryTickets.Length()
	var n uint32
	for i := uint32(0); i < stateHistoryTicketsLen; i++ {
		Ticket := stateHistoryTickets.GetTicket(i).Value()
		if Ticket.Buyer == ctx.Caller() {
			resultTickets.GetTicket(n).SetValue(Ticket)
			n++
		}
	}
}

func viewGetCurrentRound(ctx wasmlib.ScViewContext, f *GetCurrentRoundContext) {
	resultRound := f.Results.Round()
	stateRounds := f.State.Rounds()
	resultRound.SetValue(stateRounds.GetRound(f.State.CurrentRoundIdx().Value()).Value())
}

func viewGetRound(ctx wasmlib.ScViewContext, f *GetRoundContext) {
	roundIdx := f.Params.RoundIdx().Value()

	resultRound := f.Results.Round()
	stateRounds := f.State.Rounds()
	resultRound.SetValue(stateRounds.GetRound(roundIdx).Value())
}
