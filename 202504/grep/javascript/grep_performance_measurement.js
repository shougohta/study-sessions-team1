import fs from 'fs';
import { GrepImplementation } from './grep_implementation.js';
import { measurePerformance } from '../../utils/javascript/performance_measurement.js';
import { verifyResult } from '../../utils/javascript/validity_check.js';
/**
 * 入力ファイルと期待値ファイルからGrepテストデータを読み込む
 * @returns {Promise<{filePath: string, pattern: string, expectedOutput: string[]}>} テストデータ
 */
async function loadGrepTestData(fileDir) {
  // 入力データの読み込み
  let filePath = '';
  let pattern = '';
  
  try {
    const inputData = await fs.promises.readFile(fileDir+'/input.txt', 'utf8');
    const lines = inputData.split('\n');
    
    if (lines.length >= 2) {
      filePath = fileDir + '/' + lines[0].trim();
      pattern = lines[1].trim();
    } else {
      throw new Error('入力データのフォーマットが不正です');
    }
  } catch (error) {
    throw new Error(`入力ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  // 期待値の読み込み
  let expectedOutput = [];
  
  try {
    const expectedData = await fs.promises.readFile(fileDir+'/expected.txt', 'utf8');
    expectedOutput = expectedData.split('\n').filter(line => line.trim() !== '');
  } catch (error) {
    throw new Error(`期待値ファイルの読み込みに失敗しました: ${error.message}`);
  }
  
  return { filePath, pattern, expectedOutput };
}

/**
 * Grep実装のパフォーマンスと正当性を計測するクラス
 */
class GrepPerformanceMeasurement {
  constructor(filePath, implementation = new GrepImplementation()) {
    this.filePath = filePath;
    this.implementation = implementation;
  }
  
  /**
   * パフォーマンス計測と正当性検証を実行
   * @param {string} filePath - 検索対象のファイルパス
   * @param {string} pattern - 検索するパターン
   * @param {string[]} expectedOutput - 期待される出力
   * @param {number} iterations - 繰り返し回数
   * @returns {Promise<Object>} 計測結果
   */
  async runMeasurement(iterations = 1) {
    let filePath = '';
    let pattern = '';
    let expectedOutput = [];
    try {
      const testData = await loadGrepTestData(this.filePath);
      filePath = testData.filePath;
      pattern = testData.pattern;
      expectedOutput = testData.expectedOutput;
    } catch (error) {
      console.error(error.message);
      return null;
    }
    
    console.log("Grep実装のパフォーマンス計測と正当性検証:");
    console.log(`ファイル: ${filePath}`);
    console.log(`検索パターン: ${pattern}`);
    console.log(`繰り返し回数: ${iterations}`);
    
    let matchingLines = [];
    
    // 処理時間とメモリ使用量を計測
    const results = await measurePerformance("Grep", async () => {
      for (let i = 0; i < iterations; i++) {
        matchingLines = await this.implementation.search(filePath, pattern);
        // 必要に応じて結果を確認
        if (iterations === 1) {
          console.log(`ヒット数: ${matchingLines.length}`);
        }
      }
    });
    
    // 正当性検証
    const valid = verifyResult("Grep", matchingLines, expectedOutput);
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
    console.log("Grep実装のパフォーマンス計測と正当性検証");
    console.log("==============================");
    
    // Grep計測と検証
    console.log("Grep実装のテスト");
    const grepMeasurement = new GrepPerformanceMeasurement(process.argv[2]);
    const grepResults = await grepMeasurement.runMeasurement();

    // 検証結果の要約
    console.log("\n==============================");
    console.log("テスト結果サマリー");
    console.log("==============================");
    console.log(`Grep: ${grepResults && grepResults.valid ? '成功 ✓' : '失敗 ✗'}`);
    
  } catch (error) {
    console.error('エラーが発生しました:', error);
  }
}

// プログラム実行
if (process.argv[1] === new URL(import.meta.url).pathname) {
  main().catch(console.error);
}
