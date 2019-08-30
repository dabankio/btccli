package btccli

// 固定常量
const (
	RPCPortRegtest = 18443

	BitcoinBinPathEnv = "BITCOIN_BIN_PATH"

	CmdParamRegtest = "-regtest"
)

// some net id
const (
	NetRegtest = iota + 1
	NetTestnet3
	NetMainnet
)
