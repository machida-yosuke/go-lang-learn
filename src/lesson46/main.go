package main

import (
	"fmt"
)

// append make len cap

func main() {
	sl := []int{1, 2, 3}
	fmt.Println(sl)
	sl = append(sl, 300) // 末尾に追加
	fmt.Println(sl)

	sl = append(sl, 400, 500) // 末尾に複数追加
	fmt.Println(sl)

	sl2 := make([]int, 100)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	sl3 := make([]int, 100, 1001) // 1001まで拡張できる
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3)) // 1001
}
