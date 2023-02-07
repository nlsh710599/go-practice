package database

import (
	"log"
	"strings"

	"github.com/nlsh710599/go-practice/internal/database/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var res []*BlockWithoutTransaction
	if err := pg.client.Model(&model.Block{}).Order("number desc").Limit(int(n)).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (pg *postgresClient) GetBlockByNumber(n uint64) (*GetBlockByNumberResp, error) {
	var tmp *model.Block
	var res *GetBlockByNumberResp
	if err := pg.client.Model(&model.Block{}).Preload("Transactions", func(db *gorm.DB) *gorm.DB {
		return db.Select("hash, block_number")
	}).Where("number = ?", n).Find(&tmp).Error; err != nil {
		return nil, err
	}
	txHashList := make([]string, len(tmp.Transactions))
	for i := range txHashList {
		txHashList[i] = tmp.Transactions[i].Hash
	}
	res = &GetBlockByNumberResp{
		Number:       tmp.Number,
		Hash:         tmp.Hash,
		Timestamp:    tmp.Timestamp,
		ParentHash:   tmp.ParentHash,
		Transactions: txHashList,
	}

	return res, nil
}

func (pg *postgresClient) GetTransactionByHash(txHash string) (*GetTransactionByHashResp, error) {
	var tmp *model.Transaction
	var res *GetTransactionByHashResp
	if err := pg.client.Model(&model.Transaction{}).Preload("Logs").Where("hash = ?", txHash).Find(&tmp).Error; err != nil {
		return nil, err
	}
	logList := make([]Log, len(tmp.Logs))
	for i := range logList {
		logList[i].Data = tmp.Logs[i].Data
		logList[i].Index = tmp.Logs[i].Index
	}
	res = &GetTransactionByHashResp{
		Hash:  tmp.Hash,
		From:  tmp.From,
		To:    tmp.To,
		Nonce: tmp.Nonce,
		Data:  tmp.Data,
		Value: tmp.Value,
		Logs:  logList,
	}
	return res, nil
}

func (pg *postgresClient) GetOldestConfirmedBlockNumber() (uint64, error) {
	var oldestConfirmedBlockNumber uint64
	if err := pg.client.Table("blocks").Select("min(number)").Where("is_confirmed = ?", true).Scan(&oldestConfirmedBlockNumber).Error; err != nil {
		if strings.Contains(err.Error(), "converting NULL to uint64 is unsupported") {
			return 0, nil
		}
		return 0, err
	}
	return oldestConfirmedBlockNumber, nil
}

func (pg *postgresClient) GetLatestConfirmedBlockNumber() (uint64, error) {
	var latestConfirmedBlockNumber uint64
	if err := pg.client.Table("blocks").Select("max(number)").Where("is_confirmed = ?", true).Scan(&latestConfirmedBlockNumber).Error; err != nil {
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
	return pg.client.Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "number"}},
			UpdateAll: true,
		},
	).Create(&blockInfo).Error
}
