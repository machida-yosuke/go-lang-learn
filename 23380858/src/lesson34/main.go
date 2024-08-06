package main

import "fmt"

// ジェネレーター
// ジェネレーターはクロージャーの一種で、クロージャーの中で無限ループを使って値を生成する関数です。

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
}
