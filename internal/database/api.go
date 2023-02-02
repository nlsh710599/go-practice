package database

import (
	"log"

	"github.com/nlsh710599/go-practice/internal/database/model"
	"github.com/nlsh710599/go-practice/internal/web3"

	"gorm.io/gorm"
)

type RDS interface {
	CreateTable() error
	GetLatestNBlocks(n uint64) ([]*web3.BlockWithoutTransaction, error)
	GetBlockByNumber(n uint64) (*web3.GetBlockByNumberResp, error)
	GetTransactionByHash(txHash string) (*web3.GetTransactionByHashResp, error)
}

type postgresClient struct {
	client *gorm.DB
}

func (pg *postgresClient) CreateTable() error {
	if err := pg.client.Migrator().AutoMigrate(
		&model.Block{},
		&model.Transaction{},
		&model.Log{},
	); err != nil {
		log.Printf("Failed to migrate tables: %v", err)
		return err
	}

	return nil
}

func (pg *postgresClient) GetLatestNBlocks(n uint64) ([]*web3.BlockWithoutTransaction, error) {
	// TODO: implementation needed
	return make([]*web3.BlockWithoutTransaction, n), nil
}

func (pg *postgresClient) GetBlockByNumber(n uint64) (*web3.GetBlockByNumberResp, error) {
	// TODO: implementation needed
	return &web3.GetBlockByNumberResp{}, nil
}

func (pg *postgresClient) GetTransactionByHash(txHash string) (*web3.GetTransactionByHashResp, error) {
	// TODO: implementation needed
	return &web3.GetTransactionByHashResp{}, nil
}
