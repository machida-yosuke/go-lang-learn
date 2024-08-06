package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング
func main() {
	var s string = "100a"

	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T\n", i)
	}
}
