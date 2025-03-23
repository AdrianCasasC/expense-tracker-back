package services

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/AdrianCasasC/expense-tracker-back/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllExpenses() ([]models.ExpenseDto, error) {
	expensesEntities, err := repositories.GetExpensesDB()
	if err != nil {
		return []models.ExpenseDto{}, err
	}

	expensesDtos := expenseFromEntitiesToDto(expensesEntities)
	return expensesDtos, nil
}

func CreateExpense(expense models.ExpenseDto) error {
	expenseEntity := expenseFromDtoToEntity(expense)
	if err := repositories.InsertExpenseDB(expenseEntity); err != nil {
		return err
	}

	return nil
}

func UpdateExpense(expenseId string, expense models.ExpenseDto, fieldsMap map[string]bool) (models.ExpenseDto, error) {
	expenseEntity := expenseFromDtoToEntity(expense)

	fields := clearPatchObject(expenseEntity, fieldsMap)

	updatedEntity, err := repositories.ModifyExpenseDB(expenseId, fields)
	if err != nil {
		return models.ExpenseDto{}, err
	}

	return expenseFromEntityToDto(updatedEntity), nil
}

func DeleteExpense(expenseId string) (models.ExpenseDto, error) {
	deletedEntity, err := repositories.DeleteExpenseDB(expenseId)
	if err != nil {
		return models.ExpenseDto{}, err
	}

	return expenseFromEntityToDto(deletedEntity), nil
}

func clearPatchObject(entity models.ExpenseEntity, fields map[string]bool) bson.M {

	updateEntity := bson.M{}

	if fields["name"] {
		updateEntity["name"] = entity.Name
	}

	if fields["value"] {
		updateEntity["value"] = entity.Value
	}

	if fields["category"] {
		updateEntity["category"] = entity.Category
	}

	if fields["date"] {
		updateEntity["date"] = entity.Date
	}

	return updateEntity
}

func expenseFromEntitiesToDto(entities []models.ExpenseEntity) []models.ExpenseDto {
	var expenseDtos []models.ExpenseDto
	for _, entity := range entities {
		expenseDtos = append(expenseDtos, expenseFromEntityToDto(entity))
	}

	return expenseDtos
}

func expenseFromEntityToDto(entity models.ExpenseEntity) models.ExpenseDto {
	var expenseDto models.ExpenseDto
	expenseDto.ID = entity.ID
	expenseDto.Name = entity.Name
	expenseDto.Value = entity.Value
	expenseDto.Type = "expense"
	expenseDto.Category = entity.Category
	expenseDto.Date = entity.Date
	return expenseDto
}

func expenseFromDtoToEntity(dto models.ExpenseDto) models.ExpenseEntity {
	var expenseEntity models.ExpenseEntity
	expenseEntity.ID = dto.ID
	expenseEntity.Name = dto.Name
	expenseEntity.Value = dto.Value
	expenseEntity.Category = dto.Category
	expenseEntity.Date = dto.Date
	return expenseEntity
}
