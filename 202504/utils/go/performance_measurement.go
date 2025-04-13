package utils

import (
	"fmt"
	"runtime"
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
