package api

import (
	"net/http"

	"github.com/joikoi1212/Go_CRUD/core"
)

func Login_api(w http.ResponseWriter, r *http.Request) {
	DB := core.InitDB()
	DB.Exec("SELECT * FROM users")
}
