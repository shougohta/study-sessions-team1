# Grepの基本実装
class GrepImplementation
  def initialize
    # 初期化処理があれば行う
  end
  
  # ファイルから特定のパターンを検索する
  def search(file_path, pattern)
    # ここを実装する
    File.readlines(file_path, chomp: true).grep(Regexp.new(Regexp.escape(pattern)))
  end
end
