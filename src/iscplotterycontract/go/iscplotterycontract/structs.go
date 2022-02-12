// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type Round struct {
	DrawAt        int64  // unix timestamp when draw of the round
	Idx           int32  // index of lottery round
	IsDrawn       bool  // is the round drawn?
	PrizePool     int64  // iota price pool of the round
	WinningNumber string  // the winning numbers of the round
}

func NewRoundFromBytes(bytes []byte) *Round {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Round{}
	data.DrawAt        = decode.Int64()
	data.Idx           = decode.Int32()
	data.IsDrawn       = decode.Bool()
	data.PrizePool     = decode.Int64()
	data.WinningNumber = decode.String()
	decode.Close()
	return data
}

func (o *Round) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.DrawAt).
		Int32(o.Idx).
		Bool(o.IsDrawn).
		Int64(o.PrizePool).
		String(o.WinningNumber).
		Data()
}

type ImmutableRound struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableRound) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableRound) Value() *Round {
	return NewRoundFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableRound struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableRound) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableRound) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableRound) SetValue(value *Round) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableRound) Value() *Round {
	return NewRoundFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type Ticket struct {
	Amount       int64  // iota amount of the ticket
	Buyer        wasmlib.ScAgentID  // the ticket buyer
	CreatedAt    int64  // unix timestamp when the ticket create
	Id           wasmlib.ScRequestID  // request id of the ticket
	Idx          int32  // index of ticket
	IsPaid       bool  // is the prize paid?
	MatchedDigit uint8  // how many digits matched
	Number       string  // the lottery number of the ticket
	PrizeAmount  int64  // iota prize amount of the ticket
	RoundIdx     int32  // the lottery round number
}

func NewTicketFromBytes(bytes []byte) *Ticket {
	decode := wasmlib.NewBytesDecoder(bytes)
	data := &Ticket{}
	data.Amount       = decode.Int64()
	data.Buyer        = decode.AgentID()
	data.CreatedAt    = decode.Int64()
	data.Id           = decode.RequestID()
	data.Idx          = decode.Int32()
	data.IsPaid       = decode.Bool()
	data.MatchedDigit = decode.Uint8()
	data.Number       = decode.String()
	data.PrizeAmount  = decode.Int64()
	data.RoundIdx     = decode.Int32()
	decode.Close()
	return data
}

func (o *Ticket) Bytes() []byte {
	return wasmlib.NewBytesEncoder().
		Int64(o.Amount).
		AgentID(o.Buyer).
		Int64(o.CreatedAt).
		RequestID(o.Id).
		Int32(o.Idx).
		Bool(o.IsPaid).
		Uint8(o.MatchedDigit).
		String(o.Number).
		Int64(o.PrizeAmount).
		Int32(o.RoundIdx).
		Data()
}

type ImmutableTicket struct {
	objID int32
	keyID wasmlib.Key32
}

func (o ImmutableTicket) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o ImmutableTicket) Value() *Ticket {
	return NewTicketFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}

type MutableTicket struct {
	objID int32
	keyID wasmlib.Key32
}

func (o MutableTicket) Delete() {
	wasmlib.DelKey(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableTicket) Exists() bool {
	return wasmlib.Exists(o.objID, o.keyID, wasmlib.TYPE_BYTES)
}

func (o MutableTicket) SetValue(value *Ticket) {
	wasmlib.SetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES, value.Bytes())
}

func (o MutableTicket) Value() *Ticket {
	return NewTicketFromBytes(wasmlib.GetBytes(o.objID, o.keyID, wasmlib.TYPE_BYTES))
}
