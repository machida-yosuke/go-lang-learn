package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
	name  string
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		fmt.Println("ロック", v1.name)
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)

		v2.mu.Lock()
		fmt.Println("ロック", v2.name)
		defer v2.mu.Unlock()

		fmt.Println(v1.value + v2.value)
	}

	var a value = value{name: "a"}
	var b value = value{name: "b"}

	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
