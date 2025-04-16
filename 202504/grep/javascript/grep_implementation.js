import fs from 'fs';
import readline from 'readline';

/**
 * Grepの基本実装クラス
 */
export class GrepImplementation {
  constructor() {
    // 初期化処理があれば行う
  }
  
  /**
   * ファイルから特定のパターンを検索する
   * @param {string} filePath - 検索対象のファイルパス
   * @param {string} pattern - 検索するパターン
   * @returns {Promise<string[]>} マッチした行のリスト
   */
  async search(filePath, pattern) {
    const matchingLines = [];

    // ファイルを読み込むストリームを作成
    const fileStream = fs.createReadStream(filePath);
    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity
    });

    // 各行に対して処理
    for await (const line of rl) {
      if (line.includes(pattern)) {
        matchingLines.push(line);
      }
    }
    
    return matchingLines;    
  }
}
