// Copyright (C) 2019-2024, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/luxfi/node/ids"
	"github.com/luxfi/node/utils/constants"
	"github.com/luxfi/node/utils/crypto/bls"
	"github.com/luxfi/node/utils/crypto/secp256k1"
	"github.com/luxfi/node/utils/set"
	"github.com/luxfi/node/utils/units"
	"github.com/luxfi/node/vms/components/lux"
	"github.com/luxfi/node/vms/platformvm/reward"
	"github.com/luxfi/node/vms/platformvm/signer"
	"github.com/luxfi/node/vms/platformvm/stakeable"
	"github.com/luxfi/node/vms/platformvm/txs"
	"github.com/luxfi/node/vms/secp256k1fx"
	"github.com/luxfi/node/wallet/chain/p/builder"
	"github.com/luxfi/node/wallet/subnet/primary/common"
)

var (
	testKeys = secp256k1.TestKeys()

	// We hard-code [luxAssetID] and [subnetAssetID] to make
	// ordering of UTXOs generated by [testUTXOsList] is reproducible
	luxAssetID   = ids.Empty.Prefix(1789)
	subnetAssetID = ids.Empty.Prefix(2024)

	testContext = &builder.Context{
		NetworkID:                     constants.UnitTestID,
		LUXAssetID:                   luxAssetID,
		BaseTxFee:                     units.MicroLux,
		CreateSubnetTxFee:             19 * units.MicroLux,
		TransformSubnetTxFee:          789 * units.MicroLux,
		CreateBlockchainTxFee:         1234 * units.MicroLux,
		AddPrimaryNetworkValidatorFee: 19 * units.MilliLux,
		AddPrimaryNetworkDelegatorFee: 765 * units.MilliLux,
		AddSubnetValidatorFee:         1010 * units.MilliLux,
		AddSubnetDelegatorFee:         9 * units.Lux,
	}
)

// These tests create a tx, then verify that utxos included in the tx are
// exactly necessary to pay fees for it.

func TestBaseTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})
		backend = NewBackend(testContext, chainUTXOs, nil)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr), testContext, backend)

		// data to build the transaction
		outputsToMove = []*lux.TransferableOutput{{
			Asset: lux.Asset{ID: luxAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: 7 * units.Lux,
				OutputOwners: secp256k1fx.OutputOwners{
					Threshold: 1,
					Addrs:     []ids.ShortID{utxoAddr},
				},
			},
		}}
	)

	utx, err := builder.NewBaseTx(outputsToMove)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 2)
	require.Len(outs, 2)

	expectedConsumed := testContext.BaseTxFee + outputsToMove[0].Out.Amount()
	consumed := ins[0].In.Amount() + ins[1].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
	require.Equal(outputsToMove[0], outs[1])
}

func TestAddSubnetValidatorTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)

		// data to build the transaction
		subnetValidator = &txs.SubnetValidator{
			Validator: txs.Validator{
				NodeID: ids.GenerateTestNodeID(),
				End:    uint64(time.Now().Add(time.Hour).Unix()),
			},
			Subnet: subnetID,
		}
	)

	// build the transaction
	utx, err := builder.NewAddSubnetValidatorTx(subnetValidator)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 2)
	require.Len(outs, 1)

	expectedConsumed := testContext.AddSubnetValidatorFee
	consumed := ins[0].In.Amount() + ins[1].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestRemoveSubnetValidatorTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)
	)

	// build the transaction
	utx, err := builder.NewRemoveSubnetValidatorTx(
		ids.GenerateTestNodeID(),
		subnetID,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 1)
	require.Len(outs, 1)

	expectedConsumed := testContext.BaseTxFee
	consumed := ins[0].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestCreateChainTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)

		// data to build the transaction
		genesisBytes = []byte{'a', 'b', 'c'}
		vmID         = ids.GenerateTestID()
		fxIDs        = []ids.ID{ids.GenerateTestID()}
		chainName    = "dummyChain"
	)

	// build the transaction
	utx, err := builder.NewCreateChainTx(
		subnetID,
		genesisBytes,
		vmID,
		fxIDs,
		chainName,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 1)
	require.Len(outs, 1)

	expectedConsumed := testContext.CreateBlockchainTxFee
	consumed := ins[0].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestCreateSubnetTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)
	)

	// build the transaction
	utx, err := builder.NewCreateSubnetTx(subnetOwner)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 1)
	require.Len(outs, 1)

	expectedConsumed := testContext.CreateSubnetTxFee
	consumed := ins[0].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestTransferSubnetOwnershipTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)
	)

	// build the transaction
	utx, err := builder.NewTransferSubnetOwnershipTx(
		subnetID,
		subnetOwner,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 1)
	require.Len(outs, 1)

	expectedConsumed := testContext.BaseTxFee
	consumed := ins[0].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestImportTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey      = testKeys[1]
		utxos         = makeTestUTXOs(utxosKey)
		sourceChainID = ids.GenerateTestID()
		importedUTXOs = utxos[:1]
		chainUTXOs    = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
			sourceChainID:             importedUTXOs,
		})

		backend = NewBackend(testContext, chainUTXOs, nil)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr), testContext, backend)

		// data to build the transaction
		importKey = testKeys[0]
		importTo  = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs: []ids.ShortID{
				importKey.Address(),
			},
		}
	)

	// build the transaction
	utx, err := builder.NewImportTx(
		sourceChainID,
		importTo,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	importedIns := utx.ImportedInputs
	require.Empty(ins) // we spend the imported input (at least partially)
	require.Len(importedIns, 1)
	require.Len(outs, 1)

	expectedConsumed := testContext.BaseTxFee
	consumed := importedIns[0].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestExportTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})
		backend = NewBackend(testContext, chainUTXOs, nil)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr), testContext, backend)

		// data to build the transaction
		subnetID        = ids.GenerateTestID()
		exportedOutputs = []*lux.TransferableOutput{{
			Asset: lux.Asset{ID: luxAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: 7 * units.Lux,
				OutputOwners: secp256k1fx.OutputOwners{
					Threshold: 1,
					Addrs:     []ids.ShortID{utxoAddr},
				},
			},
		}}
	)

	// build the transaction
	utx, err := builder.NewExportTx(
		subnetID,
		exportedOutputs,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 2)
	require.Len(outs, 1)

	expectedConsumed := testContext.BaseTxFee + exportedOutputs[0].Out.Amount()
	consumed := ins[0].In.Amount() + ins[1].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
	require.Equal(utx.ExportedOutputs, exportedOutputs)
}

func TestTransformSubnetTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})

		subnetID       = ids.GenerateTestID()
		subnetAuthKey  = testKeys[0]
		subnetAuthAddr = subnetAuthKey.Address()
		subnetOwner    = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs:     []ids.ShortID{subnetAuthAddr},
		}
		subnets = map[ids.ID]*txs.Tx{
			subnetID: {
				Unsigned: &txs.CreateSubnetTx{
					Owner: subnetOwner,
				},
			},
		}

		backend = NewBackend(testContext, chainUTXOs, subnets)

		// builder
		utxoAddr = utxosKey.Address()
		builder  = builder.New(set.Of(utxoAddr, subnetAuthAddr), testContext, backend)

		// data to build the transaction
		initialSupply = 40 * units.MegaLux
		maxSupply     = 100 * units.MegaLux
	)

	// build the transaction
	utx, err := builder.NewTransformSubnetTx(
		subnetID,
		subnetAssetID,
		initialSupply,                 // initial supply
		maxSupply,                     // max supply
		reward.PercentDenominator,     // min consumption rate
		reward.PercentDenominator,     // max consumption rate
		1,                             // min validator stake
		100*units.MegaLux,            // max validator stake
		time.Second,                   // min stake duration
		365*24*time.Hour,              // max stake duration
		0,                             // min delegation fee
		1,                             // min delegator stake
		5,                             // max validator weight factor
		.80*reward.PercentDenominator, // uptime requirement
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	outs := utx.Outs
	require.Len(ins, 2)
	require.Len(outs, 2)

	expectedConsumedSubnetAsset := maxSupply - initialSupply
	consumedSubnetAsset := ins[0].In.Amount() - outs[1].Out.Amount()
	require.Equal(expectedConsumedSubnetAsset, consumedSubnetAsset)
	expectedConsumed := testContext.TransformSubnetTxFee
	consumed := ins[1].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestAddPermissionlessValidatorTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})
		backend = NewBackend(testContext, chainUTXOs, nil)

		// builder
		utxoAddr   = utxosKey.Address()
		rewardKey  = testKeys[0]
		rewardAddr = rewardKey.Address()
		builder    = builder.New(set.Of(utxoAddr, rewardAddr), testContext, backend)

		// data to build the transaction
		validationRewardsOwner = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs: []ids.ShortID{
				rewardAddr,
			},
		}
		delegationRewardsOwner = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs: []ids.ShortID{
				rewardAddr,
			},
		}
	)

	sk, err := bls.NewSecretKey()
	require.NoError(err)

	// build the transaction
	utx, err := builder.NewAddPermissionlessValidatorTx(
		&txs.SubnetValidator{
			Validator: txs.Validator{
				NodeID: ids.GenerateTestNodeID(),
				End:    uint64(time.Now().Add(time.Hour).Unix()),
				Wght:   2 * units.Lux,
			},
			Subnet: constants.PrimaryNetworkID,
		},
		signer.NewProofOfPossession(sk),
		luxAssetID,
		validationRewardsOwner,
		delegationRewardsOwner,
		reward.PercentDenominator,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	staked := utx.StakeOuts
	outs := utx.Outs
	require.Len(ins, 4)
	require.Len(staked, 2)
	require.Len(outs, 2)

	expectedConsumedSubnetAsset := utx.Validator.Weight()
	consumedSubnetAsset := staked[0].Out.Amount() + staked[1].Out.Amount()
	require.Equal(expectedConsumedSubnetAsset, consumedSubnetAsset)
	expectedConsumed := testContext.AddPrimaryNetworkValidatorFee
	consumed := ins[1].In.Amount() + ins[3].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func TestAddPermissionlessDelegatorTx(t *testing.T) {
	var (
		require = require.New(t)

		// backend
		utxosKey   = testKeys[1]
		utxos      = makeTestUTXOs(utxosKey)
		chainUTXOs = common.NewDeterministicChainUTXOs(require, map[ids.ID][]*lux.UTXO{
			constants.PlatformChainID: utxos,
		})
		backend = NewBackend(testContext, chainUTXOs, nil)

		// builder
		utxoAddr   = utxosKey.Address()
		rewardKey  = testKeys[0]
		rewardAddr = rewardKey.Address()
		builder    = builder.New(set.Of(utxoAddr, rewardAddr), testContext, backend)

		// data to build the transaction
		rewardsOwner = &secp256k1fx.OutputOwners{
			Threshold: 1,
			Addrs: []ids.ShortID{
				rewardAddr,
			},
		}
	)

	// build the transaction
	utx, err := builder.NewAddPermissionlessDelegatorTx(
		&txs.SubnetValidator{
			Validator: txs.Validator{
				NodeID: ids.GenerateTestNodeID(),
				End:    uint64(time.Now().Add(time.Hour).Unix()),
				Wght:   2 * units.Lux,
			},
			Subnet: constants.PrimaryNetworkID,
		},
		luxAssetID,
		rewardsOwner,
	)
	require.NoError(err)

	// check UTXOs selection and fee financing
	ins := utx.Ins
	staked := utx.StakeOuts
	outs := utx.Outs
	require.Len(ins, 4)
	require.Len(staked, 2)
	require.Len(outs, 2)

	expectedConsumedSubnetAsset := utx.Validator.Weight()
	consumedSubnetAsset := staked[0].Out.Amount() + staked[1].Out.Amount()
	require.Equal(expectedConsumedSubnetAsset, consumedSubnetAsset)
	expectedConsumed := testContext.AddPrimaryNetworkDelegatorFee
	consumed := ins[1].In.Amount() + ins[3].In.Amount() - outs[0].Out.Amount()
	require.Equal(expectedConsumed, consumed)
}

func makeTestUTXOs(utxosKey *secp256k1.PrivateKey) []*lux.UTXO {
	// Note: we avoid ids.GenerateTestNodeID here to make sure that UTXO IDs won't change
	// run by run. This simplifies checking what utxos are included in the built txs.
	const utxosOffset uint64 = 2024

	utxosAddr := utxosKey.Address()
	return []*lux.UTXO{
		{ // a small UTXO first, which should not be enough to pay fees
			UTXOID: lux.UTXOID{
				TxID:        ids.Empty.Prefix(utxosOffset),
				OutputIndex: uint32(utxosOffset),
			},
			Asset: lux.Asset{ID: luxAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: 2 * units.MilliLux,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Addrs:     []ids.ShortID{utxosAddr},
					Threshold: 1,
				},
			},
		},
		{ // a locked, small UTXO
			UTXOID: lux.UTXOID{
				TxID:        ids.Empty.Prefix(utxosOffset + 1),
				OutputIndex: uint32(utxosOffset + 1),
			},
			Asset: lux.Asset{ID: luxAssetID},
			Out: &stakeable.LockOut{
				Locktime: uint64(time.Now().Add(time.Hour).Unix()),
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: 3 * units.MilliLux,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs:     []ids.ShortID{utxosAddr},
					},
				},
			},
		},
		{ // a subnetAssetID denominated UTXO
			UTXOID: lux.UTXOID{
				TxID:        ids.Empty.Prefix(utxosOffset + 2),
				OutputIndex: uint32(utxosOffset + 2),
			},
			Asset: lux.Asset{ID: subnetAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: 99 * units.MegaLux,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Addrs:     []ids.ShortID{utxosAddr},
					Threshold: 1,
				},
			},
		},
		{ // a locked, large UTXO
			UTXOID: lux.UTXOID{
				TxID:        ids.Empty.Prefix(utxosOffset + 3),
				OutputIndex: uint32(utxosOffset + 3),
			},
			Asset: lux.Asset{ID: luxAssetID},
			Out: &stakeable.LockOut{
				Locktime: uint64(time.Now().Add(time.Hour).Unix()),
				TransferableOut: &secp256k1fx.TransferOutput{
					Amt: 88 * units.Lux,
					OutputOwners: secp256k1fx.OutputOwners{
						Threshold: 1,
						Addrs:     []ids.ShortID{utxosAddr},
					},
				},
			},
		},
		{ // a large UTXO last, which should be enough to pay any fee by itself
			UTXOID: lux.UTXOID{
				TxID:        ids.Empty.Prefix(utxosOffset + 4),
				OutputIndex: uint32(utxosOffset + 4),
			},
			Asset: lux.Asset{ID: luxAssetID},
			Out: &secp256k1fx.TransferOutput{
				Amt: 9 * units.Lux,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  0,
					Addrs:     []ids.ShortID{utxosAddr},
					Threshold: 1,
				},
			},
		},
	}
}
