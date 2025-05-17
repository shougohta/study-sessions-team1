# HashMapの基本実装
class HashMapImplementation
  def initialize
    # バケット列, 配列サイズなど、必要な要素を初期化する
    @size = 0
    @capacity = PRIMES[0]
    @buckets = Array.new(@capacity)
  end
  
  # キーのハッシュ値を計算
  def hash_key(key)
    # keyに対応するハッシュ値を計算するハッシュ関数
    murmur3(key)
  end
  
  # キーと値のペアを格納
  def put(key, value)
    resize if @size >= @capacity * LOAD_FACTOR

    index = hash_key(key) % @capacity
    probe_distance = 0

    loop do
      bucket = @buckets[index]

      if bucket.nil?
        @buckets[index] = [key, value, probe_distance]
        @size += 1
        return
      elsif bucket[0] == key
        @buckets[index][1] = value
        return
      elsif bucket[2] < probe_distance
        key, value, probe_distance, @buckets[index] = bucket[0], bucket[1], bucket[2], [key, value, probe_distance]
      end

      index = (index + 1) % @capacity
      probe_distance += 1
    end
  end
  
  # キーに対応する値を取得
  def get(key)
    index = hash_key(key) % @capacity
    probe_distance = 0

    loop do
      bucket = @buckets[index]
      return nil if bucket.nil?
      return bucket[1] if bucket[0] == key
      break if bucket[2] < probe_distance

      index = (index + 1) % @capacity
      probe_distance += 1
    end

    nil
  end
  
  # キーに対応するエントリを削除
  def remove(key)
    index = hash_key(key) % @capacity
    probe_distance = 0

    loop do
      bucket = @buckets[index]
      return if bucket.nil?

      if bucket[0] == key
        @buckets[index] = nil
        @size -= 1
        rehash_from(index)
        return
      end

      break if bucket[2] < probe_distance

      index = (index + 1) % @capacity
      probe_distance += 1
    end
  end
  
  # バケットサイズを拡張
  def resize
    # バケットサイズを拡張する処理
    old_buckets = @buckets.compact
    next_capacity = next_prime_from_table(@capacity * 2)
    @capacity = next_capacity
    @buckets = Array.new(@capacity)
    @size = 0

    old_buckets.each { |k, v, _| put(k, v) }
  end
  
  # 現在の要素数を取得
  def size
    @size
  end
  
  # 全ての要素を取得（テスト用）
  def all_entries
    @buckets.compact.map { |key, value, _| [key, value] }.to_h
  end

  private

  def murmur3(str, seed = 0)
    data = str.bytes
    length = data.length
    h = seed

    c1 = 0xcc9e2d51
    c2 = 0x1b873593

    i = 0
    while i + 4 <= length
      k = data[i] | (data[i+1] << 8) | (data[i+2] << 16) | (data[i+3] << 24)
      i += 4

      k = (k * c1) & 0xffffffff
      k = (k << 15 | k >> 17) & 0xffffffff
      k = (k * c2) & 0xffffffff

      h ^= k
      h = (h << 13 | h >> 19) & 0xffffffff
      h = (h * 5 + 0xe6546b64) & 0xffffffff
    end

    k = 0
    remain = length & 3
    if remain >= 3
      k ^= data[i+2] << 16
    end
    if remain >= 2
      k ^= data[i+1] << 8
    end
    if remain >= 1
      k ^= data[i]
      k = (k * c1) & 0xffffffff
      k = (k << 15 | k >> 17) & 0xffffffff
      k = (k * c2) & 0xffffffff
      h ^= k
    end

    h ^= length
    h ^= h >> 16
    h = (h * 0x85ebca6b) & 0xffffffff
    h ^= h >> 13
    h = (h * 0xc2b2ae35) & 0xffffffff
    h ^= h >> 16

    h & 0x7fffffff
  end

  def rehash_from(start_index)
    index = (start_index + 1) % @capacity

    while (bucket = @buckets[index])
      @buckets[index] = nil
      @size -= 1
      put(bucket[0], bucket[1])
      index = (index + 1) % @capacity
    end
  end

  def next_prime_from_table(n)
    PRIMES.find { |prime| prime >= n } || PRIMES.last
  end

  PRIMES = [
    53, 97, 193, 389, 769,
    1543, 3079, 6151, 12289, 24593,
    49157, 65521, 90001, 120071, 131071
  ]

  LOAD_FACTOR = 0.75
end
