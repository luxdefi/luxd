// Copyright (C) 2022, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	"sync"

	stdcontext "context"

<<<<<<< HEAD
	"github.com/luxdefi/luxd/database"
	"github.com/luxdefi/luxd/ids"
	"github.com/luxdefi/luxd/utils/constants"
	"github.com/luxdefi/luxd/vms/components/lux"
	"github.com/luxdefi/luxd/vms/platformvm/txs"
=======
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
)

var _ Backend = (*backend)(nil)

type ChainUTXOs interface {
	AddUTXO(ctx stdcontext.Context, destinationChainID ids.ID, utxo *lux.UTXO) error
	RemoveUTXO(ctx stdcontext.Context, sourceChainID, utxoID ids.ID) error

	UTXOs(ctx stdcontext.Context, sourceChainID ids.ID) ([]*lux.UTXO, error)
	GetUTXO(ctx stdcontext.Context, sourceChainID, utxoID ids.ID) (*lux.UTXO, error)
}

// Backend defines the full interface required to support a P-chain wallet.
type Backend interface {
	ChainUTXOs
	BuilderBackend
	SignerBackend

	AcceptTx(ctx stdcontext.Context, tx *txs.Tx) error
}

type backend struct {
	Context
	ChainUTXOs

	txsLock sync.RWMutex
	// txID -> tx
	txs map[ids.ID]*txs.Tx
}

func NewBackend(ctx Context, utxos ChainUTXOs, txs map[ids.ID]*txs.Tx) Backend {
	return &backend{
		Context:    ctx,
		ChainUTXOs: utxos,
		txs:        txs,
	}
}

func (b *backend) AcceptTx(ctx stdcontext.Context, tx *txs.Tx) error {
	txID := tx.ID()
	err := tx.Unsigned.Visit(&backendVisitor{
		b:    b,
		ctx:  ctx,
		txID: txID,
	})
	if err != nil {
		return err
	}

	producedUTXOSlice := tx.UTXOs()
	err = b.addUTXOs(ctx, constants.PlatformChainID, producedUTXOSlice)
	if err != nil {
		return err
	}

	b.txsLock.Lock()
	defer b.txsLock.Unlock()

	b.txs[txID] = tx
	return nil
}

func (b *backend) addUTXOs(ctx stdcontext.Context, destinationChainID ids.ID, utxos []*lux.UTXO) error {
	for _, utxo := range utxos {
		if err := b.AddUTXO(ctx, destinationChainID, utxo); err != nil {
			return err
		}
	}
	return nil
}

func (b *backend) removeUTXOs(ctx stdcontext.Context, sourceChain ids.ID, utxoIDs set.Set[ids.ID]) error {
	for utxoID := range utxoIDs {
		if err := b.RemoveUTXO(ctx, sourceChain, utxoID); err != nil {
			return err
		}
	}
	return nil
}

func (b *backend) GetTx(_ stdcontext.Context, txID ids.ID) (*txs.Tx, error) {
	b.txsLock.RLock()
	defer b.txsLock.RUnlock()

	tx, exists := b.txs[txID]
	if !exists {
		return nil, database.ErrNotFound
	}
	return tx, nil
}
