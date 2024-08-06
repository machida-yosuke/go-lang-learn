package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	sl := []int{1, 2, 3, 4, 5}
	if len(sl) < 0 {
		fmt.Println(sl[0])
	}
}
