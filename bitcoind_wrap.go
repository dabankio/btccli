package btccli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// StartOptions .
type StartOptions struct {
	NewTmpDir bool //创建一个临时的目录，并在完成后立即删除这个目录
}

// BitcoindWithOptions .
func BitcoindWithOptions(options StartOptions, args ...string) (func(), error) {
	if options.NewTmpDir {
		for _, arg := range args {
			if strings.Contains(arg, "-datadir=") {
				return nil, fmt.Errorf("参数里似乎已经指定了-datadir  >> %v", arg)
			}
		}

		tmpDir := strings.TrimRight(os.TempDir(), "/")

		dataDir := tmpDir + "/btccli_bitcoind_data_tmp_" + time.Now().Format(time.RFC3339) + "/"
		err := os.MkdirAll(dataDir, 0777)
		if err != nil {
			return nil, fmt.Errorf("cannot create tmp dir: %v, err: %v", dataDir, err)
		}
		args = append(args, "-datadir="+dataDir)
		killbitcoind, err := BitcoindRegtest(args...)
		if err != nil {
			return nil, err
		}
		return func() {
			os.RemoveAll(dataDir)
			killbitcoind()
		}, nil
	}
	return BitcoindRegtest(args...)
}

// BitcoindRegtest 启动bitcoind -regtest 用以测试,返回杀死bitcoind的函数
// Usage:
// killbitcoind, err := btccli.BitcoindRegtest()
// defer killbitcoind()
func BitcoindRegtest(args ...string) (func(), error) {
	if cmdIsPortContainsNameRunning(RPCPortRegtest, "bitcoin") {
		return nil, fmt.Errorf("bitcoind 似乎已经运行在18443端口了,不先杀掉的话数据可能有问题")
	}

	closeChan := make(chan struct{})

	//bitcoin/share/rpcauth$ python3 rpcauth.py rpcusr 233
	//String to be appended to bitcoin.conf:
	//rpcauth=rpcusr:656f9dabc62f0eb697c801369617dc60$422d7fca742d4a59460f941dc9247c782558367edcbf1cd790b2b7ff5624fc1b
	//Your password:
	//233
	options := []string{
		"-regtest",
		"-txindex",
		"-rpcauth=rpcusr:656f9dabc62f0eb697c801369617dc60$422d7fca742d4a59460f941dc9247c782558367edcbf1cd790b2b7ff5624fc1b",
		// "-addresstype=bech32",
		"-rpcport=18443",
	}
	options = append(options, args...)

	cmd := exec.Command(CmdBitcoind, options...)
	fmt.Println(cmd.Args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		fmt.Println("Wait for message to kill bitcoind")
		<-closeChan
		fmt.Println("Received message,killing bitcoind regtest")

		if e := cmd.Process.Kill(); e != nil {
			fmt.Println("关闭 bitcoind 时发生异常", e)
		}
		fmt.Println("关闭 bitcoind 完成")
		closeChan <- struct{}{}
	}()

	// err = cmd.Wait()
	fmt.Println("等待1.5秒,让 bitcoind 启动")
	time.Sleep(time.Millisecond * 1500)
	return func() {
		closeChan <- struct{}{}
	}, nil
}
