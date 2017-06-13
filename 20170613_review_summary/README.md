# 2017/06/13 のメモ

## メニュー

* レビューしていて気づいたこと等

## メソッド名、変数名の命名規則

* ちゃんと決めておけばよかったですね…！
* 一般的かと思われる関数名、変数名の命名規則
  * [うまくメソッド名を付けるための参考情報 - Qiita](http://qiita.com/KeithYokoma/items/2193cf79ba76563e3db6)
  * [うまくクラス名を付けるための参考情報 - Qiita](http://qiita.com/KeithYokoma/items/ee21fec6a3ebb5d1e9a8)
  * [Golang CodeReviewComments 日本語訳 - Qiita](http://qiita.com/knsh14/items/8b73b31822c109d4c497)

### 簡単な原則として

* メソッド名は動詞で始める
* Boolean を示す変数、関数は `is` とか `has` とかで始める
* メモリ確保するやつは `New` で始める
* etc...

## mutex 地獄

* https://github.com/CPSPlatform/CssPOC
* もっとシンプルになる気がします。

## context.WithCancel()

* うまく使うと便利です。
* cancelFunc は結果的に使わなかったときも呼び出す必要がある (context のリーク防止)

## テスト

* https://github.com/CPSPlatform/PostProcessService
* 早く動作させることを心がけるべき
  * できる限りスリープしない
  * モックを使っていくことでファイル IO を避ける
  * 実行が遅いとテストする気が失せていく
* 環境変数を直で扱うとテストしにくくなる
