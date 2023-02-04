package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nlsh710599/go-practice/internal/config"
	"github.com/nlsh710599/go-practice/internal/database"
	"github.com/nlsh710599/go-practice/internal/syncer"
	"github.com/nlsh710599/go-practice/internal/web3"
)

func signalHandler(sigChan chan os.Signal, aborter chan<- bool) {
	for sig := range sigChan {
		log.Printf("receive signal: %d\n", sig)
		close(aborter)
	}

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

	c := &syncer.Controller{
		RDS:         rds,
		Web3Service: web3rpc,
	}

	oldestConfirmedBlock, err := c.RDS.GetOldestConfirmedBlockNumber()
	if err != nil {
		log.Panicf("Failed to get oldest confirmed block in db: %v", err)
	}

	latestConfirmedBlock, err := c.RDS.GetLatestConfirmedBlockNumber()
	if err != nil {
		log.Panicf("Failed to get latest confirmed block in db: %v", err)
	}

	signalChan := make(chan os.Signal, 1)
	aborter := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go signalHandler(signalChan, aborter)

	wg := sync.WaitGroup{}

	if oldestConfirmedBlock != 0 {
		wg.Add(1)
		go func() {
			syncer.SyncConformedBlockBackward(0, oldestConfirmedBlock-1, c, aborter)
			wg.Done()
		}()
	}

	headers := make(chan *types.Header)
	startSyncConformedBlockForward := false

	sub, err := web3ws.GetClient().SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case header := <-headers:
				if !startSyncConformedBlockForward {
					wg.Add(1)
					go func() {
						if latestConfirmedBlock == 0 {
							syncer.SyncConformedBlockBackward(latestConfirmedBlock, header.Number.Uint64()-uint64(config.Get().ConfirmationBlockCount), c, aborter)
						} else {
							syncer.SyncConformedBlockForward(latestConfirmedBlock+1, header.Number.Uint64()-uint64(config.Get().ConfirmationBlockCount), c, aborter)
						}
						wg.Done()
					}()
					startSyncConformedBlockForward = true
				}
				wg.Add(1)
				go func() {
					syncer.SyncNewBlock(header, c, aborter)
					wg.Done()
				}()
			case <-aborter:
				log.Println("handler finished")
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}
