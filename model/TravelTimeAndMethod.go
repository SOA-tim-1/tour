package model

type TravelMethod int64

const (
	Car Status = iota
	Bicycle
	Walking
)

type TravelTimeAndMethod struct {
	ID           int64        `json:"id" gorm:"primaryKey:autoIncrement"`
	TravelTime   int64        `json:"travelTime"`
	TravelMethod TravelMethod `json:"travelMethod"`
	TourId       int64        `json:"tourId" gorm:"foreignKey"`
	Tour         Tour         `json:"-"`
}
