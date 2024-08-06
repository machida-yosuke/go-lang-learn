package main

import "fmt"

// 関数

func plus(x, y int) int { return x + y }

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func double(price int) (result int) {
	result = price * 2
	return
}

func noReturn() {
	fmt.Println("No Return")
	return
}

func main() {
	println(plus(1, 2))
	println(plus(5, 4))
	println(plus(10, 20))
	// println(plus(10, "20")) // エラー

	q, r := Div(19, 10)
	println(q, r)

	q2, _ := Div(19, 10)
	println(q2)

	i4 := double(10)
	println(i4)

	noReturn()
}
