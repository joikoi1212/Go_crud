package handlers

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	loginForm := LoginForm{}
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "templates/login.html", nil)
		return
	}
	if c.Request.Method == http.MethodPost {

		if err := c.Bind(&loginForm); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		return
	}
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}
