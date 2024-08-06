package main

import "fmt"

// for
func main() {
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("loop", i)
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue // 3の時はスキップ
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := [3]int{1, 2, 3}
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// sl := []string{"Python", "PHP", "Golang"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{"aapl": 277, "goog": 2000}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
