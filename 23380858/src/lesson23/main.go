package main

import (
	"fmt"
)

// 定数

const Pi = 3.14
const (
	URL      = "https://www.google.com"
	SiteName = "Google"
)

const (
	A = 1
	B
	C
	D = "A"
)

// var Big int = 9223372036854775807 + 1 // 型を入れるとエラー
const Big = 9223372036854775807 + 1

const (
	c1 = iota // 連番
	c2
	c3
)

func main() {
	fmt.Println(Pi)

	// Pi = 3 // エラー
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D)

	fmt.Printf("%T\n", A)
	fmt.Printf("%T\n", B)
	fmt.Printf("%T\n", C)
	fmt.Printf("%T\n", D)

	fmt.Println(Big - 1)

	fmt.Println(c1, c2, c3)
}
