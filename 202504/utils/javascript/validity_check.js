/**
 * 結果が期待値と一致するか検証する
 * @param {string} name - テスト名
 * @param {any} result - 実際の結果
 * @param {any} expected - 期待値
 * @returns {boolean} 検証結果
 */
export function verifyResult(name, result, expected) {
  // 深い比較
  const isEqual = JSON.stringify(result) === JSON.stringify(expected);
  
  if (isEqual) {
    console.log(`${name} 正当性検証: 成功 ✓`);
    return true;
  } else {
    console.log(`${name} 正当性検証: 失敗 ✗`);
    console.log(`  期待値: `, expected);
    console.log(`  実際の結果: `, result);
    return false;
  }
}
