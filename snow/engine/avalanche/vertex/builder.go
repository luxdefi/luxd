// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"context"

	"github.com/luxdefi/node/ids"
	"github.com/luxdefi/node/snow/consensus/avalanche"
	"github.com/luxdefi/node/utils"
	"github.com/luxdefi/node/utils/hashing"
)

// Builder builds a vertex given a set of parentIDs and transactions.
type Builder interface {
	// Build a new stop vertex from the parents
	BuildStopVtx(ctx context.Context, parentIDs []ids.ID) (avalanche.Vertex, error)
}

// Build a new stateless vertex from the contents of a vertex
func Build(
	chainID ids.ID,
	height uint64,
	parentIDs []ids.ID,
	txs [][]byte,
) (StatelessVertex, error) {
	return buildVtx(
		chainID,
		height,
		parentIDs,
		txs,
		func(vtx innerStatelessVertex) error {
			return vtx.verify()
		},
		false,
	)
}

// Build a new stateless vertex from the contents of a vertex
func BuildStopVertex(chainID ids.ID, height uint64, parentIDs []ids.ID) (StatelessVertex, error) {
	return buildVtx(
		chainID,
		height,
		parentIDs,
		nil,
		func(vtx innerStatelessVertex) error {
			return vtx.verifyStopVertex()
		},
		true,
	)
}

func buildVtx(
	chainID ids.ID,
	height uint64,
	parentIDs []ids.ID,
	txs [][]byte,
	verifyFunc func(innerStatelessVertex) error,
	stopVertex bool,
) (StatelessVertex, error) {
	utils.Sort(parentIDs)
	utils.SortByHash(txs)

	codecVer := codecVersion
	if stopVertex {
		// use new codec version for the "StopVertex"
		codecVer = codecVersionWithStopVtx
	}

	innerVtx := innerStatelessVertex{
		Version:   codecVer,
		ChainID:   chainID,
		Height:    height,
		Epoch:     0,
		ParentIDs: parentIDs,
		Txs:       txs,
	}
	if err := verifyFunc(innerVtx); err != nil {
		return nil, err
	}

	vtxBytes, err := c.Marshal(innerVtx.Version, innerVtx)
	vtx := statelessVertex{
		innerStatelessVertex: innerVtx,
		id:                   hashing.ComputeHash256Array(vtxBytes),
		bytes:                vtxBytes,
	}
	return vtx, err
}
