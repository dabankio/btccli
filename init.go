package btccli

import (
	"fmt"
	"os"
)

// bitcoin bin path
var (
	BasePath      = "/Users/some_user/Applications/bitcoin/bin" //see init()
	CmdBitcoind   = BasePath + "/bitcoind"
	CmdBitcoinCli = BasePath + "/bitcoin-cli"
)

func init() {
	p := os.Getenv(BitcoinBinPathEnv)
	if p == "" {
		panic("using bitcoind need bin path env: BITCOIN_BIN_PATH")
	}
	BasePath = p
	// TODO check exists
	fmt.Println("bitcoin bin path:", BasePath)
	CmdBitcoind = BasePath + "/bitcoind"      //windows may need change suffix
	CmdBitcoinCli = BasePath + "/bitcoin-cli" //windows may need change suffix

}
