package repo

import "database-example/model"

type ICheckpointRepository interface {
	FindById(id string) (model.Checkpoint, error)
	FindByTourId(tourId int64) ([]model.Checkpoint, error)
	CreateCheckpoint(checkpoint *model.Checkpoint) (model.Checkpoint, error)
}
