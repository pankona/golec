# 2017/06/13 のメモ

## メニュー

* interface を使ってモックする方法

## interface ってそもそもなに？

* メソッドの集まりの定義。

```go
type mkdirer interface {
    MkDir(path string) error
}
```

以下のような特徴がある。
* interface に定義されている関数を struct に実装させる。
* interface が定義している関数の実装を満たしていれば、  
その struct は interface として振る舞うことができるようになる。

```go
// mkdirer を実装した struct (本物の実装)
type realMkDirer struct {}
func (r *realMkdirer) MkDir(path string) error {
    // 実際に処理を行う
    return os.MkdirAll(path, 0755)
}


// mkdirer を実装した struct （モック、偽物の実装）
type mockMkDirer struct {}
func (m *mockMkdirer) MkDir(path string) error {
    // 何もしない
    return nil
}
```

## interface で実装を入れ替える

* 上記の構造体はどちらも `mkdirer interface` として扱うことができる (交換可能) 。

```go

// mkdirer を引数にとって mkdir する関数
func mkdirFunc(m mkdirer, path string) error {
    return m.MkDir(path)
}

main() {
    // ↓ の二例のように交換が可能。

    // 本当にディレクトリを作るとき
    err := mkdirFunc(&realMkDirer{}, "path")
    if err != nil {
        fmt.Println(err)
    }

    // モック（ディレクトリ作らないで済ます）のとき
    err := mkdirFunc(&mockMkDirer{}, "path")
    if err != nil {
        fmt.Println(err)
    }
}
```

## テストで使うパターン

* テストは高速にぶん回したいとする。可能であれば並列実行したいという気持ちであるとする。

### ファイル I/O をモックする

* ファイルのI/Oが発生する処理がいっぱいある場合、I/O がボトルネックになってぶんまわしにくい。
  * ファイル生成 → ファイル存在確認 → ファイル削除、みたいなテストが大量にあったら？
  * ディレクトリ作成、削除を確認するテストが大量にあったら？

* モック実装は実際にファイルI/Oは行わず（行ったことにして）、それ以外のロジックを確認する。
  * 実際にファイルI/Oをしないので速い
  * 他のテストとファイルが衝突する、みたいなことも起こらない

* とはいえ、実際に I/O を行うテストも少しはやっとく

### HTTP 通信をモックする

* HTTP 通信を行う場合
  * 対向を用意する必要があるのが手間
  * テストを並列でやろうと思ったら？バインドするポートが重複したりってことも…。

* モック実装は実際に通信を行わず（レスポンスがきたことにして）、それ以外のロジックを確認する。

## TO BE CONTINUED...
