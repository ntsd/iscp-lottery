// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package iscplotterycontract

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
    	FuncCreateTicket,
    	FuncDraw,
    	FuncGetMyHistoryTickets,
    	FuncGetMyTickets,
    	FuncInit,
    	FuncSetOwner,
    	ViewGetCurrentRound,
    	ViewGetOwner,
    	ViewGetRound,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
    	funcCreateTicketThunk,
    	funcDrawThunk,
    	funcGetMyHistoryTicketsThunk,
    	funcGetMyTicketsThunk,
    	funcInitThunk,
    	funcSetOwnerThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
    	viewGetCurrentRoundThunk,
    	viewGetOwnerThunk,
    	viewGetRoundThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type CreateTicketContext struct {
	Events  ISCPLotteryContractEvents
	Params  ImmutableCreateTicketParams
	State   MutableISCPLotteryContractState
}

func funcCreateTicketThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcCreateTicket")
	f := &CreateTicketContext{
		Params: ImmutableCreateTicketParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Number().Exists(), "missing mandatory number")
	funcCreateTicket(ctx, f)
	ctx.Log("iscplotterycontract.funcCreateTicket ok")
}

type DrawContext struct {
	Events  ISCPLotteryContractEvents
	State   MutableISCPLotteryContractState
}

func funcDrawThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcDraw")
	f := &DrawContext{
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.AccountID(), "no permission")

	funcDraw(ctx, f)
	ctx.Log("iscplotterycontract.funcDraw ok")
}

type GetMyHistoryTicketsContext struct {
	Events  ISCPLotteryContractEvents
	Results MutableGetMyHistoryTicketsResults
	State   MutableISCPLotteryContractState
}

func funcGetMyHistoryTicketsThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcGetMyHistoryTickets")
	results := wasmlib.NewScDict()
	f := &GetMyHistoryTicketsContext{
		Results: MutableGetMyHistoryTicketsResults{
			proxy: results.AsProxy(),
		},
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcGetMyHistoryTickets(ctx, f)
	ctx.Results(results)
	ctx.Log("iscplotterycontract.funcGetMyHistoryTickets ok")
}

type GetMyTicketsContext struct {
	Events  ISCPLotteryContractEvents
	Results MutableGetMyTicketsResults
	State   MutableISCPLotteryContractState
}

func funcGetMyTicketsThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcGetMyTickets")
	results := wasmlib.NewScDict()
	f := &GetMyTicketsContext{
		Results: MutableGetMyTicketsResults{
			proxy: results.AsProxy(),
		},
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcGetMyTickets(ctx, f)
	ctx.Results(results)
	ctx.Log("iscplotterycontract.funcGetMyTickets ok")
}

type InitContext struct {
	Events  ISCPLotteryContractEvents
	Params  ImmutableInitParams
	State   MutableISCPLotteryContractState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcInit(ctx, f)
	ctx.Log("iscplotterycontract.funcInit ok")
}

type SetOwnerContext struct {
	Events  ISCPLotteryContractEvents
	Params  ImmutableSetOwnerParams
	State   MutableISCPLotteryContractState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcSetOwner")
	f := &SetOwnerContext{
		Params: ImmutableSetOwnerParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// current owner of this smart contract
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.Owner().Exists(), "missing mandatory owner")
	funcSetOwner(ctx, f)
	ctx.Log("iscplotterycontract.funcSetOwner ok")
}

type GetCurrentRoundContext struct {
	Results MutableGetCurrentRoundResults
	State   ImmutableISCPLotteryContractState
}

func viewGetCurrentRoundThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("iscplotterycontract.viewGetCurrentRound")
	results := wasmlib.NewScDict()
	f := &GetCurrentRoundContext{
		Results: MutableGetCurrentRoundResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetCurrentRound(ctx, f)
	ctx.Results(results)
	ctx.Log("iscplotterycontract.viewGetCurrentRound ok")
}

type GetOwnerContext struct {
	Results MutableGetOwnerResults
	State   ImmutableISCPLotteryContractState
}

func viewGetOwnerThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("iscplotterycontract.viewGetOwner")
	results := wasmlib.NewScDict()
	f := &GetOwnerContext{
		Results: MutableGetOwnerResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetOwner(ctx, f)
	ctx.Results(results)
	ctx.Log("iscplotterycontract.viewGetOwner ok")
}

type GetRoundContext struct {
	Params  ImmutableGetRoundParams
	Results MutableGetRoundResults
	State   ImmutableISCPLotteryContractState
}

func viewGetRoundThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("iscplotterycontract.viewGetRound")
	results := wasmlib.NewScDict()
	f := &GetRoundContext{
		Params: ImmutableGetRoundParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetRoundResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableISCPLotteryContractState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.RoundIdx().Exists(), "missing mandatory roundIdx")
	viewGetRound(ctx, f)
	ctx.Results(results)
	ctx.Log("iscplotterycontract.viewGetRound ok")
}
