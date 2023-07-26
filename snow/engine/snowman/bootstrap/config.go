// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"github.com/luxdefi/node/snow/engine/common"
	"github.com/luxdefi/node/snow/engine/common/queue"
	"github.com/luxdefi/node/snow/engine/snowman/block"
)

type Config struct {
	common.Config
	common.AllGetsServer

	// Blocked tracks operations that are blocked on blocks
	//
	// It should be guaranteed that `MissingIDs` should contain all IDs
	// referenced by the `MissingDependencies` that have not already been added
	// to the queue.
	Blocked *queue.JobsWithMissing

	VM block.ChainVM

	Bootstrapped func()
}
