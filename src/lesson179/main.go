package main

import (
	"fmt"
	"sync"
	"time"
)

// type Person struct {
// 	Name string
// }

func main() {
	// mypool :=
	// 	&sync.Pool{
	// 		New: func() interface{} {
	// 			fmt.Println("Creating new instance.")
	// 			return new(Person)
	// 		},
	// 	}
	// mypool.Put(&Person{Name: "John"})
	// mypool.Put(&Person{Name: "2"})

	// instance1 := mypool.Get()
	// instance2 := mypool.Get()
	// instance3 := mypool.Get().(*Person)
	// instance3.Name = "3"

	// fmt.Println(instance1)
	// fmt.Println(instance2)
	// fmt.Println(instance3)

	// mypool.Put(instance1)
	// mypool.Put(instance2)
	// mypool.Put(instance3)

	// instance4 := mypool.Get()
	// fmt.Println(instance4)

	count := 0

	mypool :=
		&sync.Pool{
			New: func() interface{} {
				count++
				fmt.Println("Creating new instance.")
				return struct{}{}
			},
		}

	mypool.Put("manual add 1")
	mypool.Put("manual add 2")

	var wg sync.WaitGroup
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		time.Sleep(1 * time.Millisecond)
		go func() {
			defer wg.Done()
			instance := mypool.Get()
			mypool.Put(instance)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
