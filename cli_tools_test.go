package btccli

import (
	"fmt"
	"github.com/lomocoin/btccli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	killBitcoind, err := BitcoindRegtest()
	testtool.FailOnErr(t, err)
	defer killBitcoind()

	addrs, err := CliToolGetSomeAddrs(4)
	testtool.FailOnErr(t, err)
	fmt.Println(ToJsonIndent(addrs))
}
