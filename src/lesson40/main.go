package main

import "fmt"

// ラベル付きfor

func main() {
	// Loop:
	// 	for {
	// 		for {
	// 			for {
	// 				println("start")
	// 				break Loop
	// 			}
	// 			println("break 1")
	// 		}
	// 		println("break 2")
	// 	}
	// 	println("finish")
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("ここは処理したくない")
	}
}
