package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid, _ := uuid.NewUUID()
	fmt.Println(uuid.String())

	uuid2, _ := uuid.NewRandom()
	fmt.Println(uuid2.String())
}
