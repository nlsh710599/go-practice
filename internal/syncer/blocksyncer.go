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

func SyncConformedBlockBackward(from uint64, to uint64, c *Controller) {
	// TODO: implementation needed
}

func SyncConformedBlockForward(from uint64, to uint64, c *Controller) {
	// TODO: implementation needed
}

func SyncNewBlock(header *types.Header, c *Controller) {
	// TODO: implementation needed
}