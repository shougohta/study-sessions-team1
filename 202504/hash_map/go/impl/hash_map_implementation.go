package impl

import "fmt"

// HashMapImplementation はHashMapの基本実装を提供する
type HashMapImplementation struct {
	hashMap map[interface{}]interface{}
}

// NewHashMap は新しいHashMapを作成する
func NewHashMap(bucketSize int) *HashMapImplementation {
	return &HashMapImplementation{
		hashMap: make(map[interface{}]interface{}),
	}
}

// hashKey はキーのハッシュ値を計算する
// func (h *HashMapImplementation) hashKey(key string) int {
// 	# TODO: ハッシュ関数を実装する
// }

// Put はキーと値のペアを格納する
func (h *HashMapImplementation) Put(key, value interface{}) {
	h.hashMap[key] = value
}

// Get はキーに対応する値を取得する
func (h *HashMapImplementation) Get(key interface{}) (interface{}, bool) {
	value, exists := h.hashMap[key]
	if !exists {
		return nil, false
	}
	return value, true
}

// Remove はキーに対応するエントリを削除する
func (h *HashMapImplementation) Remove(key interface{}) bool {
	if _, exists := h.hashMap[key]; !exists {
		return false
	}
	delete(h.hashMap, key)
	return true
}

// resize はバケットサイズを拡張する
// func (h *HashMapImplementation) resize() {
//  # TODO: バケットサイズを拡張する
// }

// Size は現在の要素数を取得する
func (h *HashMapImplementation) Size() int {
	return len(h.hashMap)
}

// GetAllEntries は全てのエントリを取得する（テスト用）
func (h *HashMapImplementation) GetAllEntries() map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range h.hashMap {
		result[fmt.Sprintf("%v", k)] = v
	}
	return result
}
