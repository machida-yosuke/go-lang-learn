package main

import "fmt"

// 条件分岐
func main() {
	if b := 100; b == 100 {
		fmt.Println("b is 100")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x)
	}
	fmt.Println(x)

	a := 2
	if a == 2 {
		fmt.Println("a is 2")
		return
	}
	if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1")
	}
	fmt.Println("a is", a) // 早期リターンでこの行は実行されない
}
