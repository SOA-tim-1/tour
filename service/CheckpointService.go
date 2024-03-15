package service

import (
	"database-example/dtos"
	"database-example/model"
	"database-example/repo"
	"fmt"

	"github.com/rafiulgits/go-automapper"
)

type CheckpointService struct {
	CheckpointRepo repo.ICheckpointRepository
}

func (service *CheckpointService) FindCheckpoint(id string) (*dtos.CheckpointDto, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}

	// checkpointDto := dtos.CheckpointDto{
	// 	ID:          checkpoint.ID,
	// 	Name:        checkpoint.Name,
	// 	Description: checkpoint.Description,
	// 	PictureURL:  checkpoint.PictureURL,
	// 	Latitude:    checkpoint.Coordinate.Latitude,
	// 	Longitude:   checkpoint.Coordinate.Longitude,
	// 	TourId:      checkpoint.TourId,
	// }
	checkpointDto := dtos.CheckpointDto{}
	automapper.Map(checkpoint, &checkpointDto, func(src *model.Checkpoint, dst *dtos.CheckpointDto) {
		dst.Latitude = src.Coordinate.Latitude
		dst.Longitude = src.Coordinate.Longitude
	})

	return &checkpointDto, nil
}

func (service *CheckpointService) FindByTourId(tourId int64) ([]*dtos.CheckpointDto, error) {
	checkpoints, err := service.CheckpointRepo.FindByTourId(tourId)
	if err != nil {
		return nil, fmt.Errorf("failed to find tours for author with ID %d: %w", tourId, err)
	}

	var checkpointDtos []*dtos.CheckpointDto
	// for _, checkpoint := range checkpoints {
	// 	checkpointDto := dtos.CheckpointDto{
	// 		ID:          checkpoint.ID,
	// 		Name:        checkpoint.Name,
	// 		Description: checkpoint.Description,
	// 		PictureURL:  checkpoint.PictureURL,
	// 		Latitude:    checkpoint.Coordinate.Latitude,
	// 		Longitude:   checkpoint.Coordinate.Longitude,
	// 		TourId:      checkpoint.TourId,
	// 	}
	// 	checkpointDtos = append(checkpointDtos, &checkpointDto)
	// }
	for _, checkpoint := range checkpoints {
		var checkpointDto dtos.CheckpointDto
		automapper.Map(&checkpoint, &checkpointDto, func(src *model.Checkpoint, dst *dtos.CheckpointDto) {
			dst.Latitude = src.Coordinate.Latitude
			dst.Longitude = src.Coordinate.Longitude
		})
		checkpointDtos = append(checkpointDtos, &checkpointDto)
	}

	return checkpointDtos, nil
}

func (service *CheckpointService) Create(checkpointDto *dtos.CheckpointDto) (*dtos.CheckpointDto, error) {

	// checkpoint := model.Checkpoint{
	// 	ID:          checkpointDto.ID,
	// 	Name:        checkpointDto.Name,
	// 	Description: checkpointDto.Description,
	// 	PictureURL:  checkpointDto.PictureURL,
	// 	Coordinate: model.Coordinate{
	// 		Latitude:  checkpointDto.Latitude,
	// 		Longitude: checkpointDto.Longitude,
	// 	},
	// 	TourId: checkpointDto.TourId,
	// }

	var checkpoint model.Checkpoint
	automapper.Map(checkpointDto, &checkpoint, func(src *dtos.CheckpointDto, dst *model.Checkpoint) {
		dst.Coordinate = model.Coordinate{Latitude: src.Latitude, Longitude: src.Longitude}
	})

	checkpoint, err := service.CheckpointRepo.CreateCheckpoint(&checkpoint)
	if err != nil {
		return nil, err
	}

	checkpointDtoReturn := dtos.CheckpointDto{}

	automapper.Map(checkpoint, &checkpointDtoReturn, func(src *model.Checkpoint, dst *dtos.CheckpointDto) {
		dst.Latitude = src.Coordinate.Latitude
		dst.Longitude = src.Coordinate.Longitude
	})

	return &checkpointDtoReturn, nil
}
