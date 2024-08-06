package main

import "fmt"

// 型スイッチ

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + " is string")
	case int:
		fmt.Println(v, "is int")
	default:
		fmt.Println("I don't know")
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3

	// i, isInt := x.(int)
	// fmt.Println(i, isInt)

	// f, isFloat64 := x.(float64)
	// fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "is int")
	} else if f, isFloat64 := x.(float64); isFloat64 {
		fmt.Println(f, "is float64")
	} else if f, isString := x.(string); isString {
		fmt.Println(f, "is isString")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}

}
