package constant

const (
	DataShard     = 4
	ParityShards  = 2
	AllShards     = DataShard + ParityShards
	BlockPerShard = 8000
	BlockSize     = BlockPerShard * DataShard
)

