// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package iscplotterycontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncCreateTicket,        funcCreateTicketThunk)
	exports.AddFunc(FuncDraw,                funcDrawThunk)
	exports.AddFunc(FuncGetMyHistoryTickets, funcGetMyHistoryTicketsThunk)
	exports.AddFunc(FuncGetMyTickets,        funcGetMyTicketsThunk)
	exports.AddFunc(FuncInit,                funcInitThunk)
	exports.AddFunc(FuncSetOwner,            funcSetOwnerThunk)
	exports.AddView(ViewGetCurrentRound,     viewGetCurrentRoundThunk)
	exports.AddView(ViewGetOwner,            viewGetOwnerThunk)
	exports.AddView(ViewGetRound,            viewGetRoundThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type CreateTicketContext struct {
	Params  ImmutableCreateTicketParams
	State   MutableISCPLotteryContractState
}

func funcCreateTicketThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcCreateTicket")
	f := &CreateTicketContext{
		Params: ImmutableCreateTicketParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.Number().Exists(), "missing mandatory number")
	funcCreateTicket(ctx, f)
	ctx.Log("iscplotterycontract.funcCreateTicket ok")
}

type DrawContext struct {
	State   MutableISCPLotteryContractState
}

func funcDrawThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcDraw")

	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.AccountID(), "no permission")

	f := &DrawContext{
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcDraw(ctx, f)
	ctx.Log("iscplotterycontract.funcDraw ok")
}

type GetMyHistoryTicketsContext struct {
	Results MutableGetMyHistoryTicketsResults
	State   MutableISCPLotteryContractState
}

func funcGetMyHistoryTicketsThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcGetMyHistoryTickets")
	f := &GetMyHistoryTicketsContext{
		Results: MutableGetMyHistoryTicketsResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcGetMyHistoryTickets(ctx, f)
	ctx.Log("iscplotterycontract.funcGetMyHistoryTickets ok")
}

type GetMyTicketsContext struct {
	Results MutableGetMyTicketsResults
	State   MutableISCPLotteryContractState
}

func funcGetMyTicketsThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcGetMyTickets")
	f := &GetMyTicketsContext{
		Results: MutableGetMyTicketsResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcGetMyTickets(ctx, f)
	ctx.Log("iscplotterycontract.funcGetMyTickets ok")
}

type InitContext struct {
	Params  ImmutableInitParams
	State   MutableISCPLotteryContractState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcInit(ctx, f)
	ctx.Log("iscplotterycontract.funcInit ok")
}

type SetOwnerContext struct {
	Params  ImmutableSetOwnerParams
	State   MutableISCPLotteryContractState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("iscplotterycontract.funcSetOwner")

	// current owner of this smart contract
	access := ctx.State().GetAgentID(wasmlib.Key("owner"))
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	f := &SetOwnerContext{
		Params: ImmutableSetOwnerParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		State: MutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
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
	f := &GetCurrentRoundContext{
		Results: MutableGetCurrentRoundResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetCurrentRound(ctx, f)
	ctx.Log("iscplotterycontract.viewGetCurrentRound ok")
}

type GetOwnerContext struct {
	Results MutableGetOwnerResults
	State   ImmutableISCPLotteryContractState
}

func viewGetOwnerThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("iscplotterycontract.viewGetOwner")
	f := &GetOwnerContext{
		Results: MutableGetOwnerResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetOwner(ctx, f)
	ctx.Log("iscplotterycontract.viewGetOwner ok")
}

type GetRoundContext struct {
	Params  ImmutableGetRoundParams
	Results MutableGetRoundResults
	State   ImmutableISCPLotteryContractState
}

func viewGetRoundThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("iscplotterycontract.viewGetRound")
	f := &GetRoundContext{
		Params: ImmutableGetRoundParams{
			id: wasmlib.OBJ_ID_PARAMS,
		},
		Results: MutableGetRoundResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutableISCPLotteryContractState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	ctx.Require(f.Params.RoundIdx().Exists(), "missing mandatory roundIdx")
	viewGetRound(ctx, f)
	ctx.Log("iscplotterycontract.viewGetRound ok")
}