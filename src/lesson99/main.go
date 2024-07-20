package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://golang.org/?a=1&b=2#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
	fmt.Println(u.Query())

	url := &url.URL{}
	url.Scheme = "https"
	url.Host = "golang.org"
	q := url.Query()
	q.Set("a", "1")
	q.Set("q", "go")
	q.Set("日本", "だと")
	url.RawQuery = q.Encode()
	fmt.Println(url)
}
