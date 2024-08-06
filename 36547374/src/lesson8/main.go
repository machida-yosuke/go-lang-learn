package main

import "fmt"

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	i int
	s string
	b bool
)

func main() {
	// var i int
	// var i int = 1
	// var i = 1
	i := 1
	ui := uint16(i)
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	fmt.Println(ui)
	fmt.Printf("i :%[1]T %[1]v ui :%[2]T %[2]v\n", i, ui)

	f := 1.2
	s := "Hello"
	b := true
	fmt.Printf("%T %v\n", f, f)
	fmt.Printf("%T %v\n", s, s)
	fmt.Printf("%T %v\n", b, b)

	pi, title := 3.14, "Circle"
	fmt.Println(pi, title)

	x := 10
	y := 1.23
	// z := x + y // error
	z := float64(x) + y
	fmt.Println(z)

	// secret = "def" // error

	fmt.Println(Mac, Windows, Linux)

	i = 2
	fmt.Println(i)

	i += 1
	fmt.Println(i)

	i *= 2
	fmt.Println(i)
}
