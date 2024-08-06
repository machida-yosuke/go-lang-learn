package main

import "fmt"

func main() {
	var fl64 float64 = 2.4

	fl := 3.2 // float64になる

	fmt.Println(fl64)
	fmt.Println(fl + fl64)
	fmt.Printf("%T, %T\n", fl64, fl)

	// float32は基本使わない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	fmt.Println(float64(fl32))

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN
}
