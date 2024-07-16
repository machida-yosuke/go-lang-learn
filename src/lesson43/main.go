package main

import (
	"fmt"
	"time"
)

// go gorutine

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// sub()

	// for {
	// 	fmt.Println("Main Loop")
	// 	time.Sleep(200 * time.Millisecond)
	// } // sub()が終わるまでMain Loopが表示されない

	go sub()
	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
