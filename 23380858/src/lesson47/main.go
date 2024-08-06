package main

import "fmt"

// copy

func main() {
	sl := []int{1, 2, 3, 4, 5}
	sl2 := sl

	sl2[0] = 100 // メモリ参照なので、sl2[0] = 100, sl[0] = 100になる

	fmt.Println(sl)
	fmt.Println(sl2)

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i)
	fmt.Println(i2)

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)

	copy(sl4, sl3)
	sl4[0] = 100

	fmt.Println(sl3)
	fmt.Println(sl4)
}
