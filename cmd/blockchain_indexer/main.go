package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlsh710599/go-practice/internal/config"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/web3"
)

type Controller struct {
	RDS         database.RDS
	web3Service web3.Web3
}

func main() {

	rds, err := database.New(config.Get().PostgresHost, config.Get().PostgresUser, config.Get().PostgresPassword,
		config.Get().PostgresDatabase, config.Get().PostgresPort)
	if err != nil {
		log.Panicf("Failed to initialize RDS: %v", err)
	}

	web3rpc, err := web3.New(config.Get().RpcUrl)
	if err != nil {
		log.Panicf("Failed to create web3 instance: %v", err)
	}

	web3ws, err := web3.New(config.Get().WsRpcUrl)
	if err != nil {
		log.Panicf("Failed to create web3 instance: %v", err)
	}

	c := &Controller{
		RDS:         rds,
		web3Service: web3rpc,
	}

	oldestConfirmedBlock, err := c.RDS.GetOldestConfirmedBlock()
	if err != nil {
		log.Panicf("Failed to get oldest confirmed block in db: %v", err)
	}

	latestConfirmedBlock, err := c.RDS.GetLatestConfirmedBlock()
	if err != nil {
		log.Panicf("Failed to get oldest confirmed block in db: %v", err)
	}

	go syncBlockBackward(1, oldestConfirmedBlock, c)

	headers := make(chan *types.Header)
	startSyncBlockForward := false

	sub, err := web3ws.GetClient().SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			if !startSyncBlockForward {
				go syncBlockForward(latestConfirmedBlock, header.Number.Uint64()-uint64(config.Get().ConfirmationBlockCount)-1, c)
				startSyncBlockForward = true
			}
			go syncBlockListened(header, c)
		}
	}

}

func syncBlockBackward(from uint64, to uint64, c *Controller) {
	// TODO: implementation needed
}

func syncBlockForward(from uint64, to uint64, c *Controller) {
	// TODO: implementation needed
}

func syncBlockListened(header *types.Header, c *Controller) {
	// TODO: implementation needed
}
