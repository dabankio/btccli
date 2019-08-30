package btccli

import (
	"fmt"
	"testing"

	"github.com/lomocoin/btccli/testtool"
)

func TestBitcoindWithOptions(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{
		NewTmpDir: true,
	})
	testtool.FailOnErr(t, err, "bitcoind start err")
	defer func() {
		killBitcoind()
		testtool.FailOnFlag(t, cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "bitcoind should be stopped")
		t.Log("Done")
	}()

	fmt.Println("================to_get_balance=======")

	bal, err := cli.Getbalance(nil, nil, nil)
	testtool.FailOnErr(t, err, "")
	fmt.Println("bal", bal)

	testtool.FailOnFlag(t, !cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "端口现在应该已经运行")

	_, _, err = RunBitcoind(nil)
	testtool.FailOnFlag(t, err == nil, "再次运行应该返回错误")
}
