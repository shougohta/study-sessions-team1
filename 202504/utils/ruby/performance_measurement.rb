#!/usr/bin/env ruby
# performance_measurement.rb
# Grep、Sort、HashMapの各アルゴリズムのパフォーマンスと正当性を計測するプログラム

require 'benchmark'
require 'get_process_mem'
require 'json'

# ============================
# グローバル計測用ユーティリティ
# ============================

# パフォーマンス計測を行うクラス
class PerformanceMeasurement
  # 処理時間とメモリ使用量を計測する
  def self.measure(name)
    memory_before = GetProcessMem.new.mb
    
    time = Benchmark.realtime do
      yield
    end
    
    memory_after = GetProcessMem.new.mb
    memory_used = memory_after - memory_before
    
    puts "#{name} 計測結果:"
    puts "  実行時間: #{(time * 1000).round(2)} ミリ秒"
    puts "  メモリ使用量: #{memory_used.round(2)} MB"
    puts "  合計メモリ: #{memory_after.round(2)} MB"
    puts "-------------------------------"
    
    return {
      time_ms: (time * 1000).round(2),
      memory_used_mb: memory_used.round(2),
      total_memory_mb: memory_after.round(2)
    }
  end
end
