package impl

import (
	"reflect"
	"sort"
)

// SortImplementation はソートアルゴリズムの基本実装を提供する
type SortImplementation struct{}

// SortSliceOrdered は「比較演算子 < が使える型」だけを対象にソートする
func (s *SortImplementation) Sort(data []interface{}) []interface{} {
	// 元スライスをコピー
	newArr := make([]interface{}, len(data))
	copy(newArr, data)
	el := newArr[0]

	if reflect.TypeOf(el) == reflect.TypeOf(0) {
		// 数値型の場合
		sort.Slice(newArr, func(i, j int) bool {
			return newArr[i].(int) < newArr[j].(int)
		})
		return newArr
	}
	if reflect.TypeOf(el) == reflect.TypeOf("") {
		// 文字列型の場合
		sort.Slice(newArr, func(i, j int) bool {
			return newArr[i].(string) < newArr[j].(string)
		})
		return newArr
	}
	if reflect.TypeOf(el) == reflect.TypeOf(0.0) {
		// 浮動小数点型の場合
		sort.Slice(newArr, func(i, j int) bool {
			return newArr[i].(float64) < newArr[j].(float64)
		})
		return newArr
	}

	return newArr
}
