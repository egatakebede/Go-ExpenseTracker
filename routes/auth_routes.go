package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("auth")
	{
		auth.POST("/auth", controllers.Signup)
		auth.POST("/login", controllers.Login)
	}
}
