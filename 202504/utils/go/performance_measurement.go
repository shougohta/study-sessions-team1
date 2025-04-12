package performance_measurement

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ===============================================
// ユーティリティ関数：メモリ使用量計測
// ===============================================

// getMemoryUsage はメモリ使用量を計測して返す
func getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// MeasurePerformance は関数実行のパフォーマンスを計測する
func MeasurePerformance(name string, fn func()) map[string]interface{} {
	fmt.Printf("%s パフォーマンス計測開始:\n", name)
	
	runtime.GC() // 計測前にGCを実行
	memBefore := getMemoryUsage()
	startTime := time.Now()
	
	// 計測対象の関数を実行
	fn()
	
	duration := time.Since(startTime)
	memAfter := getMemoryUsage()
	memUsed := memAfter - memBefore
	
	fmt.Printf("  実行時間: %.2f ミリ秒\n", float64(duration.Microseconds())/1000.0)
	fmt.Printf("  メモリ使用量: %.2f MB\n", float64(memUsed)/(1024*1024))
	fmt.Printf("  合計メモリ: %.2f MB\n", float64(memAfter)/(1024*1024))
	fmt.Println("-------------------------------")
	
	return map[string]interface{}{
		"time_ms":         float64(duration.Microseconds()) / 1000.0,
		"memory_used_mb":  float64(memUsed) / (1024 * 1024),
		"total_memory_mb": float64(memAfter) / (1024 * 1024),
	}
}

// ===============================================
// Grep実装と計測
// ===============================================

// ===============================================
// Sort実装と計測
// ===============================================

// SortImplementation はソートアルゴリズムの基本実装を提供する
type SortImplementation struct{}

// Sort は配列をソートする（クイックソート実装）
func (s *SortImplementation) Sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	
	// クイックソートの実装
	pivot := array[len(array)/2]
	var left []int
	var middle []int
	var right []int
	
	for _, item := range array {
		if item < pivot {
			left = append(left, item)
		} else if item == pivot {
			middle = append(middle, item)
		} else {
			right = append(right, item)
		}
	}
	
	left = s.Sort(left)
	right = s.Sort(right)
	
	result := append(left, middle...)
	result = append(result, right...)
	return result
}

// loadSortTestData は入力ファイルと期待値ファイルを読み込む
func loadSortTestData() ([]int, []int, error) {
	// 入力データの読み込み
	inputData, err := ioutil.ReadFile("input.go")
	if err != nil {
		return nil, nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 入力データのパースと配列への変換
	var array []int
	inputStr := strings.TrimSpace(string(inputData))
	inputStr = strings.Trim(inputStr, "[]")
	if inputStr != "" {
		for _, numStr := range strings.Split(inputStr, ",") {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, fmt.Errorf("数値のパースに失敗しました: %v", err)
			}
			array = append(array, num)
		}
	}
	
	// 期待値の読み込み
	expectedData, err := ioutil.ReadFile("expected.go")
	if err != nil {
		return array, nil, fmt.Errorf("期待値ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 期待値のパースと配列への変換
	var expectedOutput []int
	expectedStr := strings.TrimSpace(string(expectedData))
	expectedStr = strings.Trim(expectedStr, "[]")
	if expectedStr != "" {
		for _, numStr := range strings.Split(expectedStr, ",") {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return array, nil, fmt.Errorf("期待値のパースに失敗しました: %v", err)
			}
			expectedOutput = append(expectedOutput, num)
		}
	}
	
	return array, expectedOutput, nil
}

// MeasureSortPerformance はSortの性能と正当性を計測する
func MeasureSortPerformance(array []int, expectedOutput []int, iterations int) map[string]interface{} {
	// テストデータの読み込み（引数が指定されていない場合）
	if len(array) == 0 {
		var err error
		array, expectedOutput, err = loadSortTestData()
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	
	sorter := &SortImplementation{}
	
	fmt.Printf("Sort実装のパフォーマンス計測と正当性検証:\n")
	fmt.Printf("配列サイズ: %d\n", len(array))
	fmt.Printf("繰り返し回数: %d\n", iterations)
	
	var sorted []int
	
	// 処理時間とメモリ使用量を計測
	results := measurePerformance("Sort", func() {
		for i := 0; i < iterations; i++ {
			// 配列のコピーを作成
			arrayCopy := make([]int, len(array))
			copy(arrayCopy, array)
			
			sorted = sorter.Sort(arrayCopy)
			if iterations == 1 {
				// ソート前とソート後の最初の5要素を表示
				fmt.Printf("ソート前の先頭5要素: ")
				for j := 0; j < 5 && j < len(array); j++ {
					fmt.Printf("%d ", array[j])
				}
				fmt.Println()
				
				fmt.Printf("ソート後の先頭5要素: ")
				for j := 0; j < 5 && j < len(sorted); j++ {
					fmt.Printf("%d ", sorted[j])
				}
				fmt.Println()
			}
		}
	})
	
	// 正当性検証
	valid := verifyResult("Sort", sorted, expectedOutput)
	results["valid"] = valid
	
	return results
}

// ===============================================
// HashMap実装と計測
// ===============================================

// Entry はハッシュマップのエントリ（キーと値のペア）を表す
type Entry struct {
	key   string
	value string
}

// HashMapImplementation はHashMapの基本実装を提供する
type HashMapImplementation struct {
	buckets    [][]Entry
	size       int
	loadFactor float64
}

// NewHashMap は新しいHashMapを作成する
func NewHashMap(bucketSize int) *HashMapImplementation {
	return &HashMapImplementation{
		buckets:    make([][]Entry, bucketSize),
		size:       0,
		loadFactor: 0.75,
	}
}

// hashKey はキーのハッシュ値を計算する
func (h *HashMapImplementation) hashKey(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % len(h.buckets)
	}
	return hash
}

// Put はキーと値のペアを格納する
func (h *HashMapImplementation) Put(key, value string) {
	// 負荷係数を超えた場合はリサイズ
	if float64(h.size+1)/float64(len(h.buckets)) >= h.loadFactor {
		h.resize()
	}
	
	index := h.hashKey(key)
	
	// バケットが空なら初期化
	if h.buckets[index] == nil {
		h.buckets[index] = []Entry{}
	}
	
	// 既存キーを更新
	for i, entry := range h.buckets[index] {
		if entry.key == key {
			h.buckets[index][i].value = value
			return
		}
	}
	
	// 新規キーを追加
	h.buckets[index] = append(h.buckets[index], Entry{key, value})
	h.size++
}

// Get はキーに対応する値を取得する
func (h *HashMapImplementation) Get(key string) (string, bool) {
	index := h.hashKey(key)
	
	if h.buckets[index] == nil {
		return "", false
	}
	
	for _, entry := range h.buckets[index] {
		if entry.key == key {
			return entry.value, true
		}
	}
	
	return "", false
}

// Remove はキーに対応するエントリを削除する
func (h *HashMapImplementation) Remove(key string) bool {
	index := h.hashKey(key)
	
	if h.buckets[index] == nil {
		return false
	}
	
	for i, entry := range h.buckets[index] {
		if entry.key == key {
			// 最後の要素以外の場合は最後の要素と入れ替えて削除
			lastIdx := len(h.buckets[index]) - 1
			if i != lastIdx {
				h.buckets[index][i] = h.buckets[index][lastIdx]
			}
			// 最後の要素を削除
			h.buckets[index] = h.buckets[index][:lastIdx]
			h.size--
			return true
		}
	}
	
	return false
}

// resize はバケットサイズを拡張する
func (h *HashMapImplementation) resize() {
	oldBuckets := h.buckets
	h.buckets = make([][]Entry, len(oldBuckets)*2)
	h.size = 0
	
	for _, bucket := range oldBuckets {
		if bucket == nil {
			continue
		}
		
		for _, entry := range bucket {
			h.Put(entry.key, entry.value)
		}
	}
}

// Size は現在の要素数を取得する
func (h *HashMapImplementation) Size() int {
	return h.size
}

// GetAllEntries は全てのエントリを取得する（テスト用）
func (h *HashMapImplementation) GetAllEntries() map[string]string {
	result := make(map[string]string)
	for _, bucket := range h.buckets {
		if bucket == nil {
			continue
		}
		
		for _, entry := range bucket {
			result[entry.key] = entry.value
		}
	}
	return result
}

// Operation はHashMapに対する操作を表す
type Operation struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Value  string `json:"value,omitempty"`
	Debug  bool   `json:"debug,omitempty"`
}

// loadHashMapTestData は入力ファイルと期待値ファイルを読み込む
func loadHashMapTestData() ([]Operation, map[string]string, error) {
	// 入力データの読み込み
	inputData, err := ioutil.ReadFile("input.go")
	if err != nil {
		return nil, nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 入力データのパース
	var operations []Operation
	err = json.Unmarshal(inputData, &operations)
	if err != nil {
		return nil, nil, fmt.Errorf("入力データのJSONパースに失敗しました: %v", err)
	}
	
	// 期待値の読み込み
	expectedData, err := ioutil.ReadFile("expected.go")
	if err != nil {
		return operations, nil, fmt.Errorf("期待値ファイルの読み込みに失敗しました: %v", err)
	}
	
	// 期待値のパース
	var expectedOutput map[string]string
	err = json.Unmarshal(expectedData, &expectedOutput)
	if err != nil {
		return operations, nil, fmt.Errorf("期待値データのJSONパースに失敗しました: %v", err)
	}
	
	return operations, expectedOutput, nil
}

// MeasureHashMapPerformance はHashMapの性能と正当性を計測する
func MeasureHashMapPerformance(operations []Operation, expectedOutput map[string]string, iterations int) map[string]interface{} {
	// テストデータの読み込み（引数が指定されていない場合）
	if len(operations) == 0 {
		var err error
		operations, expectedOutput, err = loadHashMapTestData()
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	
	hashMap := NewHashMap(16)
	
	fmt.Printf("HashMap実装のパフォーマンス計測と正当性検証:\n")
	fmt.Printf("操作数: %d\n", len(operations))
	fmt.Printf("繰り返し回数: %d\n", iterations)
	
	// 処理時間とメモリ使用量を計測
	results := measurePerformance("HashMap", func() {
		for i := 0; i < iterations; i++ {
			// 複数回反復する場合は新しいインスタンスで開始
			if i > 0 {
				hashMap = NewHashMap(16)
			}
			
			for _, op := range operations {
				switch op.Action {
				case "put":
					hashMap.Put(op.Key, op.Value)
				case "get":
					value, exists := hashMap.Get(op.Key)
					if iterations == 1 && op.Debug {
						fmt.Printf("取得: %s => %s (存在: %v)\n", op.Key, value, exists)
					}
				case "remove":
					hashMap.Remove(op.Key)
				}
			}
		}
	})
	
	// 正当性検証
	actualEntries := hashMap.GetAllEntries()
	valid := verifyResult("HashMap", actualEntries, expectedOutput)
	results["valid"] = valid
	
	return results
}

// ===============================================
// メイン関数
// ===============================================

func main() {
	fmt.Println("==============================")
	fmt.Println("アルゴリズム実装の性能と正当性検証")
	fmt.Println("==============================")
	
	// Grep計測と検証
	fmt.Println("\n1. Grep実装のテスト")
	grepResults := MeasureGrepPerformance("", "", nil, 1)
	
	// Sort計測と検証
	fmt.Println("\n2. Sort実装のテスト")
	sortResults := MeasureSortPerformance(nil, nil, 1)
	
	// HashMap計測と検証
	fmt.Println("\n3. HashMap実装のテスト")
	hashmapResults := MeasureHashMapPerformance(nil, nil, 1)
	
	// 検証結果の要約
	fmt.Println("\n==============================")
	fmt.Println("テスト結果サマリー")
	fmt.Println("==============================")
	
	grepValid := false
	if grepResults != nil {
		grepValid, _ = grepResults["valid"].(bool)
	}
	
	sortValid := false
	if sortResults != nil {
		sortValid, _ = sortResults["valid"].(bool)
	}
	
	hashmapValid := false
	if hashmapResults != nil {
		hashmapValid, _ = hashmapResults["valid"].(bool)
	}
	
	fmt.Printf("Grep: %s\n", boolToCheckmark(grepValid))
	fmt.Printf("Sort: %s\n", boolToCheckmark(sortValid))
	fmt.Printf("HashMap: %s\n", boolToCheckmark(hashmapValid))
}

// boolToCheckmark はブール値をチェックマーク文字列に変換
func boolToCheckmark(b bool) string {
	if b {
		return "成功 ✓"
	}
	return "失敗 ✗"
}
