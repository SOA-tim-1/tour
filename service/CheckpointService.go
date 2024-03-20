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

func (service *CheckpointService) FindCheckpoint(id int64) (*dtos.CheckpointDto, error) {
	checkpoint, err := service.CheckpointRepo.FindById(id)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %d not found", id))
	}

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

	checkpoint := model.Checkpoint{}
	automapper.Map(checkpointDto, &checkpoint)

	checkpoint.Coordinate = model.Coordinate{
		Latitude:  checkpointDto.Latitude,
		Longitude: checkpointDto.Longitude,
	}

	checkpoint, err := service.CheckpointRepo.CreateCheckpoint(&checkpoint)
	if err != nil {
		return nil, err
	}

	checkpointDtoReturn := dtos.CheckpointDto{}

	automapper.Map(checkpoint, &checkpointDtoReturn)
	checkpointDtoReturn.Latitude = checkpoint.Coordinate.Latitude
	checkpointDtoReturn.Longitude = checkpoint.Coordinate.Longitude

	return &checkpointDtoReturn, nil
}

func (service *CheckpointService) Delete(id int64) error {
	err := service.CheckpointRepo.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func (service *CheckpointService) CheckIfPointsAreValidForPublish(tourId int64) (bool, error) {
	publishable, err := service.CheckpointRepo.CheckIfPointsAreValidForPublish(tourId)
	if err != nil {
		return false, err
	}

	return publishable, nil
}
