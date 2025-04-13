package main

import (
	"fmt"
	"os"

	impl "study-session/sort/go/impl"
)

// ===============================================
// メイン関数
// ===============================================

func main() {
	fmt.Println("==============================")
	fmt.Println("Sort性能計測と正当性検証")
	fmt.Println("==============================")

	// Sort計測と検証
	fmt.Println("Sort実装のテスト")
	fileDir := os.Args[1]
	sortResults := impl.MeasureSortPerformance(fileDir, 1)

	// 検証結果の要約
	fmt.Println("\n==============================")
	fmt.Println("テスト結果サマリー")
	fmt.Println("==============================")

	sortValid := false
	if sortResults != nil {
		sortValid, _ = sortResults["valid"].(bool)
	}

	fmt.Printf("Sort: %s\n", boolToCheckmark(sortValid))
}

// boolToCheckmark はブール値をチェックマーク文字列に変換
func boolToCheckmark(b bool) string {
	if b {
		return "成功 ✓"
	}
	return "失敗 ✗"
}
