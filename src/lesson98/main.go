package main

import (
	"context"
	"fmt"
	"time"
)

// context

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("Start")
	time.Sleep(2 * time.Second)
	fmt.Println("End")
	ch <- "Result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Error")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("Success")
			break L
		}
	}

	fmt.Println("Done")
}
