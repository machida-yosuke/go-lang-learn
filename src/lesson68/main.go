package main

import "fmt"

type MyInt int

func (mi MyInt) print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt = 10
	fmt.Println(mi)
	mi.print()

	var i int = 10
	fmt.Println(i)
	// fmt.Println(i + mi) // エラー
}
