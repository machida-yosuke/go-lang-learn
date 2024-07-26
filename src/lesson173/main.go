package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	const timer = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()
		count := 0
		begin := time.Now()

		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(3 * time.Nanosecond)
			lock.Unlock()
			count++
		}

		fmt.Println("greedyWorker: %d", count)
	}

	politeWorker := func() {
		defer wg.Done()
		count := 0
		begin := time.Now()

		for time.Since(begin) <= timer {
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			lock.Lock()
			time.Sleep(1 * time.Nanosecond)
			lock.Unlock()
			count++
		}
		fmt.Println("politeWorker: %d", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
