package main

import "fmt"

// for

func main() {
	sl := []string{"A", "B", "C", "D", "E"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v)
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}
}
