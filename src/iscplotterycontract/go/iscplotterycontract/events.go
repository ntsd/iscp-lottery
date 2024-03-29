// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:gocritic
package iscplotterycontract

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ISCPLotteryContractEvents struct {
}

func (e ISCPLotteryContractEvents) Buy(address wasmtypes.ScAddress, amount uint64, roundIdx uint32) {
	evt := wasmlib.NewEventEncoder("iscplotterycontract.buy")
	evt.Encode(wasmtypes.AddressToString(address))
	evt.Encode(wasmtypes.Uint64ToString(amount))
	evt.Encode(wasmtypes.Uint32ToString(roundIdx))
	evt.Emit()
}

func (e ISCPLotteryContractEvents) Round(drawAt uint64, idx uint32, isDrawn bool, prizePool uint64, ticketCount uint32, winningNumber string) {
	evt := wasmlib.NewEventEncoder("iscplotterycontract.round")
	evt.Encode(wasmtypes.Uint64ToString(drawAt))
	evt.Encode(wasmtypes.Uint32ToString(idx))
	evt.Encode(wasmtypes.BoolToString(isDrawn))
	evt.Encode(wasmtypes.Uint64ToString(prizePool))
	evt.Encode(wasmtypes.Uint32ToString(ticketCount))
	evt.Encode(wasmtypes.StringToString(winningNumber))
	evt.Emit()
}

func (e ISCPLotteryContractEvents) RoundFinish(nextRoundIdx uint32, roundIdx uint32) {
	evt := wasmlib.NewEventEncoder("iscplotterycontract.roundFinish")
	evt.Encode(wasmtypes.Uint32ToString(nextRoundIdx))
	evt.Encode(wasmtypes.Uint32ToString(roundIdx))
	evt.Emit()
}
