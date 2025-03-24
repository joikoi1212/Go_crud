package main

import (
	"fmt"

	"github.com/joikoi1212/Go_CRUD/handlers"
)

func main() {
	router := handlers.NewRouter()
	router.LoadHTMLGlob("templates/*")
	fmt.Println("Server started at http://localhost:8081/register")
	router.Run(":8081")
}
