package main

import "fmt"

type T[A any, B []C, C *A] struct {
	a A
	b B
	c C
}

func (t T[A, B, C]) m(x int) {}

func main() {
	var a any = 1
	a = "a"
	a = 1.1
	a = true
	fmt.Println(a)
}
