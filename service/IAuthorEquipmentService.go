package service

import "database-example/dtos"

type IAuthorEquipmentService interface {
	FindAll() ([]*dtos.EquipmentDto, error)
}
