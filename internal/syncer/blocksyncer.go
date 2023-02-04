package syncer

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/web3"
)

type Controller struct {
	RDS         database.RDS
	Web3Service web3.Web3
}

func SyncConformedBlockBackward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := to; i > from; i-- {
		syncConformedBlock(i)
	}
}

func SyncConformedBlockForward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := from; i < to; i++ {
		syncConformedBlock(i)
	}
}

func SyncNewBlock(header *types.Header, c *Controller, aborter <-chan bool) {
	// TODO: implementation needed
}

func syncConformedBlock(blockNumber uint64) {
	// TODO: implementation needed
}
