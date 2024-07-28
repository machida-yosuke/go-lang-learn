package main

var data []int = []int{1, 2, 3, 4, 5}

// func writeToChan(writeChan chan<- int) {
// 	defer close(writeChan)
// 	for _, v := range data {
// 		writeChan <- v
// 	}
// }

func chanOwner() <-chan int {
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 1; i <= 5; i++ {
			results <- i
		}
	}()

	return results
}

func consumer(results <-chan int) {
	for results := range results {
		println(results)
	}
}

func main() {
	// handleData := make(chan int)
	// go writeToChan(handleData)
	// for integer := range handleData {
	// 	println(integer)
	// }

	results := chanOwner()
	consumer(results)
}
