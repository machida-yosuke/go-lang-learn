package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力から読み込む
	scanner := bufio.NewScanner(os.Stdout)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
	}
}
