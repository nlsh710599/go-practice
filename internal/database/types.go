package database

type BlockWithoutTransaction struct {
	Number      uint64
	Hash        string
	Timestamp   uint64
	ParentHash  string
	IsConfirmed bool
}

type GetBlockByNumberResp struct {
	Number       uint64
	Hash         string
	Timestamp    uint64
	ParentHash   string
	Transactions []string
}

type GetTransactionByHashResp struct {
	Hash  string
	From  string
	To    string
	Nonce uint64
	Data  string
	Value uint64
	Logs  []Log
}

type Log struct {
	Index uint64
	Data  string
}
