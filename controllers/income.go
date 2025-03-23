package controllers

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/AdrianCasasC/expense-tracker-back/services"
	"github.com/AdrianCasasC/expense-tracker-back/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIncomes(c *gin.Context) {
	incomes, err := services.GetAllIncomes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incomes"})
		return
	}

	c.JSON(http.StatusOK, incomes)
}

func PostIncome(c *gin.Context) {
	var newIncome = models.IncomeDto{}
	if err := c.BindJSON(&newIncome); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
	}

	err := services.CreateIncome(newIncome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert income"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Income added successfully!"})
}

func PatchIncome(c *gin.Context) {
	fieldsMap, err := utils.GetListOfFields(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	var body = models.IncomeDto{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	incomeId := c.Param("incomeId")

	updatedIncome, err := services.UpdateIncome(incomeId, body, fieldsMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to modify income"})
		return
	}

	c.JSON(http.StatusOK, updatedIncome)
}

func DeleteIncome(c *gin.Context) {

	incomeId := c.Param("incomeId")
	deletedIncome, err := services.DeleteIncome(incomeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete income"})
		return
	}

	c.JSON(http.StatusOK, deletedIncome)
}
