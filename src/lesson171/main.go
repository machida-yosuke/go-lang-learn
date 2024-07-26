package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var memoryAccess sync.Mutex
	var data int

	wg.Add(1)
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
		defer wg.Done()
	}()

	wg.Wait()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(data)
	}
	memoryAccess.Unlock()
}
