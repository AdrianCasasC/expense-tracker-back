package repositories

import (
	"context"
	"errors"
	"github.com/AdrianCasasC/expense-tracker-back/database"
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetExpensesDB() ([]models.ExpenseEntity, error) {
	collection := database.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.ExpenseEntity{}, err
	}
	defer cursor.Close(ctx)

	var expenses []models.ExpenseEntity
	if err = cursor.All(ctx, &expenses); err != nil {
		return []models.ExpenseEntity{}, err
	}

	return expenses, nil
}

func InsertExpenseDB(expense models.ExpenseEntity) error {
	collection := database.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, expense)
	if err != nil {
		log.Println("Error inserting expense:", err)
		return errors.New("Error inserting expense")
	}

	return nil
}

func ModifyExpenseDB(expenseId string, fields bson.M) (models.ExpenseEntity, error) {
	collection := database.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, errId := primitive.ObjectIDFromHex(expenseId)
	if errId != nil {
		return models.ExpenseEntity{}, errId
	}
	filter := bson.M{"_id": id}

	var updatedEntity models.ExpenseEntity
	err := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", fields}}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedEntity)
	if err != nil {
		return models.ExpenseEntity{}, err
	}

	return updatedEntity, nil
}

func DeleteExpenseDB(expenseId string) (models.ExpenseEntity, error) {
	collection := database.GetCollection("expenses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, errId := primitive.ObjectIDFromHex(expenseId)
	if errId != nil {
		return models.ExpenseEntity{}, errId
	}
	filter := bson.M{"_id": id}

	var deletedEntity models.ExpenseEntity
	err := collection.FindOneAndDelete(ctx, filter).Decode(&deletedEntity)
	if err != nil {
		return models.ExpenseEntity{}, err
	}

	return deletedEntity, nil
}
