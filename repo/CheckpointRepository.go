package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type CheckpointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CheckpointRepository) FindById(id string) (model.Checkpoint, error) {
	checkpoint := model.Checkpoint{}
	dbResult := repo.DatabaseConnection.First(&checkpoint, "id = ?", id)
	if dbResult != nil {
		return checkpoint, dbResult.Error
	}

	return checkpoint, nil
}

func (repo *CheckpointRepository) FindByTourId(tourId int64) ([]model.Checkpoint, error) {
	checkpoints := []model.Checkpoint{}
	dbResult := repo.DatabaseConnection.Find(&checkpoints, "tour_id = ?", tourId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return checkpoints, nil
}

func (repo *CheckpointRepository) CreateCheckpoint(checkpoint *model.Checkpoint) (model.Checkpoint, error) {
	dbResult := repo.DatabaseConnection.Create(checkpoint)
	if dbResult.Error != nil {
		return model.Checkpoint{}, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return *checkpoint, nil
}
