package btccli

import (
	"fmt"
)

// Addr .
type Addr struct {
	Address string
	Privkey string
	Pubkey  string
}

func (ad *Addr) String() string {
	return fmt.Sprintf("{Address: \"%s\", Privkey: \"%s\", Pubkey: \"%s\"}", ad.Address, ad.Privkey, ad.Pubkey)
}

// CliToolGetSomeAddrs 一次获取n个地址（包含pub-priv key)
func (cli *Cli) ToolGetSomeAddrs(n int) ([]Addr, error) {
	var addrs []Addr
	for i := 0; i < n; i++ {
		add, err := cli.Getnewaddress(nil, nil)
		if err != nil {
			return nil, err
		}
		info, err := cli.GetAddressInfo(add)
		if err != nil {
			return nil, err
		}
		dump, err := cli.Dumpprivkey(add)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, Addr{
			Address: add, Privkey: dump, Pubkey: info.Pubkey,
		})
	}
	return addrs, nil
}
