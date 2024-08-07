package main

import (
	"fmt"
	"time"
)

func main() {
	a := -1
	if a == 0 {
		println("a is 0")
	} else if a > 0 {
		println("positive")
	} else {
		println("negative")
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	// for {
	// 	fmt.Println("Infinite loop")
	// 	time.Sleep(2 * time.Second)
	// }

	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(1 * time.Second)
	}

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop
		default:
			fmt.Printf("%v", i)
		}
	}
	fmt.Printf("\n")

	items := []item{
		{price: 10.0},
		{price: 20.0},
		{price: 30.0},
	}
	for _, item := range items {
		item.price *= 1.1
	}
	fmt.Printf("%+v\n", items)

	for i, _ := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
}

type item struct {
	price float32
}
