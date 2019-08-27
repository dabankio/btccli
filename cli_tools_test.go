package btccli

import (
	"fmt"
	"github.com/lemon-sunxiansong/btccli/testtool"
	"testing"
)

func TestCliToolGetSomeAddrs(t *testing.T) {
	cc, err := BitcoindRegtest()
	testtool.FailOnErr(t, err)
	defer func() {
		cc <- struct{}{}
	}()

	addrs, err := CliToolGetSomeAddrs(4)
	testtool.FailOnErr(t, err)
	fmt.Println(ToJsonIndent(addrs))
}
