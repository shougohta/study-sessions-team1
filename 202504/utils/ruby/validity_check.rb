#!/usr/bin/env ruby
# performance_measurement.rb
# Grep、Sort、HashMapの各アルゴリズムのパフォーマンスと正当性を計測するプログラム

# 正当性検証ユーティリティ
class ValidityCheck
  # 結果が期待値と一致するかチェック
  def self.verify(name, result, expected)
    if result == expected
      puts "#{name} 正当性検証: 成功 ✓"
      return true
    else
      puts "#{name} 正当性検証: 失敗 ✗"
      puts "  期待値: #{expected.inspect}"
      puts "  実際の結果: #{result.inspect}"
      return false
    end
  end
end