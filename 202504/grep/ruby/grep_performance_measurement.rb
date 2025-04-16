require_relative 'grep_implementation'
require_relative '../../utils/ruby/performance_measurement'
require_relative '../../utils/ruby/validity_check'

# Grep実装のパフォーマンスと正当性を計測するクラス
class GrepPerformanceMeasurement
  def initialize(file_path, implementation = GrepImplementation.new)
    @file_path = file_path
    @implementation = implementation
  end
  
  # 入力ファイルと期待値ファイルを読み込む
  def load_test_data
    input_path = [@file_path, 'input.txt'].join(File::SEPARATOR)
    expected_path = [@file_path, 'expected.txt'].join(File::SEPARATOR)

    # 入力データの読み込み
    input_data = File.read(input_path).strip
    # 入力データをパースして、ファイルパスとパターンを取得
    file_name, pattern = input_data.split("\n", 2)
    pattern.strip! if pattern

    file_path = [@file_path, file_name].join(File::SEPARATOR)
    
    # 期待値の読み込み
    expected_output = []
    if File.exist?(expected_path)
      expected_output = File.readlines(expected_path).map(&:chomp)
    end
    
    return file_path, pattern, expected_output
  end
  
  # パフォーマンス計測と正当性検証を実行
  def run_measurement(iterations = 1)
    # テストデータが指定されていない場合は読み込む
    begin
      file_path, pattern, expected_output = load_test_data
    rescue => e
      puts "テストデータの読み込みに失敗しました: #{e.message}"
      return nil
    end
    
    puts "Grep実装のパフォーマンス計測と正当性検証:"
    puts "ファイル: #{file_path}"
    puts "検索パターン: #{pattern}"
    puts "繰り返し回数: #{iterations}"
    
    matching_lines = nil
    
    # 処理時間とメモリ使用量を計測
    results = PerformanceMeasurement.measure("Grep") do
      iterations.times do |i|
        matching_lines = @implementation.search(file_path, pattern)
        # 最後の反復結果を保持
        if i == iterations - 1
          puts "ヒット数: #{matching_lines.length}"
        end
      end
    end
    
    # 正当性検証
    valid = ValidityCheck.verify("Grep", matching_lines, expected_output)
    results[:valid] = valid
    
    results
  end
end

# GrepPerformanceMeasurementクラスを使用して、Grepのパフォーマンスと正当性を計測する
if __FILE__ == $0
  puts "=============================="
  puts "Grepパフォーマンス計測"
  puts "=============================="
  
  # Grep計測と検証
  puts "Grep実装のテスト"
  file_path = ARGV[0] || ''
  grep_measurement = GrepPerformanceMeasurement.new(file_path)
  grep_results = grep_measurement.run_measurement
  
  # 検証結果の要約
  puts "\n=============================="
  puts "テスト結果サマリー"
  puts "=============================="
  puts "Grep: #{grep_results && grep_results[:valid] ? '成功 ✓' : '失敗 ✗'}"
end
