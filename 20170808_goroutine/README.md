# 2017/08/08 のメモ

## メニュー

* goroutine
  * goroutine 間でのコミュニケーション
  * 排他制御

## goroutine

* C 言語で言うところの pthread 的な位置づけ。pthread より軽量。
* カジュアルにじゃんじゃん作って良い (とはいえ最大は 65535 個だったような…？) 。
* ただし、
  * 正しく終わらせる必要がある (無限ループしっぱなし → goroutine のリーク)
  * デッドロック、データの競合に注意する
  * 普通にやってたら気づかないであろう罠も若干ある

### コミュニケーション

#### 共有メモリ

* もっとも愚直だが気をつけなければならないことも多い。
* Mutex の出番。

```go
var sharedInt int

func readInt() {
    for {
        fmt.Printf("sharedInt = %d\n", sharedInt)
    }
}

func writeInt() {
    for {
        sharedInt++
        if sharedInt > 100 {
            sharedInt = 0
        }
    }
}

func main() {
    // 読み込み用 goroutine 開始
    go readInt()

    // 書き込み用 goroutine 開始
    writeInt()

    // should not reach
}
```

* `sharedInt` は読み込みと同時に書き込みが発生することがあり、内容が保証されない可能性がある。
* こういう場合は `sync.Mutex` などを用いて排他制御していく必要がある。

#### channel

* channel で goroutine 間のやり取りを行う。

```go
var sharedInt int

func readInt(c chan int) {
    for {
        select {
        case v := <-c:
            fmt.Printf("received int = %d\n", v)
        }
    }
}

func writeInt(c chan int) {
    for {
        sharedInt++
        if sharedInt > 100 {
            sharedInt = 0
        }
        c <- sharedInt
    }
}

func run() {
    c := make(chan int)
    go readInt(c)
    writeInt(c)
}

func main() {
    run()
}

```

* `select` の特徴
  * 複数のケースがあった場合、必ずどれかひとつだけ選択される
  * 上記は channel に値が書き込まれるまでブロックする (待つ) 使い方。
* channel の特徴
  * goroutine セーフ。channel に対しての操作は排他制御が必要ない。
  * int でなくても何でも送れる。文字列でも構造体でも。

### 同期

* goroutine の待ち合わせを行う方法。

#### sync.WaitGroup

```go
func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        // いつ終わるか分からない処理
        // ...

        // 終わったら wg.Done() を呼び出す
        wg.Done()
    }()
    wg.Wait()
}
```

* `wg.Add(n int)` で、待ちカウントを増やす。
* `wg.Done()` で、待ちカウントを減らす。
* `wg.Wait()` は、待ちカウントが 0 になるまでブロックする。

* 乱用するととても見難くなる

#### select と channel

* select と channel でも同期が可能。

```go
func main() {
    doneChan := make(chan struct{})

    go func() {
        // いつ終わるか分からない処理
        // ...

        // 終わったら channel に書き込む
        doneChan <- struct{}{}
    }()
    // 終わるのを待つ
    <- doneChan
}
```
