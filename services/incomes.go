package services

import (
	"github.com/AdrianCasasC/expense-tracker-back/models"
	"github.com/AdrianCasasC/expense-tracker-back/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllIncomes() ([]models.IncomeDto, error) {
	incomesEntities, err := repositories.GetIncomesDB()
	if err != nil {
		return []models.IncomeDto{}, err
	}

	incomesDtos := incomeFromEntitiesToDto(incomesEntities)
	return incomesDtos, nil
}

func CreateIncome(income models.IncomeDto) error {
	incomeEntity := incomeFromDtoToEntity(income)
	if err := repositories.InsertIncomeDB(incomeEntity); err != nil {
		return err
	}

	return nil
}

func UpdateIncome(incomeId string, income models.IncomeDto, fieldsMap map[string]bool) (models.IncomeDto, error) {
	incomeEntity := incomeFromDtoToEntity(income)

	fields := clearIncomePatchObject(incomeEntity, fieldsMap)

	updatedEntity, err := repositories.ModifyIncomeDB(incomeId, fields)
	if err != nil {
		return models.IncomeDto{}, err
	}

	return incomeFromEntityToDto(updatedEntity), nil
}

func DeleteIncome(incomeId string) (models.IncomeDto, error) {
	deletedEntity, err := repositories.DeleteIncomeDB(incomeId)
	if err != nil {
		return models.IncomeDto{}, err
	}

	return incomeFromEntityToDto(deletedEntity), nil
}

func clearIncomePatchObject(entity models.IncomeEntity, fields map[string]bool) bson.M {

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

func incomeFromEntitiesToDto(entities []models.IncomeEntity) []models.IncomeDto {
	var incomeDtos []models.IncomeDto
	for _, entity := range entities {
		incomeDtos = append(incomeDtos, incomeFromEntityToDto(entity))
	}

	return incomeDtos
}

func incomeFromEntityToDto(entity models.IncomeEntity) models.IncomeDto {
	var incomeDto models.IncomeDto
	incomeDto.ID = entity.ID
	incomeDto.Name = entity.Name
	incomeDto.Value = entity.Value
	incomeDto.Type = "income"
	incomeDto.Category = entity.Category
	incomeDto.Date = entity.Date
	return incomeDto
}

func incomeFromDtoToEntity(dto models.IncomeDto) models.IncomeEntity {
	var incomeEntity models.IncomeEntity
	incomeEntity.ID = dto.ID
	incomeEntity.Name = dto.Name
	incomeEntity.Value = dto.Value
	incomeEntity.Category = dto.Category
	incomeEntity.Date = dto.Date
	return incomeEntity
}
