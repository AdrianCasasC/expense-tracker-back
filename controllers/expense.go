package controllers

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/AdrianCasasC/expense-tracker-back/services"
	"github.com/AdrianCasasC/expense-tracker-back/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*var expenses = []any{}

func GetExpenses(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, expenses)
}*/

func GetExpenses(c *gin.Context) {
	expenses, err := services.GetAllExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func PostExpense(c *gin.Context) {
	var newExpense = models.ExpenseDto{}
	if err := c.BindJSON(&newExpense); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
	}

	err := services.CreateExpense(newExpense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert expense"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense added successfully!"})
}

func PatchExpense(c *gin.Context) {
	fieldsMap, err := utils.GetListOfFields(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	var body = models.ExpenseDto{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	expenseId := c.Param("expenseId")

	updatedExpense, err := services.UpdateExpense(expenseId, body, fieldsMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify expense"})
		return
	}

	c.JSON(http.StatusOK, updatedExpense)
}

func DeleteExpense(c *gin.Context) {

	expenseId := c.Param("expenseId")
	deletedExpense, err := services.DeleteExpense(expenseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete expense"})
		return
	}

	c.JSON(http.StatusOK, deletedExpense)
}
