package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "ABCDE")
	fmt.Println(h.Sum(nil))

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
