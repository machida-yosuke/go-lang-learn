package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup
	// m := map[string]int{"a": 1, "b": 2}

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		m["a"] = rand.Intn(100)
	// 		m["b"] = rand.Intn(100)
	// 		defer wg.Done()
	// 	}()

	// 	go func() {
	// 		m["a"] = rand.Intn(100)
	// 		m["b"] = rand.Intn(100)
	// 		defer wg.Done()
	// 	}()
	// }

	// wg.Wait()

	smap := sync.Map{}
	smap.Store("h", "w")
	smap.Store(1, 2)

	smap.Delete(1)

	v, ok := smap.Load("h")
	if ok {
		fmt.Println(v)
	}

	smap.LoadOrStore("h", "w")
	smap.LoadOrStore("a", "b")

	smap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
