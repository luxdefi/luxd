// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	stdcontext "context"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/vms/avm/txs"
	"github.com/luxdefi/node/vms/components/lux"
)

var _ txs.Visitor = (*backendVisitor)(nil)

// backendVisitor handles accepting of transactions for the backend
type backendVisitor struct {
	b    *backend
	ctx  stdcontext.Context
	txID ids.ID
}

func (*backendVisitor) BaseTx(*txs.BaseTx) error {
	return nil
}

func (*backendVisitor) CreateAssetTx(*txs.CreateAssetTx) error {
	return nil
}

func (*backendVisitor) OperationTx(*txs.OperationTx) error {
	return nil
}

func (b *backendVisitor) ImportTx(tx *txs.ImportTx) error {
	for _, in := range tx.ImportedIns {
		utxoID := in.UTXOID.InputID()
		if err := b.b.RemoveUTXO(b.ctx, tx.SourceChain, utxoID); err != nil {
			return err
		}
	}
	return nil
}

func (b *backendVisitor) ExportTx(tx *txs.ExportTx) error {
	for i, out := range tx.ExportedOuts {
		err := b.b.AddUTXO(
			b.ctx,
			tx.DestinationChain,
			&lux.UTXO{
				UTXOID: lux.UTXOID{
					TxID:        b.txID,
					OutputIndex: uint32(len(tx.Outs) + i),
				},
				Asset: lux.Asset{ID: out.AssetID()},
				Out:   out.Out,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
