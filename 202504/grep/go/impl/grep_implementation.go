package impl

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// GrepImplementation はGrepの基本実装を提供する
type GrepImplementation struct{}

// Search はファイルから特定のパターンを検索する
func (g *GrepImplementation) Search(filePath, pattern string) []string {
	var result []string

	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("failed to open file %s: %v", filePath, err)
		return result
	}
	defer f.Close()

	// 行単位でスキャン
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// パターンが含まれているか判定
		if strings.Contains(line, pattern) {
			result = append(result, line)
		}
	}

	// スキャン時のエラーがあればログに出す
	if err := scanner.Err(); err != nil {
		log.Printf("failed to read file %s: %v", filePath, err)
	}

	return result
}
