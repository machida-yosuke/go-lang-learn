package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type customConstraints interface {
	~int | float64
}

type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2))
	var i1, i2 NewInt = 1, 2
	fmt.Println(add(i1, i2))

	fmt.Println(min(1, 2))

	m1 := map[string]int{"A": 1, "B": 2, "C": 3}
	m2 := map[int]float32{1: 1.23, 2: 4.56, 3: 7.89}
	fmt.Println(sumValues(m1))
	fmt.Println(sumValues(m2))
}
