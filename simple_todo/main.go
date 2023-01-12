package main

import (
	"fmt"

	"simple_todo/app/controllers"
	"simple_todo/app/models"
)

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
