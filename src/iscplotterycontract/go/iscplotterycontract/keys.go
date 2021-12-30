// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package iscplotterycontract

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamNumber   = 0
	IdxParamOwner    = 1
	IdxParamRoundIdx = 2

	IdxResultOwner   = 3
	IdxResultRound   = 4
	IdxResultTickets = 5

	IdxStateCurrentRoundIdx = 6
	IdxStateHistoryTickets  = 7
	IdxStateOwner           = 8
	IdxStateRounds          = 9
	IdxStateTickets         = 10
)

const keyMapLen = 11

var keyMap = [keyMapLen]wasmlib.Key{
	ParamNumber,
	ParamOwner,
	ParamRoundIdx,
	ResultOwner,
	ResultRound,
	ResultTickets,
	StateCurrentRoundIdx,
	StateHistoryTickets,
	StateOwner,
	StateRounds,
	StateTickets,
}

var idxMap [keyMapLen]wasmlib.Key32
