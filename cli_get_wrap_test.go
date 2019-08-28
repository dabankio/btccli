package btccli

import (
	"fmt"
	"github.com/lomocoin/btccli/testtool"
	"testing"
)

func TestCliGetAddressInfo(t *testing.T) {
	killBitcoind, err := BitcoindRegtest()
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killBitcoind()

	var addrs []Addr

	for i := 0; i < 5; i++ {
		var newAddr string
		add := Addr{}
		{
			newAddr, err = CliGetnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			add.Address = newAddr
		}
		{
			addrInfo, err := CliGetAddressInfo(newAddr)
			testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
			fmt.Println("address info", ToJsonIndent(addrInfo))
			add.Pubkey = addrInfo.Pubkey
		}
		{
			vRes, err := CliValidateaddress(newAddr)
			testtool.FailOnFlag(t, err != nil, "Failed to validate address", err)
			fmt.Println("validate address res:", ToJsonIndent(vRes))
		}

		{
			prvk, err := CliDumpprivkey(newAddr)
			testtool.FailOnErr(t, err, "")
			add.Privkey = prvk
		}
		addrs = append(addrs, add)
	}
	fmt.Printf("%#v", addrs)
}

func TestCliGetbalance(t *testing.T) {
	killbitcoind, err := BitcoindRegtest()
	testtool.FailOnErr(t, err, "start failed")
	defer killbitcoind()

	newAddr, err := CliGetnewaddress(nil, nil)
	testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)

	_, err = CliGeneratetoaddress(101, newAddr, nil)
	testtool.FailOnErr(t, err, "failed to gen to add")

	bal, er := CliGetbalance(nil, nil, nil)
	testtool.FailOnErr(t, er, "Failed to get balance")
	testtool.FailOnFlag(t, bal == 0, "Failed to get balance")
	fmt.Println("balance: ", bal)
}
