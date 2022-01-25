// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	heightIndex "github.com/ava-labs/avalanchego/vms/components/block_height_index"
)

// BlockServer represents all requests heightIndexer can issue
// against ProposerVM. All methods must be thread-safe.
type BlockServer interface {
	LastAcceptedWrappingBlkID() (ids.ID, error)
	LastAcceptedInnerBlkID() (ids.ID, error)
	GetWrappingBlk(blkID ids.ID) (WrappingBlock, error)
	GetInnerBlk(id ids.ID) (snowman.Block, error)
}

// heightIndexDBOps groups all the operations that indexer
// need to perform on state.HeightIndex
type heightIndexDBOps interface {
	heightIndex.Getter
	heightIndex.BatchSupport
}
