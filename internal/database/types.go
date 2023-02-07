package database

type BlockWithoutTransaction struct {
	Number      uint64 `json:"block_num"`
	Hash        string `json:"block_hash"`
	Timestamp   uint64 `json:"block_time"`
	ParentHash  string `json:"parent_hash"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type GetBlockByNumberResp struct {
	Number       uint64   `json:"block_num"`
	Hash         string   `json:"block_hash"`
	Timestamp    uint64   `json:"block_time"`
	ParentHash   string   `json:"parent_hash"`
	Transactions []string `json:"transactions"`
}

type GetTransactionByHashResp struct {
	Hash  string `json:"tx_hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Nonce uint64 `json:"nence"`
	Data  string `json:"data"`
	Value string `json:"value"`
	Logs  []Log  `json:"logs"`
}

type Log struct {
	Index uint64 `json:"index"`
	Data  string `json:"data"`
}
