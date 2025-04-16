// performance_measurement.js
// Grep、Sort、HashMapの各アルゴリズムのパフォーマンスと正当性を計測するプログラム

// ============================
// パフォーマンス計測ユーティリティ
// ============================

/**
 * 処理時間とメモリ使用量を計測する
 * @param {string} name - 計測対象の名前
 * @param {Function} callback - 計測する関数
 * @returns {Promise<Object>} 計測結果
 */
export async function measurePerformance(name, callback) {
  // GCを促し、初期メモリ使用量を取得
  if (global.gc) {
    global.gc();
  }
  
  const memoryBefore = process.memoryUsage();
  console.log(`${name} パフォーマンス計測開始:`);
  
  const startTime = process.hrtime();
  
  // 実行
  await callback();
  
  const endTime = process.hrtime(startTime);
  const executionTimeMs = (endTime[0] * 1000 + endTime[1] / 1000000).toFixed(2);
  
  // メモリ使用状況を取得
  const memoryAfter = process.memoryUsage();
  
  // RSSとヒープ使用量の差分を計算
  const memoryUsedRss = (memoryAfter.rss - memoryBefore.rss) / (1024 * 1024);
  const memoryUsedHeap = (memoryAfter.heapUsed - memoryBefore.heapUsed) / (1024 * 1024);
  
  console.log(`  実行時間: ${executionTimeMs} ミリ秒`);
  console.log(`  RSS メモリ増加量: ${memoryUsedRss.toFixed(2)} MB`);
  console.log(`  ヒープ メモリ増加量: ${memoryUsedHeap.toFixed(2)} MB`);
  console.log(`  合計 RSS メモリ: ${(memoryAfter.rss / (1024 * 1024)).toFixed(2)} MB`);
  console.log(`  合計ヒープメモリ: ${(memoryAfter.heapUsed / (1024 * 1024)).toFixed(2)} MB`);
  console.log("-------------------------------");
  
  return {
    timeMs: parseFloat(executionTimeMs),
    memoryUsedRssMb: parseFloat(memoryUsedRss.toFixed(2)),
    memoryUsedHeapMb: parseFloat(memoryUsedHeap.toFixed(2)),
    totalRssMb: parseFloat((memoryAfter.rss / (1024 * 1024)).toFixed(2)),
    totalHeapMb: parseFloat((memoryAfter.heapUsed / (1024 * 1024)).toFixed(2))
  };
}
