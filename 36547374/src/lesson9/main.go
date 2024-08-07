package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ui1 uint16
	fmt.Printf("ui1 pointer:%p\n", &ui1)
	var ui2 uint16
	fmt.Printf("%p\n", &ui2)

	var p1 *uint16
	fmt.Printf("%v\n", p1)

	p1 = &ui1
	fmt.Printf("p1 value%v\n", p1)

	fmt.Printf("%d[byte]\n", unsafe.Sizeof(p1))
	fmt.Printf("p1 point:%p\n", &p1)
	fmt.Printf("p1 value:%v\n", *p1)

	*p1 = 1
	fmt.Printf("ui1 value:%v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("pp1 value:%v\n", pp1)
	fmt.Printf("pp1 point:%p\n", &pp1)
	fmt.Printf("pp1 byte:%d\n", unsafe.Sizeof(pp1))

	fmt.Printf("pp1 value:%v\n", *pp1)
	fmt.Printf("pp1 value:%v\n", **pp1)

	**pp1 = 2
	fmt.Printf("ui1 value:%v\n", ui1)

	ok, result := true, "A"
	if ok {
		result := "B"
		fmt.Printf("result:%p\n", &result)
		println(result)
	} else {
		result := "C"
		fmt.Printf("result:%p\n", &result)
		println(result)
	}
	fmt.Println(result)
	fmt.Printf("result:%p\n", &result)
}
