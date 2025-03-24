package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := gin.Default()

	// Rotas
	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/register", RegisterHandler).Methods("GET")
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	return r
}

// r := gin.Default()
// r.GET("/ping", func(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{
// 	"message": "pong",
//   })
// })
// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
