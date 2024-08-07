package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println(a1)
	var a2 = [3]int{1, 2, 3}
	fmt.Println(a2)
	a3 := [...]int{1, 2, 3}
	fmt.Println(a3)
	fmt.Println(len(a3), cap(a3))
	fmt.Printf("a3: %T\n", a3)
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)

	var s1 []int
	s2 := []int{}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1, len(s1), cap(s1))

	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Println(s1, len(s1), cap(s1))

	s4 := make([]int, 0, 2)
	fmt.Println(s4, len(s4), cap(s4))

	s4 = append(s4, 1, 2, 3, 4)
	fmt.Println(s4, len(s4), cap(s4))

	s5 := make([]int, 4, 6)
	fmt.Println(s5, len(s5), cap(s5))

	s6 := s5[1:3]
	fmt.Println(s6, len(s6), cap(s6))
	s6[1] = 10
	fmt.Println(s5, s6)

	s6 = append(s6, 2)
	fmt.Println(s6, len(s6), cap(s6))

	sc6 := make([]int, len(s5[1:3]))
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(sc6, len(sc6), cap(sc6))
	copy(sc6, s5[1:3])
	fmt.Println(sc6, len(sc6), cap(sc6))

	sc6[1] = 12
	fmt.Println(s5, sc6)

	s5 = make([]int, 4, 6)
	fs6 := s5[1:3:3]
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(fs6, len(fs6), cap(fs6))

	fs6[0] = 6
	fs6[1] = 7
	fs6 = append(fs6, 8)
	fmt.Println(s5)
	fmt.Println(fs6)

	s5[3] = 9
	fmt.Println(s5)
	fmt.Println(fs6)

	var m1 map[string]int
	fmt.Println(m1)

	m2 := map[string]int{}
	fmt.Println(m2)
	fmt.Println(m1 == nil)
	fmt.Println(m2 == nil)

	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 0
	fmt.Println(m2, len(m2), m2["C"])

	delete(m2, "A")
	fmt.Println(m2, len(m2))
	fmt.Println(m2["A"])

	v, ok := m2["A"]
	fmt.Println(v, ok)
	v, ok = m2["C"]
	fmt.Println(v, ok)

	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
