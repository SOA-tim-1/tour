package repo

import "database-example/model"

type IAuthorEquipmentRepository interface {
	FindAll() ([]model.Equipment, error)
}
