package main

import (
	"github.com/AdrianCasasC/expense-tracker-back/database"
	"github.com/AdrianCasasC/expense-tracker-back/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	server := gin.Default()

	routes.Routes(server)

	err := server.Run(":8080")

	if err != nil {
		return
	}
}
