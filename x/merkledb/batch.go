// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import "github.com/luxdefi/node/database"

var _ database.Batch = (*batch)(nil)

// batch is a write-only database that commits changes to its host database
// when Write is called.
type batch struct {
	database.BatchOps

	db *merkleDB
}

// apply all operations in order to the database and write the result to disk
func (b *batch) Write() error {
	return b.db.commitBatch(b.Ops)
}

func (b *batch) Inner() database.Batch {
	return b
}
