package main

import (
	"fmt"
	"time"
)

func doSomething(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("doSomething complete")
		defer close(completed)

		for {
			select {
			case s := <-strings:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return completed
}

func main() {
	done := make(chan interface{})
	completed := doSomething(done, nil)

	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	<-completed

	fmt.Println("main done")
}
