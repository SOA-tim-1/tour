package repo

import "database-example/model"

type ITourRepository interface {
	FindById(id int64) (model.Tour, error)
	FindByAuthorId(authorId int64) ([]model.Tour, error)
	CreateTour(tour *model.Tour) (model.Tour, error)
	PublishTour(id int64) error
	ArchiveTour(id int64) error
}
