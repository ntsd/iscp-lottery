// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "iscplotterycontract"
	ScDescription = "ISCPLotteryContract description"
	HScName       = wasmtypes.ScHname(0x884f32dd)
)

const (
	ParamNumber   = "number"
	ParamOwner    = "owner"
	ParamRoundIdx = "roundIdx"
)

const (
	ResultOwner   = "owner"
	ResultRound   = "round"
	ResultTickets = "tickets"
)

const (
	StateCurrentRoundIdx = "currentRoundIdx"
	StateHistoryTickets  = "historyTickets"
	StateOwner           = "owner"
	StateRounds          = "rounds"
	StateTickets         = "tickets"
)

const (
	FuncCreateTicket        = "createTicket"
	FuncDraw                = "draw"
	FuncGetMyHistoryTickets = "getMyHistoryTickets"
	FuncGetMyTickets        = "getMyTickets"
	FuncInit                = "init"
	FuncSetOwner            = "setOwner"
	ViewGetCurrentRound     = "getCurrentRound"
	ViewGetOwner            = "getOwner"
	ViewGetRound            = "getRound"
)

const (
	HFuncCreateTicket        = wasmtypes.ScHname(0x7a6044ef)
	HFuncDraw                = wasmtypes.ScHname(0x436fe402)
	HFuncGetMyHistoryTickets = wasmtypes.ScHname(0x2d7c6667)
	HFuncGetMyTickets        = wasmtypes.ScHname(0x12d805ce)
	HFuncInit                = wasmtypes.ScHname(0x1f44d644)
	HFuncSetOwner            = wasmtypes.ScHname(0x2a15fe7b)
	HViewGetCurrentRound     = wasmtypes.ScHname(0x7abd711b)
	HViewGetOwner            = wasmtypes.ScHname(0x137107a6)
	HViewGetRound            = wasmtypes.ScHname(0xbea7b465)
)
