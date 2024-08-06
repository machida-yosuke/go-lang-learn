package main

import "fmt"

// map

func main() {
	var m = map[string]int{
		"apple":  100,
		"banana": 200,
	}

	fmt.Println(m)

	m2 := map[string]int{
		"apple":  200,
		"banana": 400,
	}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)
	fmt.Println(m3[1])

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "Japan"
	m4[81] = "America"
	fmt.Println(m4[81])

	s, ok := m4[1]
	fmt.Println(s, ok)

	s1, ok := m4[10]
	fmt.Println(s1, ok)

	delete(m4, 1)
	fmt.Println(m4)

	delete(m4, 81)
	fmt.Println(m4)

	fmt.Println(len(m3))
}
