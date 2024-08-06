package main

import (
	"fmt"
	"os"
)

// defer

func testDefer() {
	defer fmt.Println("END") // 関数の最後に実行される
	fmt.Println("START")
}

func runDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {

	testDefer()

	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	runDefer() // 3 2 1 が出力される

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("エラーです")
	}
	defer file.Close()
	file.Write([]byte("hello"))
}
