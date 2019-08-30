package btccli

import (
	"encoding/json"
	"fmt"

	"github.com/lomocoin/btccli/btcjson"
)

func panicIf(e error, msg string) {
	if e != nil {
		panic(fmt.Errorf("【ERR】 %s %v", msg, e))
	}
}

func jsonStr(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", " ")
	return string(b)
}

// ToJSONIndent .
func ToJSONIndent(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", " ")
	return string(b)
}

// ToJSON .
func ToJSON(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

// IfOrString if flag return s ,or s2
func IfOrString(flag bool, trueS, falseS string) string {
	if flag {
		return trueS
	}
	return falseS
}

func dividePrint(msg string) {
	fmt.Printf("\n--------------%s--------------\n", msg)
}

func isCoinbaseTx(tx *btcjson.GetTransactionResult) bool {
	flag := false
	for _, dtl := range tx.Details {
		if dtl.Category == "immature" || dtl.Category == "generate" {
			flag = true
			break
		}
	}
	return len(tx.Details) == 0 && flag
}
