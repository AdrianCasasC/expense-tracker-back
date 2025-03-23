package routes

import (
	"github.com/AdrianCasasC/expense-tracker-back/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	// Expenses
	server.GET("/expenses", controllers.GetExpenses)
	server.POST("/expenses", controllers.PostExpense)
	server.PATCH("/expenses/:expenseId", controllers.PatchExpense)
	server.DELETE("/expenses/:expenseId", controllers.DeleteExpense)
	// Incomes
	server.GET("/incomes", controllers.GetIncomes)
	server.POST("/incomes", controllers.PostIncome)
	server.PATCH("/incomes/:incomeId", controllers.PatchIncome)
	server.DELETE("/incomes/:incomeId", controllers.DeleteIncome)

}
