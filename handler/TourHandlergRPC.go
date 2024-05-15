package handler

import (
	"context"
	"database-example/dtos"
	"database-example/proto/tour"
	"database-example/service"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

type TourHandlergRPC struct {
	TourService       service.ITourService
	CheckpointService service.ICheckpointService
}

func (handler *TourHandlergRPC) FindTour(ctx context.Context, in *tour.FindTourRequest) (*tour.TourDto, error) {

	tour, err := handler.TourService.FindTour(in.GetId())
	if err != nil {
		return nil, err
	}
	return ConvertTourDtoToTourResponse(tour), nil
}

func (handler *TourHandlergRPC) FindTourByAuthorId(ctx context.Context, in *tour.FindByAuthorIdRequest) (*tour.FindByAuthorIdResponse, error) {

	tours, err := handler.TourService.FindByAuthorId(in.GetAuthorId())
	if err != nil {
		return nil, err
	}

	if len(tours) == 0 {
		return &tour.FindByAuthorIdResponse{
			Tours: []*tour.TourDto{},
		}, nil
	}

	var tourResponses []*tour.TourDto
	for _, tour := range tours {
		tourResponse := ConvertTourDtoToTourResponse(tour)
		tourResponses = append(tourResponses, tourResponse)
	}

	return &tour.FindByAuthorIdResponse{
		Tours: tourResponses,
	}, nil

}

func (handler *TourHandlergRPC) CreateTour(ctx context.Context, in *tour.TourDto) (*tour.TourDto, error) {

	tourDto := ConvertTourResponseToTourDto(in)

	createdTourDto, err := handler.TourService.Create(tourDto)
	if err != nil {
		return nil, err
	}

	return ConvertTourDtoToTourResponse(createdTourDto), nil
}

func (handler *TourHandlergRPC) UpdateTour(ctx context.Context, in *tour.TourDto) (*tour.TourDto, error) {

	tourDto := ConvertTourResponseToTourDto(in)

	updatedTourDto, err := handler.TourService.Update(tourDto)
	if err != nil {
		return nil, err
	}

	return ConvertTourDtoToTourResponse(updatedTourDto), nil
}

func (handler *TourHandlergRPC) PublishTour(ctx context.Context, in *tour.PublishTourRequest, opts ...grpc.CallOption) (*tour.PublishTourResponse, error) {

	publishable, err := handler.CheckpointService.CheckIfPointsAreValidForPublish(in.GetTourId())
	if err != nil {
		return nil, err
	}

	if publishable {
		err = handler.TourService.PublishTour(in.GetTourId())
		if err != nil {
			return nil, err
		}

		return &tour.PublishTourResponse{}, nil
	} else {
		return nil, err
	}
}

func (handler *TourHandlergRPC) ArchiveTour(ctx context.Context, in *tour.ArchiveTourRequest) (*tour.ArchiveTourResponse, error) {

	err := handler.TourService.ArchiveTour(in.GetTourId())

	if err != nil {
		return nil, err
	}

	return &tour.ArchiveTourResponse{}, nil
}

func ConvertTourDtoToTourResponse(tourDto *dtos.TourDto) *tour.TourDto {

	var tourCheckpoints []*tour.CheckpointDto
	for _, dto := range tourDto.Checkpoints {
		checkpoint := &tour.CheckpointDto{
			Id:          dto.ID,
			Name:        dto.Name,
			Description: dto.Description,
			PictureUrl:  dto.PictureURL,
			Latitude:    dto.Latitude,
			Longitude:   dto.Longitude,
			TourId:      dto.TourId,
		}
		tourCheckpoints = append(tourCheckpoints, checkpoint)
	}

	var tourEquipments []*tour.EquipmentDto
	for _, dto := range tourDto.Equipments {
		equipment := &tour.EquipmentDto{
			Id:          dto.ID,
			Name:        dto.Name,
			Description: dto.Description,
		}
		tourEquipments = append(tourEquipments, equipment)
	}

	var tourTravelTimeAndMethod []*tour.TravelTimeAndMethodDto
	for _, dto := range tourDto.TravelTimeAndMethod {
		travelTimeAndMethod := &tour.TravelTimeAndMethodDto{
			TravelTime:   dto.TravelTime,
			TravelMethod: tour.TravelMethod(dto.TravelMethod),
		}
		tourTravelTimeAndMethod = append(tourTravelTimeAndMethod, travelTimeAndMethod)
	}

	publishTime, _ := ptypes.TimestampProto(*tourDto.PublishTime)

	// Similarly, convert tourDto.ArchiveTime to google.protobuf.Timestamp
	archiveTime, _ := ptypes.TimestampProto(*tourDto.ArchiveTime)

	tourResponse := &tour.TourDto{
		Id:                  tourDto.ID,
		AuthorId:            tourDto.AuthorId,
		Name:                tourDto.Name,
		Description:         tourDto.Description,
		Difficult:           tour.Difficult(tourDto.Difficult),
		Status:              tour.Status(tourDto.Status),
		Price:               float32(tourDto.Price),
		Tags:                tourDto.Tags,
		Distance:            float32(tourDto.Distance),
		Checkpoints:         tourCheckpoints,
		PublishTime:         publishTime,
		ArchiveTime:         archiveTime,
		Equipments:          tourEquipments,
		TravelTimeAndMethod: tourTravelTimeAndMethod,
	}

	return tourResponse
}

func ConvertTourResponseToTourDto(tourResponse *tour.TourDto) *dtos.TourDto {
	var dtosCheckpoints []dtos.CheckpointDto
	for _, checkpoint := range tourResponse.Checkpoints {
		dtoCheckpoint := &dtos.CheckpointDto{
			ID:          checkpoint.Id,
			Name:        checkpoint.Name,
			Description: checkpoint.Description,
			PictureURL:  checkpoint.PictureUrl,
			Latitude:    checkpoint.Latitude,
			Longitude:   checkpoint.Longitude,
			TourId:      checkpoint.TourId,
		}
		dtosCheckpoints = append(dtosCheckpoints, *dtoCheckpoint)
	}

	var dtosEquipments []dtos.EquipmentDto
	for _, equipment := range tourResponse.Equipments {
		dtoEquipment := &dtos.EquipmentDto{
			ID:          equipment.Id,
			Name:        equipment.Name,
			Description: equipment.Description,
		}
		dtosEquipments = append(dtosEquipments, *dtoEquipment)
	}

	var dtosTravelTimeAndMethod []dtos.TravelTimeAndMethodDto
	for _, travelTimeAndMethod := range tourResponse.TravelTimeAndMethod {
		dtoTravelTimeAndMethod := &dtos.TravelTimeAndMethodDto{
			TravelTime:   travelTimeAndMethod.TravelTime,
			TravelMethod: dtos.TravelMethod(travelTimeAndMethod.TravelTime),
		}
		dtosTravelTimeAndMethod = append(dtosTravelTimeAndMethod, *dtoTravelTimeAndMethod)
	}

	publishTime, _ := ptypes.Timestamp(tourResponse.PublishTime)

	// Similarly, convert tourDto.ArchiveTime to time.Time
	archiveTime, _ := ptypes.Timestamp(tourResponse.ArchiveTime)

	tourDto := &dtos.TourDto{
		ID:                  tourResponse.Id,
		AuthorId:            tourResponse.AuthorId,
		Name:                tourResponse.Name,
		Description:         tourResponse.Description,
		Difficult:           dtos.Difficult(tourResponse.Difficult),
		Status:              dtos.Status(tourResponse.Status),
		Price:               float64(tourResponse.Price),
		Tags:                tourResponse.Tags,
		Distance:            float64(tourResponse.Distance),
		Checkpoints:         dtosCheckpoints,
		PublishTime:         &publishTime,
		ArchiveTime:         &archiveTime,
		Equipments:          dtosEquipments,
		TravelTimeAndMethod: dtosTravelTimeAndMethod,
	}

	return tourDto
}
