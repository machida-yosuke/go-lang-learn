package main

import (
	"fmt"
	"sync"
)

var onceA, onceB sync.Once

func A() {
	fmt.Println("A")
	onceA.Do(B)
}

func B() {
	fmt.Println("B")
	onceB.Do(A)
}

func main() {
	// count := 0

	// increment := func() {
	// 	count++
	// }

	// decrement := func() {
	// 	count--
	// }

	var once sync.Once
	// once.Do(increment)
	// once.Do(decrement) // 呼ばれない

	// fmt.Println(count)
	once.Do(A)
}
