package dtos

type TravelMethod int64

const (
	Car Status = iota
	Bicycle
	Walking
)

type TravelTimeAndMethodDto struct {
	TravelTime   int64
	TravelMethod TravelMethod
}
