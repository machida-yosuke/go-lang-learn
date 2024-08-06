package main

import (
	"fmt"
	"sync"
)

func main() {
	// ch := make(chan string)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- "Hello!"
	// }()

	// v, ok := <-ch
	// fmt.Println(v, ok)

	// close(ch)

	// v, ok := <-ch
	// fmt.Println(v, ok)

	// ch := make(chan int)

	// go func() {
	// 	defer close(ch)
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		ch <- i
	// 	}
	// }()

	// for integer := range ch {
	// 	fmt.Println(integer)
	// }

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		fmt.Printf("start goroutine %d\n", i)

		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("begin %d\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}
