package database

import (
	"log"
	"strings"

	"github.com/nlsh710599/go-practice/internal/database/model"

	"gorm.io/gorm"
)

type RDS interface {
	GetClient() *gorm.DB
	CreateTable() error
	GetLatestNBlocks(uint64) ([]*BlockWithoutTransaction, error)
	GetBlockByNumber(uint64) (*GetBlockByNumberResp, error)
	GetTransactionByHash(string) (*GetTransactionByHashResp, error)
	GetOldestConfirmedBlockNumber() (uint64, error)
	GetLatestConfirmedBlockNumber() (uint64, error)
	InsertBlock(*model.Block) error
	UpdateBlock(uint64, bool) error
	InsertTransaction(*model.Transaction) error
	InsertLog(*model.Log) error
}

type postgresClient struct {
	client *gorm.DB
}

func (pg *postgresClient) GetClient() *gorm.DB {
	return pg.client
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

func (pg *postgresClient) UpdateBlock(number uint64, isConfirmed bool) error {
	return pg.client.Model(&model.Block{}).
		Where("number = ?", number).
		Update("is_confirmed", isConfirmed).Error
}

func (pg *postgresClient) InsertBlock(blockInfo *model.Block) error {
	return pg.client.Create(&blockInfo).Error
}

func (pg *postgresClient) InsertTransaction(transactionInfo *model.Transaction) error {
	return pg.client.Create(&transactionInfo).Error
}

func (pg *postgresClient) InsertLog(LogInfo *model.Log) error {
	return pg.client.Create(&LogInfo).Error
}
