package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
