package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joikoi1212/Go_CRUD/lib/api"

	"net/http"
)

func LoginHandler(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
		return
	}
	if c.Request.Method == http.MethodPost {
		email := c.PostForm("email")
		password := c.PostForm("password")
		log.Println("Username: (login_handler)", email)
		userID := api.Login_api(email, password)
		log.Println("User ID (login_Handler): ", userID)
		if userID == 0 {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid username or password"})
			return
		}
		c.HTML(http.StatusFound, "dashboard.html", gin.H{"userID": userID})
		return
	}
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}
