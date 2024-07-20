package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// res, _ := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	// fmt.Println(res.StatusCode)
	// fmt.Println(res.Proto)
	// fmt.Println(res.Header["Date"])
	// fmt.Println(res.Header["Content-Type"])

	// fmt.Println(res.Request.Method)
	// fmt.Println(res.Request.URL)

	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(string(body))

	vs := url.Values{}
	vs.Add("a", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("http://example.com", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
