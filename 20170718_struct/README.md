# 2017/07/18 のメモ

## メニュー

* struct の扱い諸々

## struct とは

* C 言語で言うところの構造体。
* 特徴 (主に C 言語の構造体との違い)
  * メソッドを生やすことができる
  * interface を実装することによるダックタイピング
  * 埋め込み struct/interface (後述)
  
## 使い方

基本的な使い方は以下のような感じ。

```go

// struct の定義
type Person struct {
    name   string
    age    int
    height float32
    weight float32
}

// メソッドを生やす。receiver は実体。
func (p Person) getName() string {
    return p.name
}

// メソッドを生やす。receiver は実体。
func (p Person) getAge() int {
    return p.age
}

// メソッドを生やす。receiver はポインタ。
func (p *Person) getHeight() float32 {
    return p.height
}

// メソッドを生やす。receiver はポインタ。
func (p *Person) getWeight() float32 {
    return p.weight
}

// stringer interface の実装
func (p *Person) String() string {
    return fmt.Sprintf("i'm person. name:%s, age:%d, height:%f, weight:%f", p.name, p.age, p.height, p.weight)
}

func main() {
    // 何も指定がなければ自動的に全てのメンバが初期化される
    p1 := &Person{}
    fmt.Println(p1)

    // 値を渡して初期化する例。
    // 全部の値を渡さなくても良い。
    p2 := &Person{
        name: "akatsuka",
        height: 180.0,
    }
    fmt.Println(p2)
}
```

## 便利機能 - 埋め込み struct

既に定義した struct を別の struct に埋め込むことができる。
* メソッドを引き継ぐことができる

```go
type Runner struct {
    &Person
    runningSpeed float32
}

func (r *Runner) run() {
    // 走る
}

func main() {
    r1:= &Runner{}

    // 埋め込みされた struct のメソッドを呼び出せる
    // 以下二行は同じ意味
    r1.getName()
    r1.Person.getName()

    // stringer interface は Person が実装しているので、Runner も実装済み扱い
    // 以下二行は同じ意味
    fmt.Println(r1)
    fmt.Println(r1.Person)

    // あくまで構造体を保持しているだけなので、いわゆる継承とは違う。
    // Runner 自身が Person になっているわけではない。
    r2 := &Runner{
        // 明示的に Person メンバを指定して初期化する。
        Person: &Person{
            name:   "akatsuka",
            height: 180.0,
        },  
    } 
    fmt.Println(r2)

    // 以下のようには書けない。
    /*
    r3 := &Runner{
        name:   "akatsuka",
        height: 180.0,
    }
    */
}
```

## 便利機能 - 埋め込み interface

struct には interface も埋め込める。
* とりあえずコンパイル通したいとき
* interface が大量にメソッドの実装を求めるが、一部しか使わないとき

```go
type StringPerson struct {
    name   string
    age    int
    height float32
    weight float32
    fmt.Stringer // String() の実装を要求するやつ
}

func main() {
    // とりあえずコンパイル可能。
    // ただし String() の実装がないので実行時に panic を起こす。
    sp := &StringPerson{}
    fmt.Println(sp)
}
```
