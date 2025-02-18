package controllers

import (
	"net/http"

	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/gin-gonic/gin"
)

var expenses = []models.Expense{
	{ ID: "1", Name: "Alquiler", Value: 250.0 },
	{ ID: "2", Name: "Gimnasio", Value: 32.50 },
	{ ID: "3", Name: "Comida", Value: 122.90 },
}

func GetExpenses(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, expenses)
}