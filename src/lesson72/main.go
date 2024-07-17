package main

import "fmt"

// stringerインターフェース

// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{A: 10, B: "文字列"}
	fmt.Println(p)
}
