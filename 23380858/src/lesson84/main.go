package main

import (
	"log"
	"os"
)

func main() {
	// log.SetOutput(os.Stdout)

	// log.Print("Log\n")
	// log.Println("Log2")
	// log.Printf("Log3%d\n", 3)

	// log.Fatal("Log\n")
	// log.Fatalln("Log2")
	// log.Fatalf("Log%d\n", 3)

	// log.Panic("Log\n")
	// log.Panicln("Log2")
	// log.Panicf("Log%d\n", 3)

	// f, err := os.Create("test.log")
	// if err != nil {
	// 	return
	// }
	// log.SetOutput(f)
	// log.Println("ファイルに書き込み")

	// log.SetOutput(os.Stdout)
	// log.SetFlags(log.LstdFlags)
	// log.Println("A")

	// log.SetFlags(log.LstdFlags | log.Ltime | log.Lmicroseconds)
	// log.Println("B")

	// log.SetFlags(log.Ltime | log.Lshortfile)
	// log.Println("C")

	// log.SetFlags(log.LstdFlags)
	// log.SetPrefix("[LOG]")
	// log.Println("E")

	// fmt.Println("F")

	logger := log.New(os.Stdout, "[Log]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	_, err := os.Open("test.t")
	if err != nil {
		logger.Fatalln("exit", err)
	}
}
