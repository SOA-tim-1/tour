package dtos

type CheckpointDto struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PictureURL  string  `json:"pictureURL"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	TourId      int64   `json:"tourId"`
}
