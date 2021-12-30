// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type CreateTicketCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCreateTicketParams
}

type DrawCall struct {
	Func    *wasmlib.ScFunc
}

type GetMyHistoryTicketsCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableGetMyHistoryTicketsResults
}

type GetMyTicketsCall struct {
	Func    *wasmlib.ScFunc
	Results ImmutableGetMyTicketsResults
}

type InitCall struct {
	Func    *wasmlib.ScInitFunc
	Params  MutableInitParams
}

type SetOwnerCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableSetOwnerParams
}

type GetCurrentRoundCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetCurrentRoundResults
}

type GetOwnerCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetOwnerResults
}

type GetRoundCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetRoundParams
	Results ImmutableGetRoundResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) CreateTicket(ctx wasmlib.ScFuncCallContext) *CreateTicketCall {
	f := &CreateTicketCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCreateTicket)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Draw(ctx wasmlib.ScFuncCallContext) *DrawCall {
	return &DrawCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncDraw)}
}

func (sc Funcs) GetMyHistoryTickets(ctx wasmlib.ScFuncCallContext) *GetMyHistoryTicketsCall {
	f := &GetMyHistoryTicketsCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncGetMyHistoryTickets)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetMyTickets(ctx wasmlib.ScFuncCallContext) *GetMyTicketsCall {
	f := &GetMyTicketsCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncGetMyTickets)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit, keyMap[:], idxMap[:])}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) SetOwner(ctx wasmlib.ScFuncCallContext) *SetOwnerCall {
	f := &SetOwnerCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetOwner)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) GetCurrentRound(ctx wasmlib.ScViewCallContext) *GetCurrentRoundCall {
	f := &GetCurrentRoundCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetCurrentRound)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetOwner(ctx wasmlib.ScViewCallContext) *GetOwnerCall {
	f := &GetOwnerCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetOwner)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}

func (sc Funcs) GetRound(ctx wasmlib.ScViewCallContext) *GetRoundCall {
	f := &GetRoundCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetRound)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}
