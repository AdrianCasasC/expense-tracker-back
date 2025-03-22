package controllers

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

var incomesCollection = database.GetCollection("incomes")

func GetIncomes(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := incomesCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch incomes"})
		return
	}
	defer cursor.Close(ctx)

	var incomes []models.IncomeDto
	if err = cursor.All(ctx, &incomes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding incomes"})
		return
	}

	c.JSON(http.StatusOK, incomes)
}

// TODO: Hacer la transformación de Dto a entity para meterlo en base de datos
func CreateIncome(c *gin.Context) {
	var income models.IncomeDto
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID and Date
	income.ID = primitive.NewObjectID()
	income.Date = time.Now().Format(time.RFC3339)

	// Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := incomesCollection.InsertOne(ctx, income)
	if err != nil {
		log.Println("Error inserting income:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert income"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Income added successfully!"})
}
*/
