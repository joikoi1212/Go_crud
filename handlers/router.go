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
	r.GET("/dashboard", DashboardHandler)
	return r
}
