package database

import (
	"log"

	"github.com/nlsh710599/go-practice/internal/database/model"

	"gorm.io/gorm"
)

type RDS interface {
	CreateTable() error
	GetLatestNBlocks(n uint64) ([]*BlockWithoutTransaction, error)
	GetBlockByNumber(n uint64) (*GetBlockByNumberResp, error)
	GetTransactionByHash(txHash string) (*GetTransactionByHashResp, error)
	GetOldestConfirmedBlock() (uint64, error)
	GetLatestConfirmedBlock() (uint64, error)
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

func (pg *postgresClient) GetLatestNBlocks(n uint64) ([]*BlockWithoutTransaction, error) {
	// TODO: implementation needed
	return make([]*BlockWithoutTransaction, n), nil
}

func (pg *postgresClient) GetBlockByNumber(n uint64) (*GetBlockByNumberResp, error) {
	// TODO: implementation needed
	return &GetBlockByNumberResp{}, nil
}

func (pg *postgresClient) GetTransactionByHash(txHash string) (*GetTransactionByHashResp, error) {
	// TODO: implementation needed
	return &GetTransactionByHashResp{}, nil
}

func (pg *postgresClient) GetOldestConfirmedBlock() (uint64, error) {
	// TODO: implementation needed
	return 0, nil
}

func (pg *postgresClient) GetLatestConfirmedBlock() (uint64, error) {
	// TODO: implementation needed
	return 0, nil
}
