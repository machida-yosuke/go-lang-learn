package main

import (
	"fmt"
	"time"
)

// close

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
	ch1 := make(chan int, 2)
	ch1 <- 100
	close(ch1)

	// ch1 <- 1 // panic: send on closed channel
	// fmt.Println(<-ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)

	i2, ok := <-ch1
	fmt.Println(i2, ok)

	ch2 := make(chan int)
	go reciever("1", ch2)
	go reciever("2", ch2)
	go reciever("3", ch2)

	i3 := 0
	for i3 < 100 {
		ch2 <- i3
		i3++
	}
	close(ch2)
	time.Sleep(3 * time.Second)

}
