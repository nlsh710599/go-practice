package model

type Block struct {
	Number      uint64 `json:"block_num" gorm:"primarykey"`
	Hash        string `json:"block_hash"`
	Timestamp   uint64 `json:"block_time"`
	ParentHash  string `json:"parent_hash"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type Transaction struct {
	Hash        string `json:"tx_hash" gorm:"primarykey"`
	From        string `json:"from"`
	To          string `json:"to"`
	Nonce       uint64 `json:"nonce"`
	Data        string `json:"data"`
	Value       uint64 `json:"value"`
	BlockNumber uint64 `json:"block_number"`
	Block       Block
}

type Log struct {
	ID              uint64 `json:"id" gorm:"primarykey"`
	Index           uint64 `json:"index"`
	Data            string `json:"data"`
	TransactionHash string `json:"tx_hash"`
	Transaction     Transaction
}
