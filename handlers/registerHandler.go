package handlers

import (
	"fmt"
	"net/http"

	"github.com/joikoi1212/Go_CRUD/lib/api"
)

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		http.ServeFile(res, req, "templates/register.html")
		return
	}
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		email := req.FormValue("email")
		password := req.FormValue("password")
		fmt.Println(res, "Email: %s, Username: %s, Password: %s", email, username, password)
		api.Register_api(username, email, password)
		http.Redirect(res, req, "http://localhost:8081/login", http.StatusFound)

		return
	}

}
