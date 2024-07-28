package main

import (
	"fmt"
	"os"
)

type Result struct {
	Response *os.File
	Error    error
}

func checkFiles(done <-chan interface{}, filename ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)

		for _, file := range filename {
			var result Result

			file, err := os.Open(file)
			result = Result{file, err}

			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()
	return results
}

func main() {
	done := make(chan interface{})
	defer close(done)
	filename := []string{"main.go", "bar.txt", "y.go"}
	for result := range checkFiles(done, filename...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			continue
		}

		fmt.Printf("file: %v\n", result.Response.Name())
	}
}
