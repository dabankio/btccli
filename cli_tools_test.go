package btccli

import (
	"fmt"
	"github.com/lomocoin/btccli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	testtool.FailOnErr(t, err)
	defer killBitcoind()

	addrs, err := cli.ToolGetSomeAddrs(4)
	testtool.FailOnErr(t, err)
	fmt.Println(ToJSONIndent(addrs))
}
