package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// var count int
	// var mu sync.Mutex

	// increment := func() {
	// 	mu.Lock()
	// 	defer mu.Unlock()
	// 	count++
	// }

	// increment()

	// fmt.Println(count)

	var count int64
	increment := func() {
		atomic.AddInt64(&count, 1)
	}

	increment()
	fmt.Println(count)
}
