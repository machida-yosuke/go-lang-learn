package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bt := true
	// fmt.Println(strconv.FormatBool(bt))

	// i := strconv.FormatInt(-100, 10) // 10進数
	// fmt.Println(i)

	// i2 := strconv.Itoa(100)
	// fmt.Println(i2)
	// fmt.Printf("%T\n", i2)

	// // fmt.Println(strconv.FormatFloat(12.34, 'E', -1, 64))

	// bt1, ok := strconv.ParseBool("true")
	// if ok != nil {
	// 	fmt.Println("error")
	// }
	// fmt.Println(bt1)
	// fmt.Printf("%T\n", bt1)

	// bf2, _ := strconv.ParseBool("0")
	// bf3, _ := strconv.ParseBool("f")
	// bf4, _ := strconv.ParseBool("F")
	// bf5, _ := strconv.ParseBool("FALSE")
	// bf6, _ := strconv.ParseBool("False")
	// bf7, _ := strconv.ParseBool("1")

	// fmt.Println(bf2, bf3, bf4, bf5, bf6, bf7)

	// i3, _ := strconv.ParseInt("100", 10, 0)
	// fmt.Println(i3)
	// i4, _ := strconv.ParseInt("-1", 10, 0)
	// fmt.Println(i4)

	fl1, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(fl1)
	fl2, _ := strconv.ParseFloat("0.2", 64)
	fmt.Println(fl2)

	fl3, _ := strconv.ParseFloat("-2", 64)
	fmt.Println(fl3)
}
