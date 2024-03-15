package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) FindById(id int64) (model.Tour, error) {
	tour := model.Tour{}
	dbResult := repo.DatabaseConnection.Preload("Checkpoints").First(&tour, "id = ?", id)
	if dbResult != nil {
		return tour, dbResult.Error
	}

	return tour, nil
}

func (repo *TourRepository) FindByAuthorId(authorId int64) ([]model.Tour, error) {
	tours := []model.Tour{}
	dbResult := repo.DatabaseConnection.Preload("Checkpoints").Find(&tours, "author_id = ?", authorId)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return tours, nil
}

func (repo *TourRepository) CreateTour(tour *model.Tour) (model.Tour, error) {
	// Attempt to find a tour with the same ID
	existingTour := model.Tour{}
	err := repo.DatabaseConnection.First(&existingTour, "id = ?", tour.ID).Error
	if err == nil {
		// Tour with the same ID found, update it
		dbResult := repo.DatabaseConnection.Model(&existingTour).Updates(tour)
		if dbResult.Error != nil {
			return model.Tour{}, dbResult.Error
		}
		println("Rows affected: ", dbResult.RowsAffected)
		return existingTour, nil
	}

	// Tour with the same ID not found, create a new one
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return model.Tour{}, dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return *tour, nil
}

func (repo *TourRepository) PublishTour(id int64) error {
	tour := model.Tour{}
	err := repo.DatabaseConnection.First(&tour, "id = ?", id).Error
	if err != nil {
		return err
	}

	// Update the status of the tour to indicate it's published
	tour.Status = 1

	// Save the updated tour back to the database
	err = repo.DatabaseConnection.Save(&tour).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *TourRepository) ArchiveTour(id int64) error {
	tour := model.Tour{}
	err := repo.DatabaseConnection.First(&tour, "id = ?", id).Error
	if err != nil {
		return err
	}

	// Update the status of the tour to indicate it's published
	tour.Status = 2

	// Save the updated tour back to the database
	err = repo.DatabaseConnection.Save(&tour).Error
	if err != nil {
		return err
	}

	return nil
}
