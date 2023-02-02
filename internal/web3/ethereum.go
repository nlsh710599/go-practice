package web3

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type web3Client struct {
	client  *ethclient.Client
	chainId *big.Int
}

type Web3 interface {
	Close()
	GetClient() *ethclient.Client
	GetBlockNumber() (uint64, error)
	GetBlockByNumber(blockNumber *big.Int) (*GetBlockByNumberResp, error)
	GetTransactionReceipt(hash string) (*types.Receipt, error)
}

func (wc *web3Client) Close() {
	wc.client.Close()
}

func (wc *web3Client) GetClient() *ethclient.Client {
	return wc.client
}

func (wc *web3Client) GetBlockNumber() (uint64, error) {
	return wc.client.BlockNumber(context.Background())
}

func (wc *web3Client) GetBlockByNumber(blockNumber *big.Int) (*GetBlockByNumberResp, error) {

	block, err := wc.client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		return nil, err
	}

	txs := make([]string, len(block.Transactions()))
	for i, tx := range block.Transactions() {
		txs[i] = tx.Hash().Hex()
	}

	res := &GetBlockByNumberResp{
		Number:       block.NumberU64(),
		Hash:         block.Hash().Hex(),
		Timestamp:    block.Time(),
		ParentHash:   block.ParentHash().Hex(),
		Transactions: txs,
	}

	return res, nil
}

func (wc *web3Client) GetTransactionReceipt(hash string) (*types.Receipt, error) {
	return wc.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
}

func DialEth(url string) (*web3Client, error) {
	wc := &web3Client{}
	var err error

	wc.client, err = ethclient.Dial(url)

	if err != nil {
		log.Printf("Failed to dial eth client: %v", err)
		return nil, err
	}

	wc.chainId, err = wc.client.ChainID(context.Background())

	if err != nil {
		log.Printf("Failed to get chain id: %v", err)
		return nil, err
	}

	return wc, nil
}

func New(url string) (Web3, error) {
	instance, err := DialEth(url)

	if err != nil {
		log.Printf("Failed to dial eth: %v", err)
		return nil, err
	}

	return instance, nil
}
