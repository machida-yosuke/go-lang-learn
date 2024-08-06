package main

// 無名関数

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	println(i)

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)

	println(i2)
}
