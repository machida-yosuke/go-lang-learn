package main

import "fmt"

func main() {
	var i int = 100

	// int8は-128~127までの値しか入らない
	var i8 int8 = 127
	// int16は-32768~32767までの値しか入らない
	var i16 int16 = 32767
	// int32は-2147483648~2147483647までの値しか入らない
	var i32 int32 = 2147483647
	// int64は-9223372036854775808~9223372036854775807までの値しか入らない
	var i64 int64 = 9223372036854775807

	fmt.Println(i, i8, i16, i32, i64)
	// fmt.Printf(i + i8) // 型が違うのでエラー

	// 型を調べる
	fmt.Printf("%T\n", i8) // int8

	// 型変換
	fmt.Printf("%T\n", int32(i8)) // int32

	fmt.Println(i64 - int64(i8)) //変換したのでエラーにならない
}
