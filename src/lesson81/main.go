package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max int
		msg string
		x   bool
	)

	flag.IntVar(&max, "n", 100, "max value")
	flag.StringVar(&msg, "m", "", "message")
	flag.BoolVar(&x, "x", false, "bool value")

	flag.Parse()

	fmt.Println(max, msg, x)
}
