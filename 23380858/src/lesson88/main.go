package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("foo.txt")
	bs, _ := io.ReadAll(f)
	fmt.Println(string(bs))

	if err := os.WriteFile("bar.txt", bs, 0666); err != nil {
		fmt.Fprintln(os.Stderr, "書き込みエラー:", err)
	}
}
