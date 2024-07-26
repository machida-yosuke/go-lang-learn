package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for _, name := range []string{"a", "b", "c"} {
		go func(name string) {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println("Wait")
			cond.Wait()
			fmt.Println(name)
		}(name)
	}

	fmt.Println("ready..")
	time.Sleep(2 * time.Second)
	fmt.Println("Go!")

	// for i := 0; i < 3; i++ {
	// 	cond.Signal()
	// 	time.Sleep(time.Second)
	// }
	cond.Broadcast()
	time.Sleep(time.Second)
	fmt.Println("Done")
}
