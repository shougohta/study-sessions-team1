require_relative 'hash_map_implementation'
require_relative '../../utils/ruby/performance_measurement'
require_relative '../../utils/ruby/validity_check'

# HashMap実装のパフォーマンスと正当性を計測するクラス
class HashMapPerformanceMeasurement
  def initialize(file_path, implementation = HashMapImplementation.new)
    @file_path = file_path
    @implementation = implementation
  end
  
  # 入力ファイルと期待値ファイルを読み込む
  def load_test_data
    input_path = [@file_path, 'input.txt'].join(File::SEPARATOR)
    expected_path = [@file_path, 'expected.txt'].join(File::SEPARATOR)

    # 入力データの読み込み
    operations = []
    if File.exist?(input_path)
      input_text = File.read(input_path).strip
      begin
        operations = eval(input_text)
        raise "操作の配列が期待されていますが、#{operations.class} が得られました" unless operations.is_a?(Array)
      rescue => e
        puts "入力データのパースに失敗しました: #{e.message}"
        operations = []
      end
    end
    
    # 期待値の読み込み
    expected_output = {}
    if File.exist?(expected_path)
      expected_text = File.read(expected_path).strip
      begin
        expected_output = eval(expected_text)
        raise "ハッシュが期待されていますが、#{expected_output.class} が得られました" unless expected_output.is_a?(Hash)
      rescue => e
        puts "期待値データのパースに失敗しました: #{e.message}"
        expected_output = {}
      end
    end
    expected_output = expected_output.transform_keys(&:to_s)

    return operations, expected_output
  end
  
  # パフォーマンス計測と正当性検証を実行
  def run_measurement(iterations = 1)
    begin
      operations, expected_output = load_test_data
    rescue => e
      puts "テストデータの読み込みに失敗しました: #{e.message}"
      return nil
    end
    
    puts "HashMap実装のパフォーマンス計測と正当性検証:"
    puts "操作数: #{operations.length}"
    puts "繰り返し回数: #{iterations}"
    
    # 処理時間とメモリ使用量を計測
    results = PerformanceMeasurement.measure("HashMap") do
      iterations.times do
        # 新しいインスタンスで開始
        @implementation = HashMapImplementation.new if iterations > 1
        
        operations.each do |operation|
          case operation[:action].to_sym
          when :put
            @implementation.put(operation[:key], operation[:value])
          when :get
            value = @implementation.get(operation[:key])
            puts "取得: #{operation[:key]} => #{value}" if iterations == 1 && operation[:debug]
          when :remove
            @implementation.remove(operation[:key])
          end
        end
      end
    end
    
    # 正当性検証：全エントリを取得して期待値と比較
    actual_entries = @implementation.all_entries
    valid = ValidityCheck.verify("HashMap", actual_entries, expected_output)
    results[:valid] = valid
    
    results
  end
end

if __FILE__ == $0
  puts "=============================="
  puts "HashMapパフォーマンス計測"
  puts "=============================="
  
  # HashMap計測と検証
  puts "HashMap実装のテスト"
  file_path = ARGV[0] || ''
  hashmap_measurement = HashMapPerformanceMeasurement.new(file_path)
  hashmap_results = hashmap_measurement.run_measurement
  
  # 検証結果の要約
  puts "\n=============================="
  puts "テスト結果サマリー"
  puts "=============================="
  puts "HashMap: #{hashmap_results && hashmap_results[:valid] ? '成功 ✓' : '失敗 ✗'}"
end
