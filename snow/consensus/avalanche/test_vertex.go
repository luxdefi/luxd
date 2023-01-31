// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/snow/consensus/snowstorm"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ Vertex = (*TestVertex)(nil)

// TestVertex is a useful test vertex
type TestVertex struct {
	choices.TestDecidable

	VerifyErrV    error
	ParentsV      []Vertex
	ParentsErrV   error
	HasWhitelistV bool
	WhitelistV    set.Set[ids.ID]
	WhitelistErrV error
	HeightV       uint64
	HeightErrV    error
	TxsV          []snowstorm.Tx
	TxsErrV       error
	BytesV        []byte
}

<<<<<<< HEAD
<<<<<<< HEAD
func (v *TestVertex) Verify(context.Context) error {
=======
func (v *TestVertex) Verify() error {
>>>>>>> 55bd9343c (Add EmptyLines linter (#2233))
=======
func (v *TestVertex) Verify(context.Context) error {
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
	return v.VerifyErrV
}

func (v *TestVertex) Parents() ([]Vertex, error) {
	return v.ParentsV, v.ParentsErrV
}

func (v *TestVertex) HasWhitelist() bool {
	return v.HasWhitelistV
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
func (v *TestVertex) Whitelist(context.Context) (set.Set[ids.ID], error) {
=======
func (v *TestVertex) Whitelist() (ids.Set, error) {
>>>>>>> 55bd9343c (Add EmptyLines linter (#2233))
=======
func (v *TestVertex) Whitelist(context.Context) (ids.Set, error) {
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
=======
func (v *TestVertex) Whitelist(context.Context) (set.Set[ids.ID], error) {
>>>>>>> 87ce2da8a (Replace type specific sets with a generic implementation (#1861))
	return v.WhitelistV, v.WhitelistErrV
}

func (v *TestVertex) Height() (uint64, error) {
	return v.HeightV, v.HeightErrV
}

<<<<<<< HEAD
<<<<<<< HEAD
func (v *TestVertex) Txs(context.Context) ([]snowstorm.Tx, error) {
=======
func (v *TestVertex) Txs() ([]snowstorm.Tx, error) {
>>>>>>> 55bd9343c (Add EmptyLines linter (#2233))
=======
func (v *TestVertex) Txs(context.Context) ([]snowstorm.Tx, error) {
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
	return v.TxsV, v.TxsErrV
}

func (v *TestVertex) Bytes() []byte {
	return v.BytesV
}
