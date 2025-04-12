require_relative 'sort_implementation'
require_relative '../../utils/ruby/performance_measurement'
require_relative '../../utils/ruby/validity_check'

# Sort実装のパフォーマンスと正当性を計測するクラス
class SortPerformanceMeasurement
  def initialize(file_path, implementation = SortImplementation.new)
    @file_path = file_path
    @implementation = implementation
  end
  
  # 入力ファイルと期待値ファイルを読み込む
  def load_test_data
    input_path = [@file_path, 'input.txt'].join(File::SEPARATOR)
    expected_path = [@file_path, 'expected.txt'].join(File::SEPARATOR)

    # 入力データの読み込み
    array = []
    if File.exist?(input_path)
      input_text = File.read(input_path).strip
      # 入力データをevalして配列を取得
      begin
        array = eval(input_text)
        raise "配列が期待されていますが、#{array.class} が得られました" unless array.is_a?(Array)
      rescue => e
        puts "入力データのパースに失敗しました: #{e.message}"
        array = []
      end
    end
    
    # 期待値の読み込み
    expected_output = []
    if File.exist?(expected_path)
      expected_text = File.read(expected_path).strip
      begin
        expected_output = eval(expected_text)
        raise "配列が期待されていますが、#{expected_output.class} が得られました" unless expected_output.is_a?(Array)
      rescue => e
        puts "期待値データのパースに失敗しました: #{e.message}"
        expected_output = []
      end
    end
    
    return array, expected_output
  end
  
  # パフォーマンス計測と正当性検証を実行
  def run_measurement(array = nil, expected_output = nil, iterations = 1)
    # テストデータが指定されていない場合は読み込む
    if array.nil? || expected_output.nil?
      begin
        array, expected_output = load_test_data
      rescue => e
        puts "テストデータの読み込みに失敗しました: #{e.message}"
        return nil
      end
    end
    
    puts "Sort実装のパフォーマンス計測と正当性検証:"
    puts "配列サイズ: #{array.length}"
    puts "データ型: #{array.first.class}" if array.any?
    puts "繰り返し回数: #{iterations}"
    
    sorted_array = nil
    
    # 処理時間とメモリ使用量を計測
    results = PerformanceMeasurement.measure("Sort") do
      iterations.times do |i|
        sorted_array = @implementation.sort(array.dup)
        # 最後の反復の結果を保持
        if i == iterations - 1
          puts "ソート前の先頭5要素: #{array.first(5)}" 
          puts "ソート後の先頭5要素: #{sorted_array.first(5)}"
        end
      end
    end
    
    # 正当性検証
    valid = ValidityCheck.verify("Sort", sorted_array, expected_output)
    results[:valid] = valid
    
    results
  end
end

if __FILE__ == $0
  puts "=============================="
  puts "Sortパフォーマンス計測"
  puts "=============================="
  
  # Sort計測と検証
  puts "Sort実装のテスト"
  file_path = ARGV[0] || ''
  sort_measurement = SortPerformanceMeasurement.new(file_path)
  sort_results = sort_measurement.run_measurement
  
  # 検証結果の要約
  puts "\n=============================="
  puts "テスト結果サマリー"
  puts "=============================="
  puts "Sort: #{sort_results && sort_results[:valid] ? '成功 ✓' : '失敗 ✗'}"
end
