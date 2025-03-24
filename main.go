package main

import (
	"fmt"

	"github.com/joikoi1212/Go_CRUD/handlers"
)

func main() {
	router := handlers.NewRouter()
	fmt.Println("Server started at http://localhost:8081/register")
	router.Run()
}

// func (h *EmailIDHandler) Routes(rg *gin.RouterGroup) {

// 	rg.POST("/", middlewares.ValidateRequest(&request.Email{}), h.sendEmail)
