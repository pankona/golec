# 2017/06/02 のメモ

## メニュー

* 環境構築
* fizzbuzz
* fizzbuzz テスト
* GoCodeReviewComment の紹介

## fizzbuzz

### ルール

* 数値を入力してもらい (入力された値を n とする) 、1 から順番に n まで数を表示する
* 表示しようとしている数が
  * 3 で割り切れるなら "Fizz" 
  * 5 で割り切れるなら "Buzz"
  * 両方で割り切れるなら "FizzBuzz" と表示する

### 実行例

```
$ ./fizzbuzz 34
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
```

### サンプルコード

コマンドライン引数を取るところまでのコードは以下。

```go
package main

import (
	"log"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	// ...

}
```

## fizzbuzz テスト

* 作った処理をテストするように書いてみる。
* テストしにくかったらテストしやすいように変更していく。
* テストコードの書き方は以下のような感じ。

```go
package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	// ...
}
```

## GoCodeReviewComment

* Effective Go の軽い版みたいなやつ
* 日本語訳されたものがあります
  * [#golang CodeReviewComments 日本語翻訳 - Qiita](http://qiita.com/knsh14/items/8b73b31822c109d4c497)

