package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	fmt.Println("1")
	go func() {
		fmt.Println("4")
		defer fmt.Println("close")
		defer close(ch)
		for i := 0; i < 5; i++ {
			fmt.Printf("write channel %d\n", i)
			ch <- i
		}
	}()

	fmt.Println("2")
	for integer := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("read channel %d\n", integer)
	}
	fmt.Println("3")
}
