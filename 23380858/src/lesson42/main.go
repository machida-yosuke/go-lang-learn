package main

import "fmt"

// panicとrecover

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			fmt.Println("END")
		}
	}()
	panic("ランタイムエラーが発生しました")
	fmt.Println("Start") // この行は実行されない
}
