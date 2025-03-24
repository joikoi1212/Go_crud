package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type EmailIDHandler struct {
	Log *log.Logger
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	test()
	loginTemplate := template.Must(template.ParseFiles(filepath.Join("templates", "login.html")))

	if req.Method == http.MethodGet {
		loginTemplate.Execute(res, nil)
		return
	}
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		fmt.Fprintf(res, "Username: %s, Password: %s", username, password)
		return
	}
	http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)

}

func test() {
	panic("unimplemented")
}

func (s *EmailIDHandler) test() {
	s.Log.Println("test")
}
