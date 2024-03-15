package model

type Checkpoint struct {
	ID          int64      `json:"id" gorm:"primaryKey:autoincrement"`
	Name        string     `json:"name" gorm:"type:string"`
	Description string     `json:"description" gorm:"type:string"`
	PictureURL  string     `json:"pictureURL" gorm:"type:string"`
	Coordinate  Coordinate `gorm:"embedded"`
	TourId      int64      `json:"tourId" gorm:"foreignKey"`
	Tour        Tour       `json:"-"`
}
