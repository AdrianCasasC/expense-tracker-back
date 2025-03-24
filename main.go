package main

import (
	"github.com/AdrianCasasC/expense-tracker-back/database"
	"github.com/AdrianCasasC/expense-tracker-back/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	database.ConnectDB()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://another.com"}, // Allowed domains
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "PUT"},        // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},      // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                               // Headers exposed to the client
		AllowCredentials: true,                                                     // Allow cookies/auth headers
	}))

	routes.Routes(server)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := server.Run(":" + port)

	if err != nil {
		return
	}
}
