# 2017/06/09 のメモ

## メニュー

Go 言語でのテストに関する Tips

* テストの書き方
* テストしやすいコードにする
* Table Driven Testing
* カバレッジを計測する
* interface を使ったモック

## テストの書き方おさらい

* 一般的には `xxx_test.go` という名前のファイルを作る。
  * `conduct.go` → `conduct_test.go` というような法則
* 雛形は以下のような形。

```go
package {任意のパッケージ}

import (
    "testing"
)

// テスト時に実行される関数。
func TestXXX(t *testing.T) {
    // ここにテストを書く
}
```

* `testing` パッケージを import する
* Test〜 で始まる関数を定義する。  
テストを走らせると Test〜が順次呼び出される。

## テストしやすいコードにする

* テストコードは任意のパッケージに属する
  * パッケージ内の private な関数も呼び出せる
  * テストしやすい単位に関数を分割する

* [dspector.go](https://github.com/CPSPlatform/CoreService/blob/b3de7f8dad95cc4f676632d2b240b554751c6578/src/dspector/dspector.go)
  * 無限ループする関数一個のみで構成されていて、このままではテストしずらい一面
  * テストしやすい単位に処理を分割する (関数にわける)

### テストしにくい関数

* 入力したっきり出力がない (void を返すタチの) 
* 無限ループして帰ってこない (任意のタイミングで終わらせれない) 
* グローバル変数を参照しまく数 (入力を再現しにくい)
* タイミングによって出力が変わる
* 等…

## Table Driven Testing

* Go のテストではよく使われる気がする手法
* 入力と期待する出力を構造体で定義し、配列を作ってテーブル化する
  * テストケースを増やすのが楽
  * 人が見てもわかりやすい
  * 似たようなコードをコピペしまくらなくて済む
* 以下はこの前の fizzbuzz 関数のテスト (再掲)

```go
package main

import "testing"

type testData struct {
	input  int    // 入力
	expect string // 期待する出力
}

func TestFizzBuzz(t *testing.T) {
	testDataTable := []testData{ // 入力と期待する出力の表
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
	}

	// ループを回す
	for _, v := range testDataTable {
		if fizzbuzz(v.input) != v.expect {
			t.Errorf("unexpected result. [got] %d, [want] %d", fizzbuzz(v.input), v.expect)
		}
	}
}
```

## カバレッジを計測する

* テスト時にカバレッジ測定のためのパラメータを入れるとカバレッジ計ってくれる

```bash
# テストしたいパッケージ内に cd した後に、
$ go test -coverprofile=cover.out .
$ go tool cover -html=cover.out -o cover.html
# cover.html をブラウザで開く
```

* コードが実行されたかどうかが分かる。
  * テストが手薄 (というかされていない) 部分が分かる
  * 実行されていれば、ひとまず実行されていないよりは安心

## interface を使ったモック (option)

* 関数を分割していったとしても、以下のようなケースはやはりテストしにくい
  * 通信を行う処理 (通信相手を準備しないといけない)
  * 都合の良いところでエラーを返してほしいとき
    * A やって B やって C やる、みたいな関数で B のところでエラー返してほしい、みたいな。
* テスト用に擬似的に処理結果を返すモジュールを作成する (モック)

* 参考
  * [Golangにおけるinterfaceをつかったテスト技法 - SOTA](http://deeeet.com/writing/2016/10/25/go-interface-testing/)



