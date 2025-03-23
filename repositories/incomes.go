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

func GetIncomesDB() ([]models.IncomeEntity, error) {
	collection := database.GetCollection("incomes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []models.IncomeEntity{}, err
	}
	defer cursor.Close(ctx)

	var incomes []models.IncomeEntity
	if err = cursor.All(ctx, &incomes); err != nil {
		return []models.IncomeEntity{}, err
	}

	return incomes, nil
}

func InsertIncomeDB(income models.IncomeEntity) error {
	collection := database.GetCollection("incomes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, income)
	if err != nil {
		log.Println("Error inserting income:", err)
		return errors.New("Error inserting income")
	}

	return nil
}

func ModifyIncomeDB(incomeId string, fields bson.M) (models.IncomeEntity, error) {
	collection := database.GetCollection("incomes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, errId := primitive.ObjectIDFromHex(incomeId)
	if errId != nil {
		return models.IncomeEntity{}, errId
	}
	filter := bson.M{"_id": id}

	var updatedEntity models.IncomeEntity
	err := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", fields}}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedEntity)
	if err != nil {
		return models.IncomeEntity{}, err
	}

	return updatedEntity, nil
}

func DeleteIncomeDB(incomeId string) (models.IncomeEntity, error) {
	collection := database.GetCollection("incomes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, errId := primitive.ObjectIDFromHex(incomeId)
	if errId != nil {
		return models.IncomeEntity{}, errId
	}
	filter := bson.M{"_id": id}

	var deletedEntity models.IncomeEntity
	err := collection.FindOneAndDelete(ctx, filter).Decode(&deletedEntity)
	if err != nil {
		return models.IncomeEntity{}, err
	}

	return deletedEntity, nil
}
