package main

import "fmt"

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	byteB := []byte("HI")
	fmt.Println(byteB)

	fmt.Println(string(byteB))
}
