package main

import (
	"fmt"

	"gotodo/app/controllers"
	"gotodo/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
