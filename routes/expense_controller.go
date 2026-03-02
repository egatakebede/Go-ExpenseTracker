package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisteredExpenseRoutes(r *gin.Engine) {
	expense := r.Group("expense")
	{
		expense.POST("/", controllers.CreateExpense)
		expense.GET("?", controllers.getExpenses)
		expense.PUT("/:id", controllers.UpdateExpenses)
		expense.DELETE("/:id", controllers.DeleteExpenses)
	}
}
