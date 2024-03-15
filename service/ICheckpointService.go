package service

import (
	"database-example/dtos"
)

type ICheckpointService interface {
	FindCheckpoint(id string) (*dtos.CheckpointDto, error)
	FindByTourId(tourId int64) ([]*dtos.CheckpointDto, error)
	Create(checkpointDto *dtos.CheckpointDto) (*dtos.CheckpointDto, error)
}
