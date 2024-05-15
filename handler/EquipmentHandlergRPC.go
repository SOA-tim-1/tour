package handler

import (
	"context"
	"database-example/dtos"
	"database-example/proto/equipment"
	"database-example/service"
)

type EquipmentHandlergRPC struct {
	AuthorEquipmentService service.IAuthorEquipmentService
	equipment.UnimplementedEquipmentServiceServer
}

func (handler *EquipmentHandlergRPC) FindAllEquipments(ctx context.Context, in *equipment.FindAllRequest) (*equipment.FindAllResponse, error) {

	equipments, err := handler.AuthorEquipmentService.FindAll()
	if err != nil {
		return nil, err
	}

	if len(equipments) == 0 {
		return &equipment.FindAllResponse{
			Equipments: []*equipment.EquipmentDto{},
		}, nil
	}

	var equipmentResponses []*equipment.EquipmentDto
	for _, equipment := range equipments {
		equipmentResponse := ConvertEquipmentDtoToEquipmentResponse(equipment)
		equipmentResponses = append(equipmentResponses, equipmentResponse)
	}

	return &equipment.FindAllResponse{
		Equipments: equipmentResponses,
	}, nil
}

func ConvertEquipmentDtoToEquipmentResponse(equipmentDto *dtos.EquipmentDto) *equipment.EquipmentDto {

	equipmentResponse := &equipment.EquipmentDto{
		Id:          equipmentDto.ID,
		Name:        equipmentDto.Name,
		Description: equipmentDto.Description,
	}

	return equipmentResponse
}

func ConvertEquipmentResponseToEquipmentDto(equipmentResponse *equipment.EquipmentDto) *dtos.EquipmentDto {

	equipmentDto := &dtos.EquipmentDto{
		ID:          equipmentResponse.Id,
		Name:        equipmentResponse.Name,
		Description: equipmentResponse.Description,
	}

	return equipmentDto
}
