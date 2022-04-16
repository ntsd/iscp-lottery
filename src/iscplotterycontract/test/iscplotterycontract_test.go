// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"github.com/ntsd/iscp-lottery/src/iscplotterycontract/go/iscplotterycontract"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, iscplotterycontract.ScName, iscplotterycontract.OnLoad)
	require.NoError(t, ctx.ContractExists(iscplotterycontract.ScName))
}
