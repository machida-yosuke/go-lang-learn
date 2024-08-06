package main

import (
	"fmt"
	"lesson75/alib"
)

// テスト

func IsOne(n int) bool {
	return n == 1
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(2))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
