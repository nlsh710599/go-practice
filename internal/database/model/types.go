package model

type Block struct {
	Number      int64  `json:"block_num" gorm:"primarykey"`
	Hash        string `json:"block_hash"`
	Timestamp   int64  `json:"block_time"`
	ParentHash  string `json:"parent_hash"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type Transaction struct {
	Hash        string `json:"tx_hash" gorm:"primarykey"`
	From        string `json:"from"`
	To          string `json:"to"`
	Nonce       string `json:"nonce"`
	Data        string `json:"data"`
	Value       string `json:"value"`
	BlockNumber int64  `json:"block_number"`
	Block       Block
}

type Log struct {
	ID              int64  `json:"id" gorm:"primarykey"`
	Index           int64  `json:"index"`
	Data            string `json:"data"`
	TransactionHash string `json:"tx_hash"`
	Transaction     Transaction
}
