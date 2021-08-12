package rs

const (
	DataShard     = 4
	ParityShard   = 2
	AllShard      = DataShard + ParityShard
	BlockPerShard = 8000
	BlockSie      = BlockPerShard * DataShard
)
