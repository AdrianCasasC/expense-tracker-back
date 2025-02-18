package controllers

import (
	"net/http"

	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/gin-gonic/gin"
)

var incomes = []models.Income{
	{ ID: "1", Name: "Salario", Value: 1820.50 },
	{ ID: "2", Name: "Inversiones", Value: 50.25 },
	{ ID: "3", Name: "Rentabilidad", Value: 33.75 },
}

func GetIncomes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, incomes)
}