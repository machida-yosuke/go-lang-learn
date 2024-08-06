package main

import "fmt"

func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "2")

	fmt.Println(5 - 1)

	fmt.Println(2 * 3)
	fmt.Println(10 / 2)
	fmt.Println(10 % 3)

	n := 1
	n += 4
	fmt.Println(n) // 5

	n++
	fmt.Println(n) // 6

	n--
	fmt.Println(n) // 5

	s := "ABC"
	s += "DEF"
	fmt.Println(s) // ABCDEF
}
