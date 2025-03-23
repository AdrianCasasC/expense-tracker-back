package controllers

import (
	"github.com/AdrianCasasC/expense-tracker-back/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCostsByYearAndMonth(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return
	}
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return
	}

	costs, err := services.GetCostsByYearAndMonth(yearInt, monthInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch costs"})
		return
	}

	c.JSON(http.StatusOK, costs)
}
