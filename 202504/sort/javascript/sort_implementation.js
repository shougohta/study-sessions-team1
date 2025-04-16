/**
 * ソートアルゴリズムの基本実装クラス
 */
export class SortImplementation {
  constructor() {
    // 初期化処理があれば行う
  }
  
  /**
   * 配列をソートする (クイックソート実装)
   * @param {Array} array - ソートする配列
   * @returns {Array} ソートされた配列
   */
  sort(array) {
    if (array.length <= 1) {
      return array;
    }
    
    return array.sort((a, b) => a - b);
  }
}
