package services

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"sort"
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

	return models.CostDto{Expenses: getSameDayCosts(filteredExpenses), Incomes: getSameDayCosts(filteredIncomes)}, nil

}

func filterExpenseByYearAndMonth(slice []models.ExpenseDto, year int, month int) []models.GraphCost {
	if len(slice) == 0 {
		return []models.GraphCost{}
	}

	var filtered []models.GraphCost
	for _, event := range slice {
		if event.Date.Month() == time.Month(month) && event.Date.Year() == year {
			filtered = append(filtered, models.GraphCost{Day: event.Date.Day(), Value: event.Value})
		}
	}

	return filtered
}

func filterIncomeByYearAndMonth(slice []models.IncomeDto, year int, month int) []models.GraphCost {
	if len(slice) == 0 {
		return []models.GraphCost{}
	}

	var filtered []models.GraphCost
	for _, event := range slice {
		if event.Date.Month() == time.Month(month) && event.Date.Year() == year {
			filtered = append(filtered, models.GraphCost{Day: event.Date.Day(), Value: event.Value})
		}
	}

	return filtered
}

func getSameDayCosts(costs []models.GraphCost) []models.GraphCost {
	if len(costs) == 0 {
		return []models.GraphCost{}
	}

	sumMap := make(map[int]float64)
	for _, c := range costs {
		sumMap[c.Day] += c.Value
	}

	// Convert the map back to a slice.
	var result []models.GraphCost
	for day, value := range sumMap {
		result = append(result, models.GraphCost{Day: day, Value: value})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Day < result[j].Day
	})

	return result
}
