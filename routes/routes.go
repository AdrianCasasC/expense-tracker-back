package routes

import (
	"github.com/AdrianCasasC/expense-tracker-back/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	server.GET("/expenses", controllers.GetExpenses)
	server.GET("/incomes", controllers.GetIncomes)
}