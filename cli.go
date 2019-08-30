package btccli

import (
	"fmt"
)

// Cli .
type Cli struct {
	args         []string
	IgnoreCliCmd bool //not print cmd
	IgnoreCliOut bool //not print cli out
}

// RunningBitcoindOptions .
type RunningBitcoindOptions struct {
	RPCPort     uint
	RPCUser     string
	RPCPassword string
	DataDir     string
	NetID       uint32
}

// NewCliFromRunningBitcoind .
func NewCliFromRunningBitcoind(options RunningBitcoindOptions) (*Cli, error) {
	args := []string{}
	if options.RPCPort != 0 {
		args = append(args, fmt.Sprintf("-rpcport=%d", options.RPCPort))
	}
	if options.RPCUser != "" {
		args = append(args, fmt.Sprintf("-rpcuser=%s", options.RPCUser))
	}
	if options.RPCPassword != "" {
		args = append(args, fmt.Sprintf("-rpcpassword=%s", options.RPCPassword))
	}
	if options.DataDir != "" {
		args = append(args, fmt.Sprintf("-datadir=%s", options.DataDir))
	}

	switch options.NetID {
	case NetRegtest:
		args = append(args, "-regtest")
	case NetTestnet3:
		args = append(args, "-testnet")
	case NetMainnet:
	case 0:
	default:
		return nil, fmt.Errorf("unknown net %d", options.NetID)
	}
	return &Cli{args: args}, nil
}

// AppendArgs .
func (cli *Cli) AppendArgs(args ...string) []string {
	return append(cli.args, args...)
}
