package handler

import (
	"context"
	"database-example/dtos"
	"database-example/proto/checkpoint"
	"database-example/service"
)

type CheckpointHandlergRPC struct {
	CheckpointService service.ICheckpointService
	checkpoint.UnimplementedCheckpointServiceServer
}

func (handler *CheckpointHandlergRPC) FindCheckpoint(ctx context.Context, in *checkpoint.FindCheckpointRequest) (*checkpoint.CheckpointDto, error) {

	checkpoint, err := handler.CheckpointService.FindCheckpoint(in.GetId())
	if err != nil {
		return nil, err
	}

	return ConvertCheckpointDtoToCheckpointResponse(checkpoint), nil
}

func (handler *CheckpointHandlergRPC) FindCheckpointByTourId(ctx context.Context, in *checkpoint.FindByTourIdRequest) (*checkpoint.FindByTourIdResponse, error) {

	checkpoints, err := handler.CheckpointService.FindByTourId(in.GetTourId())
	if err != nil {
		return nil, err
	}

	if len(checkpoints) == 0 {
		return &checkpoint.FindByTourIdResponse{
			Checkpoints: []*checkpoint.CheckpointDto{},
		}, nil
	}

	var checkpointResponses []*checkpoint.CheckpointDto
	for _, checkpoint := range checkpoints {
		checkpointResponse := ConvertCheckpointDtoToCheckpointResponse(checkpoint)
		checkpointResponses = append(checkpointResponses, checkpointResponse)
	}

	return &checkpoint.FindByTourIdResponse{
		Checkpoints: checkpointResponses,
	}, nil
}

func (handler *CheckpointHandlergRPC) CreateCheckpoint(ctx context.Context, in *checkpoint.CheckpointDto) (*checkpoint.CheckpointDto, error) {

	checkpointDto := ConvertCheckpointResponseToCheckpointDto(in)
	createdCheckpointDto, err := handler.CheckpointService.Create(checkpointDto)
	if err != nil {
		return nil, err
	}

	return ConvertCheckpointDtoToCheckpointResponse(createdCheckpointDto), nil
}

func (handler *CheckpointHandlergRPC) DeleteCheckpoint(ctx context.Context, in *checkpoint.DeleteRequest) (*checkpoint.DeleteResponse, error) {

	err := handler.CheckpointService.Delete(in.GetId())
	if err != nil {
		return nil, err
	}

	return &checkpoint.DeleteResponse{}, nil
}

func ConvertCheckpointDtoToCheckpointResponse(checkpointDto *dtos.CheckpointDto) *checkpoint.CheckpointDto {

	checkpointResponse := &checkpoint.CheckpointDto{
		Id:          checkpointDto.ID,
		Name:        checkpointDto.Name,
		Description: checkpointDto.Description,
		PictureUrl:  checkpointDto.PictureURL,
		Latitude:    checkpointDto.Latitude,
		Longitude:   checkpointDto.Longitude,
		TourId:      checkpointDto.TourId,
	}

	return checkpointResponse
}

func ConvertCheckpointResponseToCheckpointDto(checkpointResponse *checkpoint.CheckpointDto) *dtos.CheckpointDto {

	checkpointDto := &dtos.CheckpointDto{
		ID:          checkpointResponse.Id,
		Name:        checkpointResponse.Name,
		Description: checkpointResponse.Description,
		PictureURL:  checkpointResponse.PictureUrl,
		Latitude:    checkpointResponse.Latitude,
		Longitude:   checkpointResponse.Longitude,
		TourId:      checkpointResponse.TourId,
	}

	return checkpointDto
}
