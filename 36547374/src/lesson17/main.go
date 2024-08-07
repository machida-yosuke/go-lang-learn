package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func divide(x, y int) float32 {
	if y == 0 {
		return 0
	}
	return float32(x) / float32(y)
}

func main() {
	x, y := 10, 5
	fmt.Println(Add(x, y))
	fmt.Println(divide(x, y))
}
