package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type AuthorEquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AuthorEquipmentRepository) FindAll() ([]model.Equipment, error) {
	equipments := []model.Equipment{}
	dbResult := repo.DatabaseConnection.Find(&equipments)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return equipments, nil
}
