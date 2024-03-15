package model

type Equipment struct {
	ID          int64 `gorm:"primaryKey:autoIncrement"`
	Name        string
	Description string
	Tours       []Tour `gorm:"many2many:tourequipment;"`
}
