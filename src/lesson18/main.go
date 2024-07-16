package main

import (
	"fmt"
)

// 配列
func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	fmt.Println(arr3[0])

	arr4 := [...]string{"a", "b", "c", "d"} // 配列の要素数を省略
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])
	// fmt.Println(arr2[3]) // エラー

	arr2[2] = "c"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1 // エラー
	// fmt.Println(arr5)

	fmt.Println(len(arr2))
}
