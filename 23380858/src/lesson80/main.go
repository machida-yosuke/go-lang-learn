package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("表示")
	// fmt.Print("hello")
	// fmt.Printf("world")

	// fmt.Println("hello", "world")
	// fmt.Println("hello", "world!")

	// fmt.Printf("%s\n", "hello")
	// fmt.Printf("%#v\n", "hello")

	// s := fmt.Sprint("hello")
	// s1 := fmt.Sprintf("%v\n", "hello")
	// s2 := fmt.Sprintln("hello")

	// fmt.Println(s)
	// fmt.Println(s1)
	// fmt.Println(s2)

	fmt.Fprint(os.Stdout, "hello")
	fmt.Fprintf(os.Stdout, "world")
	fmt.Fprintln(os.Stdout, "golang")

	f, _ := os.Create("test1.txt")
	defer f.Close()
	fmt.Fprintln(f, "hello")
}
