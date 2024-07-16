package main

import (
	"fmt"
)

// select

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
	// ch1 := make(chan int, 3)
	// ch2 := make(chan string, 2)
	// ch2 <- "A"
	// ch2 <- "B"
	// ch1 <- 1
	// ch1 <- 2

	// // v1 := <-ch1
	// // v2 := <-ch2
	// // fmt.Println(v1)
	// // fmt.Println(v2)

	// select {
	// case v1 := <-ch1:
	// 	fmt.Println(v1 * 100)
	// case v2 := <-ch2:
	// 	fmt.Println(v2 + "!?") // ランダムに処理される
	// case v2 := <-ch2:
	// 	fmt.Println(v2 + "**") // ランダムに処理される
	// default: // どのcaseも処理されなかった場合
	// 	fmt.Println("default")
	// }

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L
			}
		}
	}
}
