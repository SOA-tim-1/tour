package service

import (
	"database-example/dtos"
	"database-example/repo"
	"fmt"

	"github.com/rafiulgits/go-automapper"
)

type AuthorEquipmentService struct {
	AuthorEquipmentRepo repo.IAuthorEquipmentRepository
}

func (service *AuthorEquipmentService) FindAll() ([]*dtos.EquipmentDto, error) {
	equipments, err := service.AuthorEquipmentRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find equipments: %w", err)
	}

	var equipmentDtos []*dtos.EquipmentDto
	for _, equipment := range equipments {
		var equipmentDto dtos.EquipmentDto
		automapper.Map(&equipment, &equipmentDto)
		equipmentDtos = append(equipmentDtos, &equipmentDto)
	}

	return equipmentDtos, nil
}
