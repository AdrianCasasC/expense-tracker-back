package routes

import (
	"github.com/AdrianCasasC/expense-tracker-back/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/expenses", controllers.GetExpenses)
	server.POST("/expenses", controllers.PostExpense)
	server.PATCH("/expenses/:expenseId", controllers.PatchExpense)
	server.DELETE("/expenses/:expenseId", controllers.DeleteExpense)
	//server.GET("/incomes", controllers.GetIncomes)
}
