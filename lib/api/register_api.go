package api

import (
	"github.com/joikoi1212/Go_CRUD/core"
)

func Register_api(username string, email string, password string) {
	DB := core.InitDB()
	DB.Exec("Insert into users (username, email, pass) values (?, ?, ?)", username, email, password)
}
