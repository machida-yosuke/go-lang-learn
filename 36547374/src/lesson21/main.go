package main

import (
	"fmt"
	"runtime"
)

func main() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(1 * time.Second)
	// }()

	// fmt.Println(<-ch)

	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	fmt.Println(runtime.NumGoroutine())

	ch2 := make(chan int, 1)

	ch2 <- 2
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)
}
