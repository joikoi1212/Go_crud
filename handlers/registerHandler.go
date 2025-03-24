package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joikoi1212/Go_CRUD/lib/api"
)

func RegisterHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "register.html", nil)
		return
	}
	if c.Request.Method == http.MethodPost {

		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		api.Register_api(username, email, password)
		c.HTML(http.StatusOK, "login.html", nil)
		fmt.Println("Redirected to login page")
		return
	}
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}
