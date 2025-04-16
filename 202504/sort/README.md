# Sort 実装ガイド

## 📂 ファイル構造

```
sort/
├── go/
│   ├── main.go
│   └── impl/
│       ├── sort_implementation.go     # 実装ファイル
│       └── sort_performance_measurement.go
│
├── javascript/
│   ├── sort_implementation.js         # 実装ファイル
│   └── sort_performance_measurement.js
│
├── ruby/
│   ├── sort_implementation.rb         # 実装ファイル
│   └── sort_performance_measurement.rb
│
└── test_cases/                        # テストケース
    ├── case1/
    ├── case2/
    └── ...
```

## 🔍 ファイルの役割

| ファイル名 | 説明 |
|------------|------|
| **sort_implementation.xx** | ソートアルゴリズムのメイン実装ファイル。ここに実装を記述します |
| **sort_performance_measurement.xx** | 実装したソートの検証とパフォーマンス計測を行うためのファイル |

## ⚙️ 実装方法

実装は各言語の **sort_implementation** ファイルに記述してください。

- ファイル分割は自由に行っていただいて構いません
- ただし、インターフェイスは **sort_implementation** で定義されているものに合わせてください
  - インターフェイスが一致しないと、正確なパフォーマンス測定および正当性検証ができません

## 🧪 テストと計測方法

### Go

```bash
go run <your_path>/sort/go/main.go "<your_path>/sort/test_cases/<テストケース>"
```

**例:**
```bash
go run sort/go/main.go "./sort/test_cases/case1"
```

### JavaScript (Node.js)

```bash
node <your_path>/sort/javascript/sort_performance_measurement.js "<your_path>/sort/test_cases/<テストケース>"
```

**例:**
```bash
node sort/javascript/sort_performance_measurement.js "./sort/test_cases/case1"
```

### Ruby

```bash
ruby <your_path>/sort/ruby/sort_performance_measurement.rb "<your_path>/sort/test_cases/<テストケース>"
```

**例:**
```bash
ruby sort/ruby/sort_performance_measurement.rb "./sort/test_cases/case1"
```

## 📝 実装のポイント

- **効率的なアルゴリズム**: 様々なデータサイズに対して効率的に動作するソートアルゴリズムを選択
- **任意のデータ型対応**: 数値、文字列など異なるデータ型に対応できるようにする
- **メモリ効率**: 大きな配列でもメモリ効率良く処理できるよう実装
- **安定なソート**: 同値要素の順序が保持されなくても良いが、実装によっては安定ソートも可能

## 💡 テストケースの構成

各テストケースディレクトリには以下のファイルが含まれています:

- `input.txt`: ソートする配列データ（JSON形式）
- `expected.txt`: ソート後の期待される配列（JSON形式）

これらのファイルを使用して、実装したソートアルゴリズムの正確性が検証されます。
