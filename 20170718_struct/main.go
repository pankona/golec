package main

import "fmt"

// struct の定義
type Person struct {
	name   string
	age    int
	height float32
	weight float32
	fmt.Stringer
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
	return fmt.Sprintf("name:%s, age:%d, height:%f, weight:%f", p.name, p.age, p.height, p.weight)
}

type Runner struct {
	Person
	runningSpeed float32
}

func (r Runner) run() {
	// 走る
}

func main() {
	// 何も指定がなければ自動的に全てのメンバが初期化される
	p1 := &Person{}
	fmt.Println(p1)

	// 値を渡して初期化する例。
	// 全部の値を渡さなくても良い。
	p2 := &Person{
		name:   "akatsuka",
		height: 180.0,
	}
	fmt.Println(p2)

	r1 := &Runner{}
	fmt.Println(r1)

	r2 := &Runner{
		Person: Person{
			name:   "akatsuka",
			height: 180.0,
		},
	}
	fmt.Println(r2.getName())
	fmt.Println(r2.Person.getName())
}
