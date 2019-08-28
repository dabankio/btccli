package btccli

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lomocoin/btccli/btcjson"
)

// CliGetbestblockhash .
func CliGetbestblockhash() (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getbestblockhash",
	))
	//TODO validate hash
	return cmdPrint, nil
}

// CliGetAddressInfo .
func CliGetAddressInfo(addr string) (*btcjson.GetAddressInfoResp, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getaddressinfo", addr,
	))
	var resp btcjson.GetAddressInfoResp
	err := json.Unmarshal([]byte(cmdPrint), &resp)
	return &resp, err
}

// CliGetbalance .
func CliGetbalance(_dummy *string, minconf *int, includeWatchonly *bool) (float64, error) {
	args := []string{CmdParamRegtest, "getbalance"}
	if _dummy == nil {
		args = append(args, "*")
	} else {
		args = append(args, *_dummy)
	}

	if minconf != nil {
		args = append(args, strconv.Itoa(*minconf))
	} else {
		args = append(args, "0")
	}

	if includeWatchonly != nil {
		if *includeWatchonly {
			args = append(args, "true")
		} else {
			args = append(args, "false")
		}
	} else {
		args = append(args, "false")
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	return strconv.ParseFloat(cmdPrint, 64)
}

// CliGetWalletInfo .
func CliGetWalletInfo() map[string]interface{} {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "getwalletinfo",
	))
	var info map[string]interface{}
	json.Unmarshal([]byte(cmdPrint), &info)
	return info
}

// CliGetblockcount .
func CliGetblockcount() (int, error) {
	cmd := exec.Command(CmdBitcoinCli, CmdParamRegtest, "getblockcount")
	cmdPrint := cmdAndPrint(cmd)
	cmdPrint = strings.TrimSpace(cmdPrint)
	return strconv.Atoi(cmdPrint)
}

// CliGetblockhash .
func CliGetblockhash(height int) (string, error) {
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, CmdParamRegtest, "getblockhash", strconv.Itoa(height)))
	//TODO validate hash
	return strings.TrimSpace(cmdPrint), nil
}

// CliGetblock https://bitcoin.org/en/developer-reference#getblock
func CliGetblock(hash string, verbosity int) (*string, *btcjson.GetBlockResultV1, *btcjson.GetBlockResultV2, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest,
		"getblock",
		hash,
		strconv.Itoa(verbosity),
	))
	var (
		hex string
		b   btcjson.GetBlockResultV1
		b2  btcjson.GetBlockResultV2
		err error
	)
	switch verbosity {
	case 0:
		hex = cmdPrint
	case 1:
		err = json.Unmarshal([]byte(cmdPrint), &b)
	case 2:
		err = json.Unmarshal([]byte(cmdPrint), &b2)
	default:
		err = fmt.Errorf("verbosity must one of 0/1/2, got: %d", verbosity)
	}
	return &hex, &b, &b2, err
}

// CliGetnewaddress https://bitcoin.org/en/developer-reference#getnewaddress
func CliGetnewaddress(labelPtr, addressTypePtr *string) (hexedAddress string, err error) {
	label := ""
	if labelPtr != nil {
		label = *labelPtr
	}
	args := []string{CmdParamRegtest, "getnewaddress", label}
	if addressTypePtr != nil {
		args = append(args, *addressTypePtr)
	}
	cmdPrint := cmdAndPrint(exec.Command(CmdBitcoinCli, args...))
	//TODO validate address
	return cmdPrint, nil
}

// CliGettransaction https://bitcoin.org/en/developer-reference#gettransaction
func CliGettransaction(txid string, includeWatchonly bool) (*btcjson.GetTransactionResult, error) {
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, CmdParamRegtest, "gettransaction", txid, strconv.FormatBool(includeWatchonly),
	))
	var tx btcjson.GetTransactionResult
	err := json.Unmarshal([]byte(cmdPrint), &tx)
	return &tx, err
}

// CliGetrawtransaction .
func CliGetrawtransaction(cmd btcjson.GetRawTransactionCmd) (*btcjson.RawTx, error) {
	args := []string{ //TODO verbose and blockhash process
		CmdParamRegtest,
		"getrawtransaction",
		cmd.Txid,
		strconv.FormatBool(true),
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	var tx btcjson.RawTx
	err := json.Unmarshal([]byte(cmdPrint), &tx)
	return &tx, err
}

// CliGetreceivedbyaddress https://bitcoin.org/en/developer-reference#getreceivedbyaddress
func CliGetreceivedbyaddress(addr string, minconf int) (string, error) {
	args := []string{
		CmdParamRegtest,
		"getreceivedbyaddress",
		addr,
		strconv.Itoa(minconf),
	}
	cmdPrint := cmdAndPrint(exec.Command(
		CmdBitcoinCli, args...,
	))
	return cmdPrint, nil
}
