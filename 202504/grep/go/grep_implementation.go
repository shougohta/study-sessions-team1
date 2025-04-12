package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GrepImplementation はGrepの基本実装を提供する
type GrepImplementation struct{}

// Search はファイルから特定のパターンを検索する
func (g *GrepImplementation) Search(filePath, pattern string) []string {
	matchingLines := []string{}
	
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("ファイル読み込みエラー: %v\n", err)
		return matchingLines
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		found := false
		for i := 0; i <= len(line)-len(pattern); i++ {
			if line[i:i+len(pattern)] == pattern {
				found = true
				break
			}
		}
		if found {
			matchingLines = append(matchingLines, line)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("ファイル読み込みエラー: %v\n", err)
	}
	
	return matchingLines
}
