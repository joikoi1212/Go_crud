package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashboardHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "register.html", nil)
		return
	}
	if c.Request.Method == http.MethodPost {

		message := c.PostForm("message")

		api.crud_operations()
		c.HTML(http.StatusOK, "login.html", nil)
		fmt.Println("Redirected to login page")
		return
	}
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}
