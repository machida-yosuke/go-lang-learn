package main

import (
	"fmt"
)

// interface
// 全ての型と互換性がある

func main() {
	var x interface{}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 1
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 2
	// fmt.Println(x + 2) // エラーになる
}
