package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	ch1 <- 10
	close(ch1)

	v, ok := <-ch1
	fmt.Println(v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 2
	ch2 <- 3
	close(ch2)
	v, ok = <-ch2
	fmt.Println(v, ok)
	v, ok = <-ch2
	fmt.Println(v, ok)
	v, ok = <-ch2
	fmt.Println(v, ok)

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
		time.Sleep(3 * time.Second)
	}

	// nCh := make(chan struct{})
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("started %v\n", i)
	// 		<-nCh
	// 	}(i)
	// }

	// time.Sleep(2 * time.Second)
	// close(nCh)
	// wg.Wait()
	// fmt.Println("finish")
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("write")
		}
		defer close(ch)
	}()
	return ch
}
