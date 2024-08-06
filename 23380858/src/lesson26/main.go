package main

import "fmt"

// 論理演算子

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || false)
	fmt.Println(true || false == true)
	fmt.Println(false || false == true)
	fmt.Println(!true)

	fmt.Println(!true == false)
}
