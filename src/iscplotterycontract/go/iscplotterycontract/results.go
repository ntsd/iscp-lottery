// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ArrayOfImmutableTicket struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableTicket) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableTicket) GetTicket(index uint32) ImmutableTicket {
	return ImmutableTicket{proxy: a.proxy.Index(index)}
}

type ImmutableGetMyHistoryTicketsResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetMyHistoryTicketsResults) Tickets() ArrayOfImmutableTicket {
	return ArrayOfImmutableTicket{proxy: s.proxy.Root(ResultTickets)}
}

type ArrayOfMutableTicket struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableTicket) AppendTicket() MutableTicket {
	return MutableTicket{proxy: a.proxy.Append()}
}

func (a ArrayOfMutableTicket) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableTicket) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableTicket) GetTicket(index uint32) MutableTicket {
	return MutableTicket{proxy: a.proxy.Index(index)}
}

type MutableGetMyHistoryTicketsResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetMyHistoryTicketsResults) Tickets() ArrayOfMutableTicket {
	return ArrayOfMutableTicket{proxy: s.proxy.Root(ResultTickets)}
}

type ImmutableGetMyTicketsResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetMyTicketsResults) Tickets() ArrayOfImmutableTicket {
	return ArrayOfImmutableTicket{proxy: s.proxy.Root(ResultTickets)}
}

type MutableGetMyTicketsResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetMyTicketsResults) Tickets() ArrayOfMutableTicket {
	return ArrayOfMutableTicket{proxy: s.proxy.Root(ResultTickets)}
}

type ImmutableGetCurrentRoundResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetCurrentRoundResults) Round() ImmutableRound {
	return ImmutableRound{proxy: s.proxy.Root(ResultRound)}
}

type MutableGetCurrentRoundResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetCurrentRoundResults) Round() MutableRound {
	return MutableRound{proxy: s.proxy.Root(ResultRound)}
}

type ImmutableGetOwnerResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetOwnerResults) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ResultOwner))
}

type MutableGetOwnerResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetOwnerResults) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ResultOwner))
}

type ImmutableGetRoundResults struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRoundResults) Round() ImmutableRound {
	return ImmutableRound{proxy: s.proxy.Root(ResultRound)}
}

type MutableGetRoundResults struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRoundResults) Round() MutableRound {
	return MutableRound{proxy: s.proxy.Root(ResultRound)}
}
