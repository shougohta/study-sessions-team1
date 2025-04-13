package main

import (
	"fmt"
	"os"

	impl "study-session/hash_map/go/impl"
)

// ===============================================
// メイン関数
// ===============================================

func main() {
	fmt.Println("==============================")
	fmt.Println("HashMap性能計測と正当性検証")
	fmt.Println("==============================")

	// HashMap計測と検証
	fmt.Println("HashMap実装のテスト")
	fileDir := os.Args[1]
	hashmapResults := impl.MeasureHashMapPerformance(fileDir, 1)

	// 検証結果の要約
	fmt.Println("\n==============================")
	fmt.Println("テスト結果サマリー")
	fmt.Println("==============================")

	hashmapValid := false
	if hashmapResults != nil {
		hashmapValid, _ = hashmapResults["valid"].(bool)
	}

	fmt.Printf("HashMap: %s\n", boolToCheckmark(hashmapValid))
}

// boolToCheckmark はブール値をチェックマーク文字列に変換
func boolToCheckmark(b bool) string {
	if b {
		return "成功 ✓"
	}
	return "失敗 ✗"
}
