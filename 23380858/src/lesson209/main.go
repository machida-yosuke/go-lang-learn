package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func shortProcess(ctx context.Context) (bool, error) {
	shortWork := time.NewTicker(1 * time.Second)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		return false, fmt.Errorf("short process cancelled")
	case <-shortWork.C:
	}
	return true, nil
}

func longProcess(ctx context.Context) (bool, error) {
	longWork := time.NewTicker(5 * time.Second)

	select {
	case <-ctx.Done():
		return false, fmt.Errorf("long process cancelled")
	case <-longWork.C:
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
