package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello, World!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Hello, Go!"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) // 初期値の0, ""が入ってる

	i3 = 100
	s3 = "Hello, World!"
	fmt.Println(i3, s3)

	// 暗黙的な宣言
	i4 := 400
	s4 := "Hello, 400!"
	fmt.Println(i4, s4)

	i4 = 450
	// i4 := 100 // 再定義はエラー
	// s4 = 100 // 肩エラー
	fmt.Println(i4)
	// fmt.Println(s5) // 宣言していない変数はエラー

	outer()

	var s5 string = "not user" // 使ってないのでエラー
}
