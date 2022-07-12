// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stateful

// var _ Block = &ProposalBlock{}

// // ProposalBlock is a proposal to change the chain's state.
// //
// // A proposal may be to:
// // 	1. Advance the chain's timestamp (*AdvanceTimeTx)
// //  2. Remove a staker from the staker set (*RewardStakerTx)
// //  3. Add a new staker to the set of pending (future) stakers
// //     (*AddValidatorTx, *AddDelegatorTx, *AddSubnetValidatorTx)
// //
// // The proposal will be enacted (change the chain's state) if the proposal block
// // is accepted and followed by an accepted Commit block
// type ProposalBlock struct {
// 	*stateless.ProposalBlock
// 	*commonBlock
// }

// // NewProposalBlock creates a new block that proposes to issue a transaction.
// // The parent of this block has ID [parentID].
// // The parent must be a decision block.
// func NewProposalBlock(
// 	manager Manager,
// 	ctx *snow.Context,
// 	parentID ids.ID,
// 	height uint64,
// 	tx *txs.Tx,
// ) (*ProposalBlock, error) {
// 	statelessBlk, err := stateless.NewProposalBlock(parentID, height, tx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return toStatefulProposalBlock(statelessBlk, manager, ctx, choices.Processing)
// }

// func toStatefulProposalBlock(
// 	statelessBlk *stateless.ProposalBlock,
// 	manager Manager,
// 	ctx *snow.Context,
// 	status choices.Status,
// ) (*ProposalBlock, error) {
// 	pb := &ProposalBlock{
// 		ProposalBlock: statelessBlk,
// 		commonBlock: &commonBlock{
// 			Manager: manager,
// 			baseBlk: &statelessBlk.CommonBlock,
// 		},
// 	}

// 	pb.Tx.Unsigned.InitCtx(ctx)
// 	return pb, nil
// }

// func (pb *ProposalBlock) Verify() error {
// 	return pb.VerifyProposalBlock(pb.ProposalBlock)
// }

// func (pb *ProposalBlock) Accept() error {
// 	return pb.AcceptProposalBlock(pb.ProposalBlock)
// }

// func (pb *ProposalBlock) Reject() error {
// 	return pb.RejectProposalBlock(pb.ProposalBlock)
// }

// func (pb *ProposalBlock) conflicts(s ids.Set) (bool, error) {
// 	return pb.conflictsProposalBlock(pb, s)
// }

// // Options returns the possible children of this block in preferential order.
// func (pb *ProposalBlock) Options() ([2]snowman.Block, error) {
// 	blkID := pb.ID()
// 	nextHeight := pb.Height() + 1

// 	preferCommit := pb.preferredCommit(blkID)
// 	commit, err := NewCommitBlock(
// 		pb.Manager,
// 		blkID,
// 		nextHeight,
// 	)
// 	if err != nil {
// 		return [2]snowman.Block{}, fmt.Errorf(
// 			"failed to create commit block: %w",
// 			err,
// 		)
// 	}
// 	abort, err := NewAbortBlock(
// 		pb.Manager,
// 		blkID,
// 		nextHeight,
// 	)
// 	if err != nil {
// 		return [2]snowman.Block{}, fmt.Errorf(
// 			"failed to create abort block: %w",
// 			err,
// 		)
// 	}

// 	if preferCommit {
// 		return [2]snowman.Block{commit, abort}, nil
// 	}
// 	return [2]snowman.Block{abort, commit}, nil
// }

// func (pb *ProposalBlock) setBaseState() {
// 	pb.setBaseStateProposalBlock(pb)
// }
