package main

import "fmt"

type Vector[T any] []T

type IntVector Vector[int]

func main() {
	var v Vector[int] = Vector[int]{1, 2, 3, 4, 5}
	fmt.Println(v)

	var v2 Vector[float32] = Vector[float32]{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(v2)

	v3 := IntVector{1, 2, 3, 4, 5}
	fmt.Println(v3)
}
