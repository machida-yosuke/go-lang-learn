package main

import (
	"fmt"
	"time"
)

func generator() <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 100; i++ {
			intStream <- i
		}
	}()
	return intStream
}

func signal(after time.Duration) <-chan interface{} {
	done := make(chan interface{})
	go func() {
		defer close(done)
		defer fmt.Println("close done")
		time.Sleep(after)
	}()
	return done
}

func orDone(done <-chan interface{}, c <-chan int) <-chan interface{} {
	valCh := make(chan interface{})
	go func() {
		defer close(valCh)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}

				fmt.Println(v)
				select {
				case valCh <- v:
				case <-done:
				}
			}
		}
	}()
	return valCh
}

func main() {
	start := time.Now()
	done := signal(1000 * time.Microsecond)
	intStream := generator()

	// loop:
	// 	for {
	// 		select {
	// 		case <-done:
	// 			break loop
	// 		case v, ok := <-intStream:
	// 			if !ok {
	// 				break loop
	// 			}
	// 			fmt.Println(v)
	// 		}
	// 	}

	for val := range orDone(done, intStream) {
		fmt.Println(val)
	}

	fmt.Println("Elapsed:", time.Since(start))
}
