package main

import "fmt"

// 可変長数の引数

func sum(s ...int) int {
	n := 0

	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3))

	sl := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(sl...))
}
