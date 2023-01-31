<<<<<<< HEAD
<<<<<<< HEAD
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
>>>>>>> 53a8245a8 (Update consensus)
=======
// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
=======
// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
>>>>>>> 34554f662 (Update LICENSE)
>>>>>>> c5eafdb72 (Update LICENSE)
// See the file LICENSE for licensing terms.

package vertex

import (
	"testing"

	"github.com/stretchr/testify/require"

<<<<<<< HEAD
	"github.com/luxdefi/luxd/ids"
=======
	"github.com/ava-labs/avalanchego/ids"
>>>>>>> 53a8245a8 (Update consensus)
)

func TestBuildInvalid(t *testing.T) {
	chainID := ids.ID{1}
	height := uint64(2)
	parentIDs := []ids.ID{{4}, {5}}
	txs := [][]byte{{6}, {6}}
	_, err := Build(
		chainID,
		height,
		parentIDs,
		txs,
	)
	require.Error(t, err, "build should have errored because restrictions were provided in epoch 0")
}

func TestBuildValid(t *testing.T) {
	chainID := ids.ID{1}
	height := uint64(2)
	parentIDs := []ids.ID{{4}, {5}}
	txs := [][]byte{{7}, {6}}
	vtx, err := Build(
		chainID,
		height,
		parentIDs,
		txs,
	)
	require.NoError(t, err)
	require.Equal(t, chainID, vtx.ChainID())
	require.Equal(t, height, vtx.Height())
	require.Equal(t, parentIDs, vtx.ParentIDs())
	require.Equal(t, txs, vtx.Txs())
}
