package api

import (
	"log"

	"github.com/joikoi1212/Go_CRUD/core"
)

func Login_api(email string, password string) int {
	DB := core.InitDB()
	var userID int
	DB.QueryRow("SELECT id FROM users where email = ?", email).Scan(&userID)
	log.Println("User ID login_api: ", userID)
	return userID
}
