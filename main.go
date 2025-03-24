package main

import (
	"fmt"
	"net/http"

	"github.com/joikoi1212/Go_CRUD/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	fmt.Println("Server started at http://localhost:8081/register")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// func (h *EmailIDHandler) Routes(rg *gin.RouterGroup) {

// 	rg.POST("/", middlewares.ValidateRequest(&request.Email{}), h.sendEmail)
