package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	flags := log.Lshortfile
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARNING: ", flags)
	errLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)
	warnLogger.Println("This is a warning message")
	errLogger.Fatalln("critical error")
}
