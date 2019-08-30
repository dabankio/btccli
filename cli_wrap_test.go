package btccli

import (
	"fmt"
	"testing"

	"github.com/lomocoin/btccli/testtool"

	"github.com/lomocoin/btccli/btcjson"
)

func TestCliCreatemultisig(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killBitcoind()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [3]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := cli.Getnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			info, err := cli.GetAddressInfo(add)
			testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
			addrs[i].pubkey = info.Pubkey

			privkey, err := cli.Dumpprivkey(add)
			testtool.FailOnFlag(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
			// keys = append(keys, info.addr)
		}
		multisigResp, err = cli.Createmultisig(2, keys, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to create multi sig", err)
		fmt.Println("keys", keys)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		info, err := cli.GetAddressInfo(multisigResp.Address)
		testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
		fmt.Println("multisigAddres info", ToJSONIndent(info))
	}
	{
		vRes, err := cli.Validateaddress(multisigResp.Address)
		testtool.FailOnFlag(t, err != nil, "Failed to validate address info", err)
		fmt.Println("validate multisig address", ToJSONIndent(vRes))
	}

}

func TestCliAddmultisigaddress(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killBitcoind()

	type addrinfo struct {
		addr, privkey, pubkey string
	}
	var addrs [5]addrinfo
	{ //获取几个新地址
		for i := 0; i < len(addrs); i++ {
			add, err := cli.Getnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			addrs[i] = addrinfo{addr: add}

			info, err := cli.GetAddressInfo(add)
			testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
			addrs[i].pubkey = info.Pubkey

			privkey, err := cli.Dumpprivkey(add)
			testtool.FailOnFlag(t, err != nil, "Failed to dump privkey", err)
			addrs[i].privkey = privkey
		}
		fmt.Println("addrs", addrs)
	}

	var multisigResp btcjson.CreateMultiSigResult
	{ //create multisig address
		var keys []string
		for _, info := range addrs {
			keys = append(keys, info.pubkey)
		}

		multisigResp, err = cli.Addmultisigaddress(btcjson.AddMultisigAddressCmd{
			NRequired: 3, Keys: keys,
		})
		testtool.FailOnFlag(t, err != nil, "Failed to add multi sig address", err)
		fmt.Println("multisig address:", jsonStr(multisigResp))
	}

	{
		info, err := cli.GetAddressInfo(multisigResp.Address)
		testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
		fmt.Println("multisigAddres info", ToJSONIndent(info))
	}
}

func TestCliGetAddressInfo(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	// killBitcoind, err := BitcoindRegtest("-addresstype=legacy")
	testtool.FailOnFlag(t, err != nil, "Failed to start btcd", err)
	defer killBitcoind()

	var addrs []Addr

	for i := 0; i < 5; i++ {
		var newAddr string
		add := Addr{}
		{
			newAddr, err = cli.Getnewaddress(nil, nil)
			testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
			add.Address = newAddr
		}
		{
			addrInfo, err := cli.GetAddressInfo(newAddr)
			testtool.FailOnFlag(t, err != nil, "Failed to get address info", err)
			fmt.Println("address info", ToJSONIndent(addrInfo))
			add.Pubkey = addrInfo.Pubkey
		}
		{
			vRes, err := cli.Validateaddress(newAddr)
			testtool.FailOnFlag(t, err != nil, "Failed to validate address", err)
			fmt.Println("validate address res:", ToJSONIndent(vRes))
		}

		{
			prvk, err := cli.Dumpprivkey(newAddr)
			testtool.FailOnErr(t, err, "")
			add.Privkey = prvk
		}
		addrs = append(addrs, add)
	}
	fmt.Printf("%#v", addrs)
}

func TestCliGetbalance(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	testtool.FailOnErr(t, err, "start failed")
	defer killBitcoind()

	newAddr, err := cli.Getnewaddress(nil, nil)
	testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)

	_, err = cli.Generatetoaddress(101, newAddr, nil)
	testtool.FailOnErr(t, err, "failed to gen to add")

	bal, er := cli.Getbalance(nil, nil, nil)
	testtool.FailOnErr(t, er, "Failed to get balance")
	testtool.FailOnFlag(t, bal == 0, "Failed to get balance")
	fmt.Println("balance: ", bal)
}

func TestCliListunspent(t *testing.T) {
	cli, killBitcoind, err := RunBitcoind(&RunOptions{NewTmpDir: true})
	testtool.FailOnFlag(t, err != nil, "Failed to start d", err)
	defer killBitcoind()

	var newaddr string
	{
		newaddr, err = cli.Getnewaddress(nil, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to get new address", err)
	}
	{
		const leng = 102
		hashs, err := cli.Generatetoaddress(leng, newaddr, nil)
		testtool.FailOnFlag(t, err != nil, "Failed to gen to addr", err)
		testtool.FailOnFlag(t, len(hashs) != leng, "len not equal", leng, hashs)
	}
	{
		unspents, err := cli.Listunspent(0, 999, []string{newaddr}, btcjson.Bool(true), nil)
		testtool.FailOnFlag(t, err != nil, "Fail on listunspent", err)
		fmt.Println(jsonStr(unspents))
	}
}
