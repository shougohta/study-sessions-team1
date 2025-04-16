import fs from 'fs';
import { SortImplementation } from './sort_implementation.js';
import { measurePerformance } from '../../utils/javascript/performance_measurement.js';
import { verifyResult } from '../../utils/javascript/validity_check.js';

/**
 * 入力ファイルと期待値ファイルからSortテストデータを読み込む
 * @returns {Promise<{array: Array, expectedOutput: Array}>} テストデータ
 */
async function loadSortTestData(fileDir) {
  // 入力データの読み込み
  let array = [];
  
  try {
    const inputData = await fs.promises.readFile(fileDir+'/input.txt', 'utf8');
    try {
      array = JSON.parse(inputData);
      if (!Array.isArray(array)) {
        throw new Error('入力データは配列であるべきです');
      }
    } catch (parseError) {
      throw new Error(`入力データのJSONパースに失敗しました: ${parseError.message}`);
    }
  } catch (error) {
    throw new Error(`入力ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  // 期待値の読み込み
  let expectedOutput = [];
  
  try {
    const expectedData = await fs.promises.readFile(fileDir+'/expected.txt', 'utf8');
    try {
      expectedOutput = JSON.parse(expectedData);
      if (!Array.isArray(expectedOutput)) {
        throw new Error('期待値は配列であるべきです');
      }
    } catch (parseError) {
      throw new Error(`期待値データのJSONパースに失敗しました: ${parseError.message}`);
    }
  } catch (error) {
    throw new Error(`期待値ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  return { array, expectedOutput };
}

/**
 * Sort実装のパフォーマンスと正当性を計測するクラス
 */
class SortPerformanceMeasurement {
  constructor(filePath, implementation = new SortImplementation()) {
    this.filePath = filePath;
    this.implementation = implementation;
  }
  
  /**
   * パフォーマンス計測と正当性検証を実行
   * @param {Array} array - ソート対象の配列
   * @param {Array} expectedOutput - 期待される出力
   * @param {number} iterations - 繰り返し回数
   * @returns {Promise<Object>} 計測結果
   */
  async runMeasurement(iterations = 1) {
    let array = [];
    let expectedOutput = [];
    try {
      const testData = await loadSortTestData(this.filePath);
      array = testData.array;
      expectedOutput = testData.expectedOutput;
    } catch (error) {
      console.error(error.message);
      return null;
    }
    
    console.log("Sort実装のパフォーマンス計測と正当性検証:");
    console.log(`配列サイズ: ${array.length}`);
    console.log(`データ型: ${typeof array[0]}`);
    console.log(`繰り返し回数: ${iterations}`);
    
    let sortedArray = [];
    
    // 処理時間とメモリ使用量を計測
    const results = await measurePerformance("Sort", async () => {
      for (let i = 0; i < iterations; i++) {
        // 配列のコピーを作成してソート
        const arrayCopy = [...array];
        sortedArray = this.implementation.sort(arrayCopy);
        
        // 必要に応じて結果を確認
        if (iterations === 1) {
          console.log(`ソート前の先頭5要素: ${array.slice(0, 5)}`);
          console.log(`ソート後の先頭5要素: ${sortedArray.slice(0, 5)}`);
        }
      }
    });
    
    // 正当性検証
    const valid = verifyResult("Sort", sortedArray, expectedOutput);
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
    console.log("Sort実装のパフォーマンス計測と正当性検証");
    console.log("==============================");
    
    // Sort計測と検証
    console.log("Sort実装のテスト");
    const sortMeasurement = new SortPerformanceMeasurement(process.argv[2]);
    const sortResults = await sortMeasurement.runMeasurement();
    
    // 検証結果の要約
    console.log("\n==============================");
    console.log("テスト結果サマリー");
    console.log("==============================");
    console.log(`Sort: ${sortResults && sortResults.valid ? '成功 ✓' : '失敗 ✗'}`);
    
  } catch (error) {
    console.error('エラーが発生しました:', error);
  }
}

// プログラム実行
if (process.argv[1] === new URL(import.meta.url).pathname) {
  main().catch(console.error);
}
