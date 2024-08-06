package main

import (
	"fmt"
)

// for

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			fmt.Println(name, "is done.")
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name, "is END.")
}

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
