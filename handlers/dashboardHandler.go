package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashboardHandler(c *gin.Context) {
	userID, _ := c.Get("userID")
	log.Println("User ID from session: ", userID)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"userID": userID})

}
