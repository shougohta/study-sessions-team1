import fs from 'fs';
import { HashMapImplementation } from './hash_map_implementation.js';
import { measurePerformance } from '../../utils/javascript/performance_measurement.js';
import { verifyResult } from '../../utils/javascript/validity_check.js';
/**
 * 入力ファイルと期待値ファイルからHashMapテストデータを読み込む
 * @returns {Promise<{operations: Array, expectedOutput: Object}>} テストデータ
 */
async function loadHashMapTestData(fileDir) {
  // 入力データの読み込み
  let operations = [];
  
  try {
    const inputData = await fs.promises.readFile(fileDir+'/input.txt', 'utf8');
    try {
      operations = JSON.parse(inputData);
      if (!Array.isArray(operations)) {
        throw new Error('入力データは操作の配列であるべきです');
      }
    } catch (parseError) {
      throw new Error(`入力データのJSONパースに失敗しました: ${parseError.message}`);
    }
  } catch (error) {
    throw new Error(`入力ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  // 期待値の読み込み
  let expectedOutput = {};
  
  try {
    const expectedData = await fs.promises.readFile(fileDir+'/expected.txt', 'utf8');
    try {
      expectedOutput = JSON.parse(expectedData);
      if (typeof expectedOutput !== 'object' || Array.isArray(expectedOutput)) {
        throw new Error('期待値はオブジェクトであるべきです');
      }
    } catch (parseError) {
      throw new Error(`期待値データのJSONパースに失敗しました: ${parseError.message}`);
    }
  } catch (error) {
    throw new Error(`期待値ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  return { operations, expectedOutput };
}

/**
 * HashMap実装のパフォーマンスと正当性を計測するクラス
 */
class HashMapPerformanceMeasurement {
  constructor(filePath, implementation = new HashMapImplementation()) {
    this.filePath = filePath;
    this.implementation = implementation;
  }
  
  /**
   * パフォーマンス計測と正当性検証を実行
   * @param {Object[]} operations - 実行する操作のリスト
   * @param {Object} expectedOutput - 期待される出力
   * @param {number} iterations - 繰り返し回数
   * @returns {Promise<Object>} 計測結果
   */
  async runMeasurement(iterations = 1) {
    let operations = [];
    let expectedOutput = {};
    try {
      const testData = await loadHashMapTestData(this.filePath);
      operations = testData.operations;
      expectedOutput = testData.expectedOutput;
    } catch (error) {
      console.error(error.message);
      return null;
    }
    
    console.log("HashMap実装のパフォーマンス計測と正当性検証:");
    console.log(`操作数: ${operations.length}`);
    console.log(`繰り返し回数: ${iterations}`);
    
    // 処理時間とメモリ使用量を計測
    const results = await measurePerformance("HashMap", async () => {
      for (let i = 0; i < iterations; i++) {
        // 複数回反復する場合は新しいインスタンスで開始
        if (i > 0) {
          this.implementation = new HashMapImplementation();
        }
        
        for (const operation of operations) {
          switch (operation.action) {
            case 'put':
              this.implementation.put(operation.key, operation.value);
              break;
            case 'get':
              const value = this.implementation.get(operation.key);
              if (iterations === 1 && operation.debug) {
                console.log(`取得: ${operation.key} => ${value}`);
              }
              break;
            case 'remove':
              this.implementation.remove(operation.key);
              break;
          }
        }
      }
    });
    
    // 正当性検証
    const actualEntries = this.implementation.getAllEntries();
    const valid = verifyResult("HashMap", actualEntries, expectedOutput);
    results.valid = valid;
    
    return results;
  }
}

/**
 * メイン実行関数
 */
async function main() {
  try {
    console.log("==============================");
    console.log("HashMap実装のパフォーマンス計測と正当性検証");
    console.log("==============================");
    
    // HashMap計測と検証
    console.log("HashMap実装のテスト");
    const hashmapMeasurement = new HashMapPerformanceMeasurement(process.argv[2]);
    const hashmapResults = await hashmapMeasurement.runMeasurement();
    
    // 検証結果の要約
    console.log("\n==============================");
    console.log("テスト結果サマリー");
    console.log("==============================");
    console.log(`HashMap: ${hashmapResults && hashmapResults.valid ? '成功 ✓' : '失敗 ✗'}`);
    
  } catch (error) {
    console.error('エラーが発生しました:', error);
  }
}

// プログラム実行
if (process.argv[1] === new URL(import.meta.url).pathname) {
  main().catch(console.error);
}
