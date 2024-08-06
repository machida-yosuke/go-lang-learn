package main

import "fmt"

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()

	return intStream
}

// func double(x []int) []int {
// 	doubleSlice := make([]int, len(x))
// 	for i, v := range x {
// 		doubleSlice[i] = v * 2
// 	}
// 	return doubleSlice
// }

// func add(x []int) []int {
// 	addSlice := make([]int, len(x))
// 	for i, v := range x {
// 		addSlice[i] = v + 1
// 	}
// 	return addSlice
// }

func double(done <-chan interface{}, intStream <-chan int) <-chan int {
	doubleStream := make(chan int)

	go func() {
		defer close(doubleStream)

		for v := range intStream {
			fmt.Println("double", v)
			select {
			case <-done:
				return
			case doubleStream <- v * 2:
			}
		}
	}()

	return doubleStream
}

func add(done <-chan interface{}, intStream <-chan int) <-chan int {
	addStream := make(chan int)

	go func() {
		defer close(addStream)
		for v := range intStream {
			fmt.Println("add", v)
			select {
			case <-done:
				return
			case addStream <- v + 1:
			}
		}
	}()

	return addStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4, 5)

	for v := range double(done, add(done, intStream)) {
		fmt.Println(v)
	}
}
