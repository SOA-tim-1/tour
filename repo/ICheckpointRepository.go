package repo

import "database-example/model"

type ICheckpointRepository interface {
	FindById(id int64) (model.Checkpoint, error)
	FindByTourId(tourId int64) ([]model.Checkpoint, error)
	CreateCheckpoint(checkpoint *model.Checkpoint) (model.Checkpoint, error)
	DeleteById(id int64) error
}
