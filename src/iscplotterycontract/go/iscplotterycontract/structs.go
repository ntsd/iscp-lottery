// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type Round struct {
	DrawAt        uint64  // unix timestamp when draw of the round
	Idx           uint32  // index of lottery round
	IsDrawn       bool  // is the round drawn?
	PrizePool     uint64  // iota price pool of the round
	TicketCount   uint32  // number of tickets for the round
	WinningNumber string  // the winning numbers of the round
}

func NewRoundFromBytes(buf []byte) *Round {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Round{}
	data.DrawAt        = wasmtypes.Uint64Decode(dec)
	data.Idx           = wasmtypes.Uint32Decode(dec)
	data.IsDrawn       = wasmtypes.BoolDecode(dec)
	data.PrizePool     = wasmtypes.Uint64Decode(dec)
	data.TicketCount   = wasmtypes.Uint32Decode(dec)
	data.WinningNumber = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *Round) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.DrawAt)
		wasmtypes.Uint32Encode(enc, o.Idx)
		wasmtypes.BoolEncode(enc, o.IsDrawn)
		wasmtypes.Uint64Encode(enc, o.PrizePool)
		wasmtypes.Uint32Encode(enc, o.TicketCount)
		wasmtypes.StringEncode(enc, o.WinningNumber)
	return enc.Buf()
}

type ImmutableRound struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableRound) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableRound) Value() *Round {
	return NewRoundFromBytes(o.proxy.Get())
}

type MutableRound struct {
	proxy wasmtypes.Proxy
}

func (o MutableRound) Delete() {
	o.proxy.Delete()
}

func (o MutableRound) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableRound) SetValue(value *Round) {
	o.proxy.Set(value.Bytes())
}

func (o MutableRound) Value() *Round {
	return NewRoundFromBytes(o.proxy.Get())
}

type Ticket struct {
	Amount       uint64  // iotas amount of the ticket
	Buyer        wasmtypes.ScAgentID  // the ticket buyer
	CreatedAt    uint64  // unix timestamp when the ticket create
	Id           wasmtypes.ScRequestID  // request id of the ticket
	Idx          uint32  // index of ticket
	IsPaid       bool  // is the prize paid?
	MatchedDigit uint8  // digits matched of the prize
	Number       string  // the lottery number of the ticket
	PrizeAmount  uint64  // iotas prize amount of the ticket
	RoundIdx     uint32  // the lottery round number
}

func NewTicketFromBytes(buf []byte) *Ticket {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Ticket{}
	data.Amount       = wasmtypes.Uint64Decode(dec)
	data.Buyer        = wasmtypes.AgentIDDecode(dec)
	data.CreatedAt    = wasmtypes.Uint64Decode(dec)
	data.Id           = wasmtypes.RequestIDDecode(dec)
	data.Idx          = wasmtypes.Uint32Decode(dec)
	data.IsPaid       = wasmtypes.BoolDecode(dec)
	data.MatchedDigit = wasmtypes.Uint8Decode(dec)
	data.Number       = wasmtypes.StringDecode(dec)
	data.PrizeAmount  = wasmtypes.Uint64Decode(dec)
	data.RoundIdx     = wasmtypes.Uint32Decode(dec)
	dec.Close()
	return data
}

func (o *Ticket) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.Amount)
		wasmtypes.AgentIDEncode(enc, o.Buyer)
		wasmtypes.Uint64Encode(enc, o.CreatedAt)
		wasmtypes.RequestIDEncode(enc, o.Id)
		wasmtypes.Uint32Encode(enc, o.Idx)
		wasmtypes.BoolEncode(enc, o.IsPaid)
		wasmtypes.Uint8Encode(enc, o.MatchedDigit)
		wasmtypes.StringEncode(enc, o.Number)
		wasmtypes.Uint64Encode(enc, o.PrizeAmount)
		wasmtypes.Uint32Encode(enc, o.RoundIdx)
	return enc.Buf()
}

type ImmutableTicket struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableTicket) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableTicket) Value() *Ticket {
	return NewTicketFromBytes(o.proxy.Get())
}

type MutableTicket struct {
	proxy wasmtypes.Proxy
}

func (o MutableTicket) Delete() {
	o.proxy.Delete()
}

func (o MutableTicket) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableTicket) SetValue(value *Ticket) {
	o.proxy.Set(value.Bytes())
}

func (o MutableTicket) Value() *Ticket {
	return NewTicketFromBytes(o.proxy.Get())
}
