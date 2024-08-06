package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doSomething(done chan interface{}) <-chan int {
	readStream := make(chan int)

	go func() {
		defer close(readStream)
		defer fmt.Println("doSomething done")
		for {
			select {
			case readStream <- rand.Intn(100):
			case <-done:
				return
			}
		}
	}()

	return readStream
}

func main() {
	done := make(chan interface{})
	readStream := doSomething(done)

	for i := 0; i < 3; i++ {
		fmt.Println(<-readStream)
	}

	close(done)

	time.Sleep(1 * time.Second)

	fmt.Println("main done")
}
