package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var lock sync.RWMutex
	var wg sync.WaitGroup

	increment := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("increment")
		count++
		time.Sleep(1 * time.Second)
	}

	read := func(wg *sync.WaitGroup, l sync.Locker) {
		l.Lock()
		defer l.Unlock()
		defer wg.Done()

		fmt.Println("read")
		fmt.Println(count)
		time.Sleep(1 * time.Second)
	}

	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg, &lock)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(&wg, lock.RLocker())
	}
	wg.Wait()

	end := time.Now()

	fmt.Println(end.Sub(start))
}
