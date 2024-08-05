package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func shortProcess(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	switch isDone, err := doSomething(ctx); {
	case err != nil:
		return false, err
	case isDone:
		return isDone, nil
	}
	return false, fmt.Errorf("UnSupported case")
}

func longProcess(ctx context.Context) (bool, error) {
	longWork := time.NewTicker(5 * time.Second)

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-longWork.C:
	}
	return true, nil
}

func doSomething(ctx context.Context) (bool, error) {
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println("deadline:", deadline)
		if deadline.Sub(time.Now().Add(2*time.Second)) <= 0 {
			return false, context.DeadlineExceeded
		}
	}

	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case <-time.After(3 * time.Second):
	}
	return true, nil
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cancel()
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := shortProcess(ctx); err != nil {
			fmt.Println(err)
			fmt.Println(isDone)
			cancel()
			return
		}
		fmt.Println("short process done")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := longProcess(ctx); err != nil {
			fmt.Println(err)
			fmt.Println(isDone)
			cancel()
			return
		}
		fmt.Println("long process done")
	}()
	wg.Wait()
}
