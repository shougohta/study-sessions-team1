package impl

import (
	"bufio"
	"log"
	"os"
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

	// Boyer-Mooreの事前処理（bad character ルール）
	bcTable := buildBadCharTable(pattern)

	// 行単位でスキャン
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Boyer-Mooreで検索
		if boyerMooreSearch(line, pattern, bcTable) {
			result = append(result, line)
		}
	}

	// スキャン時のエラーがあればログに出す
	if err := scanner.Err(); err != nil {
		log.Printf("failed to read file %s: %v", filePath, err)
	}

	return result
}

// buildBadCharTable は bad character テーブルを構築する
func buildBadCharTable(pattern string) [256]int {
	var table [256]int
	patLen := len(pattern)
	for i := 0; i < 256; i++ {
		table[i] = patLen
	}
	for i := 0; i < patLen-1; i++ {
		table[pattern[i]] = patLen - 1 - i
	}
	return table
}

// boyerMooreSearch は対象文字列 text に pattern が含まれているかを返す
func boyerMooreSearch(text, pattern string, bcTable [256]int) bool {
	n, m := len(text), len(pattern)
	if m == 0 {
		return true
	}
	i := m - 1
	for i < n {
		j := m - 1
		for j >= 0 && text[i] == pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return true
		}
		skip := bcTable[text[i]]
		if skip < 1 {
			skip = 1
		}
		i += skip
	}
	return false
}
