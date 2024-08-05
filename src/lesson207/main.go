package main

import (
	"fmt"
	"sync"
	"time"
)

func shortProcess(done <-chan interface{}) (bool, error) {
	shortWork := time.NewTicker(1 * time.Second)

	select {
	case <-done:
		return false, fmt.Errorf("short process cancelled")
	case <-shortWork.C:
	}
	return true, nil
}
func longProcess(done <-chan interface{}) (bool, error) {
	longWork := time.NewTicker(5 * time.Second)

	select {
	case <-done:
		return false, fmt.Errorf("long process cancelled")
	case <-longWork.C:
	}
	return true, nil
}

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	// defer close(done)
	go func() {
		time.Sleep(1 * time.Second)
		close(done)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := shortProcess(done); err != nil {
			fmt.Println(err)
			fmt.Println(isDone)
			return
		}
		fmt.Println("short process done")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if isDone, err := longProcess(done); err != nil {
			fmt.Println(err)
			fmt.Println(isDone)
			return
		}
		fmt.Println("long process done")
	}()
	wg.Wait()
}
