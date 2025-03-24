package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joikoi1212/Go_CRUD/lib/api"
)

type registerForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

func RegisterHandler(c *gin.Context) {
	var RegisterForm registerForm
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "templates/register.html", nil)
		return
	}
	if c.Request.Method == http.MethodPost {

		if err := c.Bind(&RegisterForm); err != nil {
			api.Register_api(RegisterForm.Username, RegisterForm.Email, RegisterForm.Password)
			c.Redirect(http.StatusFound, "templates/login.html")
			fmt.Println("Redirected to login page")
			return
		}

		return
	}
	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
}
