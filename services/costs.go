package services

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"time"
)

func GetCostsByYearAndMonth(year int, month int) (models.CostDto, error) {
	expenses, err := GetAllExpenses()
	if err != nil {
		return models.CostDto{}, err
	}

	incomes, err := GetAllIncomes()
	if err != nil {
		return models.CostDto{}, err
	}
	filteredExpenses := filterExpenseByYearAndMonth(expenses, year, month)
	filteredIncomes := filterIncomeByYearAndMonth(incomes, year, month)
	return models.CostDto{Expenses: filteredExpenses, Incomes: filteredIncomes}, nil

}

func filterExpenseByYearAndMonth(slice []models.ExpenseDto, year int, month int) []models.GraphCost {
	var filtered []models.GraphCost
	for _, event := range slice {
		if event.Date.Month() == time.Month(month) && event.Date.Year() == year {
			filtered = append(filtered, models.GraphCost{Day: event.Date.Day(), Value: event.Value})
		}
	}
	if len(filtered) == 0 {
		return []models.GraphCost{}
	}
	return filtered
}

func filterIncomeByYearAndMonth(slice []models.IncomeDto, year int, month int) []models.GraphCost {
	var filtered []models.GraphCost
	for _, event := range slice {
		if event.Date.Month() == time.Month(month) && event.Date.Year() == year {
			filtered = append(filtered, models.GraphCost{Day: event.Date.Day(), Value: event.Value})
		}
	}
	if len(filtered) == 0 {
		return []models.GraphCost{}
	}
	return filtered
}
