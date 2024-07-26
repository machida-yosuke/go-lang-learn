package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(time.Second) {
			cond.Broadcast()
		}
	}()

	var flag [2]bool

	takeStep := func() {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
		fmt.Println(flag)
	}

	var wg sync.WaitGroup

	p0 := func() {
		defer wg.Done()
		flag[0] = true
		takeStep()

		for flag[1] {
			fmt.Println("p0")
			takeStep()
			flag[0] = false
			takeStep()

			if flag[0] != flag[1] {
				break
			}
			takeStep()

			flag[0] = true
			takeStep()
		}
	}

	p1 := func() {
		defer wg.Done()
		flag[1] = true
		takeStep()

		for flag[0] {
			fmt.Println("p1")
			takeStep()
			flag[1] = false
			takeStep()

			if flag[1] != flag[0] {
				break
			}
			takeStep()

			flag[1] = true
			takeStep()
		}
	}

	wg.Add(2)
	go p0()
	go p1()
	wg.Wait()
}
