// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/luxdefi/node/codec"
	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/utils/crypto"
	"github.com/luxdefi/node/vms/avm/fxs"
	"github.com/luxdefi/node/vms/avm/txs"
	"github.com/luxdefi/node/vms/components/avax"
	"github.com/luxdefi/node/vms/secp256k1fx"
)

var (
	networkID uint32 = 10
	chainID          = ids.GenerateTestID()
	keys             = crypto.BuildTestKeys()
	assetID          = ids.GenerateTestID()
)

func TestStandardBlocks(t *testing.T) {
	// check standard block can be built and parsed
	require := require.New(t)

	parser, err := NewParser([]fxs.Fx{
		&secp256k1fx.Fx{},
	})
	require.NoError(err)

	blkTimestamp := time.Now()
	parentID := ids.GenerateTestID()
	height := uint64(2022)
	cm := parser.Codec()
	txs, err := createTestTxs(cm)
	require.NoError(err)

	standardBlk, err := NewStandardBlock(parentID, height, blkTimestamp, txs, cm)
	require.NoError(err)

	// parse block
	parsed, err := parser.ParseBlock(standardBlk.Bytes())
	require.NoError(err)

	// compare content
	require.Equal(standardBlk.ID(), parsed.ID())
	require.Equal(standardBlk.Parent(), parsed.Parent())
	require.Equal(standardBlk.Height(), parsed.Height())
	require.Equal(standardBlk.Bytes(), parsed.Bytes())
	require.Equal(standardBlk.Timestamp(), parsed.Timestamp())

	parsedStandardBlk, ok := parsed.(*StandardBlock)
	require.True(ok)

	require.Equal(txs, parsedStandardBlk.Txs())
	require.Equal(parsed.Txs(), parsedStandardBlk.Txs())
}

func createTestTxs(cm codec.Manager) ([]*txs.Tx, error) {
	countTxs := 1
	testTxs := make([]*txs.Tx, 0, countTxs)
	for i := 0; i < countTxs; i++ {
		// Create the tx
		tx := &txs.Tx{Unsigned: &txs.BaseTx{BaseTx: avax.BaseTx{
			NetworkID:    networkID,
			BlockchainID: chainID,
			Outs: []*avax.TransferableOutput{{
				Asset: avax.Asset{ID: assetID},
				Out: &secp256k1fx.TransferOutput{
					Amt: uint64(12345),
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs:     []ids.ShortID{keys[0].PublicKey().Address()},
					},
				},
			}},
			Ins: []*avax.TransferableInput{{
				UTXOID: avax.UTXOID{
					TxID:        ids.ID{'t', 'x', 'I', 'D'},
					OutputIndex: 1,
				},
				Asset: avax.Asset{ID: assetID},
				In: &secp256k1fx.TransferInput{
					Amt: uint64(54321),
					Input: secp256k1fx.Input{
						SigIndices: []uint32{2},
					},
				},
			}},
			Memo: []byte{1, 2, 3, 4, 5, 6, 7, 8},
		}}}
		if err := tx.SignSECP256K1Fx(cm, [][]*crypto.PrivateKeySECP256K1R{{keys[0]}}); err != nil {
			return nil, err
		}
		testTxs = append(testTxs, tx)
	}
	return testTxs, nil
}
