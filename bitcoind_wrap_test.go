package btccli

import (
	"testing"
	"time"

	"github.com/lomocoin/btccli/testtool"
)

func TestBitcoindRegtest(t *testing.T) {
	killBitcoind, err := BitcoindRegtest()
	testtool.FailOnErr(t, err, "bitcoind start err")
	defer func() {
		killBitcoind()
		time.Sleep(time.Second)
		testtool.FailOnFlag(t, cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "bitcoind should be stopped")
		t.Log("Done")
	}()

	testtool.FailOnFlag(t, !cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "端口现在应该已经运行")

	_, err = BitcoindRegtest()
	testtool.FailOnFlag(t, err == nil, "再次运行应该返回错误")
}

func TestBitcoindWithOptions(t *testing.T) {
	killBitcoind, err := BitcoindWithOptions(StartOptions{NewTmpDir: true})
	testtool.FailOnErr(t, err, "bitcoind start err")
	defer func() {
		killBitcoind()
		testtool.FailOnFlag(t, cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "bitcoind should be stopped")
		t.Log("Done")
	}()

	testtool.FailOnFlag(t, !cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin"), "端口现在应该已经运行")

	_, err = BitcoindRegtest()
	testtool.FailOnFlag(t, err == nil, "再次运行应该返回错误")
}
