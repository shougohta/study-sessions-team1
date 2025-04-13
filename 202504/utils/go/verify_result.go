package utils

import (
	"fmt"
	"reflect"
)

// VerifyResult は結果の正当性を検証する
func VerifyResult(name string, result interface{}, expected interface{}) bool {
	if reflect.DeepEqual(result, expected) {
		fmt.Printf("%s 正当性検証: 成功 ✓\n", name)
		return true
	} else {
		fmt.Printf("%s 正当性検証: 失敗 ✗\n", name)
		fmt.Printf("  期待値: %v\n", expected)
		fmt.Printf("  実際の結果: %v\n", result)
		return false
	}
}
