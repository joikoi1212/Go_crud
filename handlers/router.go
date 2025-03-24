package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Rotas
	r.GET("/login", LoginHandler)
	r.POST("/login", LoginHandler)
	r.GET("/register", RegisterHandler)
	r.POST("/register", RegisterHandler)
	return r
}

// r := gin.Default()
// r.GET("/ping", func(c *gin.Context) {
//   c.JSON(http.StatusOK, gin.H{
// 	"message": "pong",
//   })
// })
// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
