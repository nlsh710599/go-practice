package syncer

import (
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/database/model"
	"github.com/nlsh710599/go-practice/internal/web3"
	"gorm.io/gorm"
)

type Controller struct {
	RDS         database.RDS
	Web3Service web3.Web3
}

func SyncConformedBlockBackward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := to; i > from; i-- {
		syncConformedBlock(i, c, aborter)
	}
}

func SyncConformedBlockForward(from uint64, to uint64, c *Controller, aborter <-chan bool) {
	for i := from; i < to; i++ {
		syncConformedBlock(i, c, aborter)
	}
}

func SyncNewBlock(header *types.Header, c *Controller, aborter <-chan bool) {
	// TODO: implementation needed
}

func syncConformedBlock(blockNumber uint64, c *Controller, aborter <-chan bool) {
	log.Println("I'm going to sync block No.", blockNumber)
	blockInfo, err := c.Web3Service.GetBlockByNumber(big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Panicf("Failed to get block by number : %v", err)
	}

	err = c.RDS.GetClient().Transaction(func(tx *gorm.DB) error {
		err = c.RDS.InsertBlock(&model.Block{
			Number:      blockInfo.Number,
			Hash:        blockInfo.Hash,
			Timestamp:   blockInfo.Timestamp,
			ParentHash:  blockInfo.ParentHash,
			IsConfirmed: true,
		})
		if err != nil {
			log.Panicf("Failed to insert block : %v", err)
		}

		for _, tx := range blockInfo.Transactions {
			go syncTx(blockNumber, tx, c)
		}
		return nil
	})

	if err != nil {
		log.Panicf("Failed to sync confirmed block : %v", err)
	}
}

func syncTx(blockNumber uint64, tx string, c *Controller) {
	transactionReceipt, err := c.Web3Service.GetTransactionReceipt(tx)
	if err != nil {
		log.Panicf("Failed to get transaction receipt : %v", err)
	}

	transactionInfo, _, err := c.Web3Service.GetTransactionByHash(tx)
	if err != nil {
		log.Panicf("Failed to get transaction by hash : %v", err)
	}
	from, err := types.Sender(types.LatestSignerForChainID(transactionInfo.ChainId()), transactionInfo)
	if err != nil {
		log.Panicf("Failed to get transaction sender : %v", err)
	}
	err = c.RDS.InsertTransaction(&model.Transaction{
		Hash:        tx,
		From:        from.Hex(),
		To:          transactionInfo.To().Hex(),
		Nonce:       uint64(transactionInfo.Nonce()),
		Data:        hex.EncodeToString(transactionInfo.Data()),
		Value:       transactionInfo.Value().Uint64(),
		BlockNumber: blockNumber,
	})

	if err != nil {
		log.Panicf("Failed to insert transaction : %v", err)
	}

	for _, Log := range transactionReceipt.Logs {
		go syncLog(tx, Log, c)
	}
}

func syncLog(tx string, Log *types.Log, c *Controller) {
	err := c.RDS.InsertLog(&model.Log{
		Index:           uint64(Log.Index),
		Data:            hex.EncodeToString(Log.Data),
		TransactionHash: tx,
	})
	if err != nil {
		log.Panicf("Failed to insert log : %v", err)
	}
}
