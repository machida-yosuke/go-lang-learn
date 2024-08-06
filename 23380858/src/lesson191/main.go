package main

import (
	"fmt"
)

func main() {
	// a := make(chan int)
	// b := make(chan int)
	// close(a)
	// close(b)

	// select {
	// case <-a:
	// 	println("a")
	// case <-b:
	// 	println("b")
	// default:
	// 	println("default")
	// }

	// 	ch1 := make(chan int)
	// 	ch2 := make(chan int)

	// 	start := time.Now()
	// 	done := make(chan interface{})
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		close(done)
	// 	}()

	// 	go func() {
	// 		defer close(ch1)
	// 		for i := 0; i < 5; i++ {
	// 			time.Sleep(1 * time.Second)
	// 			ch1 <- i
	// 		}
	// 	}()

	// 	go func() {
	// 		defer close(ch2)
	// 		for i := 0; i < 10; i++ {
	// 			time.Sleep(1 * time.Second)
	// 			ch2 <- i
	// 		}
	// 	}()

	// loop:
	// 	for {
	// 		select {
	// 		case <-done:
	// 			break loop
	// 		case <-time.After(1 * time.Second):
	// 			break loop
	// 		case v, _ := <-ch1:
	// 			println("ch1", v)
	// 		case v, _ := <-ch2:
	// 			println("ch2", v)
	// 		}
	// 	}
	// 	end := time.Now()

	// 	fmt.Println("done", end.Sub(start))

	var ch <-chan int
	select {
	case <-ch:
		fmt.Println("read from ch")
	default:
		fmt.Println("default")
	}

	// ch <- 1
}
