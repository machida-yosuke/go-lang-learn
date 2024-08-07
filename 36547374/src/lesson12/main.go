package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func funcDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("hello")
}
func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, file := range files {
		out = append(out, strings.TrimSuffix(file, ".csv"))
	}

	return out
}

func fileChecker(name string) (string, error) {
	f, err := os.Open(name)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer f.Close()
	return name, nil
}

func addExt(f func(file string) string, file string) {
	fmt.Println(f(file))
}

func countUp() func(int) int {
	count := 0
	return func(n int) int {
		count += n
		return count
	}
}

func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}

func main() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))

	name, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)

	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))

	f2 := func(file string) string {
		return file + ".csv"
	}

	addExt(f2, "file")

	f3 := multiply()
	fmt.Println(f3(10))

	f4 := countUp()
	for i := 0; i < 10; i++ {
		v := f4(2)
		fmt.Println(v)
	}
}
