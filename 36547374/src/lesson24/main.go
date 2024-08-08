package main

import (
	"fmt"
	"sync"
	"time"
)

const bufSize = 5

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, bufSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "hello"
		}
	}()

	for i := 0; i < 3; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("default")
		}
		time.Sleep(1500 * time.Millisecond)
	}
}
