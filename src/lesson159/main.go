package main

import (
	"fmt"
)

type Number interface {
	~int | int32 | int64 | float32 | float64
}

type MyInt int

func Max[T Number](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(1.1, 2))
	fmt.Println(Max[int](1, 2))
	fmt.Println(Max[MyInt](1, 2))

	var x, y MyInt = 1, 2
	fmt.Println(Max(x, y))
}
