package main

import (
	"github.com/AdrianCasasC/expense-tracker-back/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.Routes(server)

	server.Run(":8080")
}