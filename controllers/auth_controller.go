package controllers

import (
	"backend/models"
	"expense-tracker/models"
	"net/http"

	"github.com/gin-gonic"
	"github.com/gin-gonic/gin"
)

var jwtkey = []byte("supersecretkey")

func Signup(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
}
