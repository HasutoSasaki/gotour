# A Tour of Go - 写経練習

このリポジトリは[A Tour of Go](https://go.dev/tour/)の内容を写経しながら学習するためのものです。

## 進捗状況

- [x] 0. イントロダクション
- [ ] 1.  Basics
- [ ] 2.  Packages, variables, and functions
- [ ] 3.  Flow control statements
- [ ] 4.  More types: structs, slices, and maps
- [ ] 5.  Methods and interfaces
- [ ] 6.  Interfaces
- [ ] 7.  Concurrency

## ディレクトリ構造

```
gotour/
├── 01-basics/              # 基本的な構文
│   ├── hello/             # Hello, World!
│   ├── sandbox/           # 乱数生成
│   └── packages/          # パッケージの使用
├── 02-packages-variables-functions/  # パッケージ、変数、関数
│   ├── functions/         # 関数の定義
│   ├── multiple-results/  # 複数の戻り値
│   └── variables/         # 変数の宣言
├── 03-flow-control/        # フロー制御文
│   ├── for-loop/          # for文
│   └── if-statement/      # if文
├── 04-more-types/          # より多くの型
├── 05-methods/             # メソッド
├── 06-interfaces/          # インターフェース
├── 07-concurrency/         # 並行性
└── README.md
```

## 実行方法

各プログラムは以下のように実行できます：

```bash
cd 01-basics/hello && go run main.go
cd 02-packages-variables-functions/functions && go run main.go
```

## 学習メモ

各章で学んだことや気づいたことをここに記録していきます。
