package database

import (
	"log"
	"strings"

	"github.com/nlsh710599/go-practice/internal/database/model"

	"gorm.io/gorm"
)

type RDS interface {
	CreateTable() error
	GetLatestNBlocks(n uint64) ([]*BlockWithoutTransaction, error)
	GetBlockByNumber(n uint64) (*GetBlockByNumberResp, error)
	GetTransactionByHash(txHash string) (*GetTransactionByHashResp, error)
	GetOldestConfirmedBlockNumber() (uint64, error)
	GetLatestConfirmedBlockNumber() (uint64, error)
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

func (pg *postgresClient) GetOldestConfirmedBlockNumber() (uint64, error) {
	var oldestConfirmedBlockNumber uint64
	if err := pg.client.Table("blocks").Select("min(number)").Scan(&oldestConfirmedBlockNumber).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint64 is unsupported") {
			return 0, nil
		}
		return 0, err
	}
	return oldestConfirmedBlockNumber, nil
}

func (pg *postgresClient) GetLatestConfirmedBlockNumber() (uint64, error) {
	var latestConfirmedBlockNumber uint64
	if err := pg.client.Table("blocks").Select("max(number)").Scan(&latestConfirmedBlockNumber).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint64 is unsupported") {
			return 0, nil
		}
		return 0, err
	}
	return latestConfirmedBlockNumber, nil
}
