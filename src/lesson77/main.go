package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("start")

	// defer func() {
	// 	fmt.Println("defer")
	// }()
	// os.Exit(0)

	// 	_, err := os.Open("test.txt")
	// 	if err != nil {
	// 		log.Fatal(err) // log.Fatal は os.Exit(1) を呼び出す
	// 	}

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])

	// fmt.Printf("length=%d\n", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Printf("os.Args[%d]=%s\n", i, v)
	// }

	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello, World\n"))
	// f.WriteAt([]byte("golang"), 6)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("test")

	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// bs := make([]byte, 128)
	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// bs2 := make([]byte, 128)
	// nn, err := f.ReadAt(bs2, 10) // 10byte目から読み込む
	// fmt.Println(nn)
	// fmt.Println(string(bs2))

	f, err := os.OpenFile("foo.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
