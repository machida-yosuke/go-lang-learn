package main

import "fmt"

// クロージャー

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}
func main() {
	f := later()
	fmt.Printf("%s\n", f("hello"))
	fmt.Printf("%s\n", f("go"))
}
