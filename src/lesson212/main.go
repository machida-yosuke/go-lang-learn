package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(workerNum int, ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d: %d\n", workerNum, num)
		wg.Done()
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= runtime.NumCPU(); i++ {
		go worker(i, ch, &wg)
	}

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(x int) {
			ch <- x * 2
		}(i)
	}

	wg.Wait()
	close(ch)
}
