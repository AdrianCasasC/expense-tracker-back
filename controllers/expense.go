package controllers

import (
	"context"
	"github.com/AdrianCasasC/expense-tracker-back/database"
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

/*var expenses = []any{}

func GetExpenses(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, expenses)
}*/

func GetExpenses(c *gin.Context) {
	var expensesCollection = database.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := expensesCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	defer cursor.Close(ctx)

	var expenses []models.ExpenseDto
	if err = cursor.All(ctx, &expenses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

/*
import (
	"context"
	"github.com/AdrianCasasC/expense-tracker-back/database"
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var expenseCollection = database.GetCollection("expenses")

func GetExpenses(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := expenseCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}
	defer cursor.Close(ctx)

	var expenses []models.ExpenseDto
	if err = cursor.All(ctx, &expenses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

// TODO: Hacer la transformación de Dto a entity para meterlo en base de datos
func CreateExpense(c *gin.Context) {
	expenseCollection := database.GetCollection("expenses")
	var expense models.ExpenseDto
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID and Date
	expense.ID = primitive.NewObjectID()
	expense.Date = time.Now().Format(time.RFC3339)

	// Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := expenseCollection.InsertOne(ctx, expense)
	if err != nil {
		log.Println("Error inserting expense:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert expense"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Expense added successfully!"})
}
*/
