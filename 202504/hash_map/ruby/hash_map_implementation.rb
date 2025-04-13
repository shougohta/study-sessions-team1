# HashMapの基本実装
class HashMapImplementation
  def initialize
    # バケット列, 配列サイズなど、必要な要素を初期化する
    @hash = {}
  end
  
  # キーのハッシュ値を計算
  def hash_key(key)
    # keyに対応するハッシュ値を計算するハッシュ関数
  end
  
  # キーと値のペアを格納
  def put(key, value)
    @hash[key] = value
  end
  
  # キーに対応する値を取得
  def get(key)
    @hash[key]
  end
  
  # キーに対応するエントリを削除
  def remove(key)
    @hash.delete(key)
  end
  
  # バケットサイズを拡張
  def resize
    # バケットサイズを拡張する処理
  end
  
  # 現在の要素数を取得
  def size
    @hash.size
  end
  
  # 全ての要素を取得（テスト用）
  def all_entries
    @hash.dup
  end
end
