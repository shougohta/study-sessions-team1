package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fs-19-koudai-yagisawa/study_session/utils/go/performance_measurement"
	"github.com/fs-19-koudai-yagisawa/study_session/utils/go/verify_result"
)
// loadGrepTestData は入力ファイルと期待値ファイルを読み込む
func loadGrepTestData(fileDir string) (string, string, []string, error) {
	// 入力データの読み込み
	inputData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "input.txt"}, "/"))
	if err != nil {
		return "", "", nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 入力データのパース
	inputLines := strings.Split(string(inputData), "\n")
	if len(inputLines) < 2 {
		return "", "", nil, fmt.Errorf("入力データが不足しています")
	}
	
	filePath := strings.TrimSpace(inputLines[0])
	pattern := strings.TrimSpace(inputLines[1])
	
	// 期待値の読み込み
	expectedData, err := ioutil.ReadFile(strings.Join([]string{fileDir, "expected.txt"}, "/"))
	if err != nil {
		return filePath, pattern, nil, fmt.Errorf("期待値ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 期待値のパース
	var expectedOutput []string
	for _, line := range strings.Split(string(expectedData), "\n") {
		if line != "" {
			expectedOutput = append(expectedOutput, line)
		}
	}
	
	return filePath, pattern, expectedOutput, nil
}

// MeasureGrepPerformance はGrepの性能と正当性を計測する
func MeasureGrepPerformance(fileDir string, iterations int) map[string]interface{} {
	// テストデータの読み込み（引数が指定されていない場合）
	var err error
	filePath, pattern, expectedOutput, err = loadGrepTestData(fileDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	grep := &GrepImplementation{}
	
	fmt.Printf("Grep実装のパフォーマンス計測と正当性検証:\n")
	fmt.Printf("ファイル: %s\n", filePath)
	fmt.Printf("検索パターン: %s\n", pattern)
	fmt.Printf("繰り返し回数: %d\n", iterations)
	
	var matchingLines []string
	
	// 処理時間とメモリ使用量を計測
	results := MeasurePerformance("Grep", func() {
		for i := 0; i < iterations; i++ {
			matchingLines = grep.Search(filePath, pattern)
			if iterations == 1 {
				fmt.Printf("ヒット数: %d\n", len(matchingLines))
			}
		}
	})
	
	// 正当性検証
	valid := VerifyResult("Grep", matchingLines, expectedOutput)
	results["valid"] = valid
	
	return results
}


// ===============================================
// メイン関数
// ===============================================

func main() {
	fmt.Println("==============================")
	fmt.Println("Grep性能計測と正当性検証")
	fmt.Println("==============================")
	
	// Grep計測と検証
	fmt.Println("Grep実装のテスト")
	var fileDir string
	grepResults := MeasureGrepPerformance(fileDir, 1)
	
	// 検証結果の要約
	fmt.Println("\n==============================")
	fmt.Println("テスト結果サマリー")
	fmt.Println("==============================")
	
	grepValid := false
	if grepResults != nil {
		grepValid, _ = grepResults["valid"].(bool)
	}
	
	fmt.Printf("Grep: %s\n", boolToCheckmark(grepValid))
}

// boolToCheckmark はブール値をチェックマーク文字列に変換
func boolToCheckmark(b bool) string {
	if b {
		return "成功 ✓"
	}
	return "失敗 ✗"
}
