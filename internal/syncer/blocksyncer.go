package syncer

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlsh710599/go-practice/internal/config"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/web3"
)

type Controller struct {
	RDS         database.RDS
	Web3Service web3.Web3
}

func SyncConfirmedBlockBackward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := to; i > from; i-- {
		select {
		case <-aborter:
			return
		default:
			syncConfirmedBlock(i, c, aborter)
		}
	}
}

func SyncConfirmedBlockForward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := from; i < to; i++ {
		select {
		case <-aborter:
			return
		default:
			syncConfirmedBlock(i, c, aborter)
		}
	}
}

func SyncNewBlock(header *types.Header, c *Controller, aborter <-chan bool) {
	select {
	case <-aborter:
		return
	default:
		syncBlock(header.Number.Uint64(), false, c, aborter)
		if header.Number.Uint64() > uint64(config.Get().ConfirmationBlockCount) {
			syncConfirmedBlock(header.Number.Uint64()-uint64(config.Get().ConfirmationBlockCount), c, aborter)
		}
	}

}

func syncConfirmedBlock(blockNumber uint64, c *Controller, aborter <-chan bool) {
	select {
	case <-aborter:
		return
	default:
		syncBlock(blockNumber, true, c, aborter)
		err := c.RDS.UpdateBlock(blockNumber, true)
		if err != nil {
			log.Panicf("Failed to insert block : %v", err)
		}
	}
}

func syncBlock(blockNumber uint64, isConfirmed bool, c *Controller, aborter <-chan bool) {
	select {
	case <-aborter:
		return
	default:
		log.Println("I'm going to sync block No.", blockNumber)
		blockInfo, err := c.Web3Service.GetBlockByNumber(big.NewInt(int64(blockNumber)))
		if err != nil {
			syncBlock(blockNumber, isConfirmed, c, aborter)
			return
		}
		err = c.RDS.InsertBlock(blockInfo)
		if err != nil {
			syncBlock(blockNumber, isConfirmed, c, aborter)
			return
		}
	}

}
