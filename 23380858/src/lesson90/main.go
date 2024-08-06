package main

import (
	"fmt"
	"sync"
	"time"
)

var st struct{ A, B, C int }

// sync

var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	mutex.Lock()
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Microsecond)
	fmt.Println(st)
	mutex.Unlock()
}

func main() {
	// mutex = new(sync.Mutex)
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		for i := 0; i < 1000; i++ {
	// 			UpdateAndPrint(i)
	// 		}
	// 	}()
	// }
	// for {
	// }

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("A:", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("B:", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println("C:", i)
		}
		wg.Done()
	}()

	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("終了")
}
