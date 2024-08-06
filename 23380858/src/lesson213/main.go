package main

import (
	"fmt"
	"os"
	"strings"
)

func readGoFile(path string) chan string {
	promise := make(chan string)
	go func() {
		defer close(promise)
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			promise <- string(content)
		}
	}()

	return promise
}

func printFunc(futureSource chan string) chan []string {
	promise := make(chan []string)
	go func() {
		defer close(promise)
		var result []string
		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func main() {
	futureSource := readGoFile("main.go")
	futureFunc := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFunc, "\n"))
}
