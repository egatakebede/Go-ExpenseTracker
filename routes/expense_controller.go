package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateExpenseInput defines what data we expect for creating an expense
type CreateExpenseInput struct {
	Title    string  `json:"title" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Category string  `json:"category"`
	Notes    string  `json:"notes"`
	SpentAt  string  `json:"spent_at"` // ISO string
}

// UpdateExpenseInput defines what can be updated
type UpdateExpenseInput struct {
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Notes    string  `json:"notes"`
	SpentAt  string  `json:"spent_at"`
}

// CreateExpense handler

func CreateExpense(c *gin.Context) {
	var input CreateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
