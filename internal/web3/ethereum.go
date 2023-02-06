package web3

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nlsh710599/go-practice/internal/database/model"
)

type web3Client struct {
	client  *ethclient.Client
	chainId *big.Int
}

type Web3 interface {
	Close()
	GetClient() *ethclient.Client
	GetBlockByNumber(*big.Int) (*model.Block, error)
	GetTransactionDetail(string) (model.Transaction, error)
	GetTransactionReceipt(string) (*types.Receipt, error)
	GetTransactionByHash(string) (*types.Transaction, bool, error)
}

func (wc *web3Client) Close() {
	wc.client.Close()
}

func (wc *web3Client) GetClient() *ethclient.Client {
	return wc.client
}

func (wc *web3Client) GetBlockByNumber(blockNumber *big.Int) (*model.Block, error) {

	block, err := wc.client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		return nil, err
	}
	txs := make([]model.Transaction, len(block.Transactions()))
	for i, tx := range block.Transactions() {
		txs[i], err = wc.GetTransactionDetail(tx.Hash().Hex())
		if err != nil {
			return nil, err
		}
	}

	return &model.Block{
		Number:       block.NumberU64(),
		Hash:         block.Hash().Hex(),
		Timestamp:    block.Time(),
		ParentHash:   block.ParentHash().Hex(),
		Transactions: txs,
	}, nil
}

func (wc *web3Client) GetTransactionDetail(hash string) (model.Transaction, error) {
	txInfo, _, err := wc.GetTransactionByHash(hash)
	if err != nil {
		return model.Transaction{}, err
	}
	txReceipt, err := wc.GetTransactionReceipt(hash)
	if err != nil {
		return model.Transaction{}, err
	}

	logs := make([]model.Log, len(txReceipt.Logs))
	for i, log := range txReceipt.Logs {
		logs[i] = model.Log{
			Index:           uint64(log.Index),
			Data:            hex.EncodeToString(log.Data),
			TransactionHash: hash,
		}

	}

	from, err := types.Sender(types.LatestSignerForChainID(txInfo.ChainId()), txInfo)
	if err != nil {
		log.Panicf("Failed to get transaction sender : %v", err)
	}

	var to string
	if txInfo.To() == nil {
		to = "0x0000000000000000000000000000000000000000"
	} else {
		to = txInfo.To().Hex()
	}

	return model.Transaction{
		Hash:        txInfo.Hash().Hex(),
		From:        from.Hex(),
		To:          to,
		Nonce:       txInfo.Nonce(),
		Data:        hex.EncodeToString(txInfo.Data()),
		Value:       txInfo.Value().String(),
		BlockNumber: 3,
		Logs:        logs,
	}, nil

}

func (wc *web3Client) GetTransactionReceipt(hash string) (*types.Receipt, error) {
	return wc.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
}

func (wc *web3Client) GetTransactionByHash(hash string) (*types.Transaction, bool, error) {
	return wc.client.TransactionByHash(context.Background(), common.HexToHash(hash))
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
