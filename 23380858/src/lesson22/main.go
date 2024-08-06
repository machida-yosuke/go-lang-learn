package main

import (
	"fmt"
)

// 型変換

func main() {
	// var i int = 1
	// f64 := float64(i)
	// fmt.Println(f64)
	// fmt.Printf("%T\n", f64)

	// f2 := 10.6
	// i2 := int(f2)
	// fmt.Println(i2)
	// fmt.Printf("%T\n", i2)

	// var s string = "100"
	// i, _ := strconv.Atoi(s) // _は破棄できる

	// // var s string = "100.1" // エラーになる
	// // i, err := strconv.Atoi(s)
	// // if err != nil {
	// // fmt.Println("Error")
	// // }

	// fmt.Println(i)

	// var i2 int = 200
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("%T\n", s2)

	var h string = "Hello, World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}
