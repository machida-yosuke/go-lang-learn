package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hello from goroutine")
	// }()
	// wg.Wait()
	// fmt.Printf("num %d\n", runtime.NumGoroutine())
	// fmt.Println("finish")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}

	defer trace.Stop()
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("the number of logical CPU cores:", runtime.NumCPU())

	// Task(ctx, "task1")
	// Task(ctx, "task2")
	// Task(ctx, "task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTack(ctx, &wg, "task1")
	go cTack(ctx, &wg, "task2")
	go cTack(ctx, &wg, "task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}

func Task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func cTack(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)

}
