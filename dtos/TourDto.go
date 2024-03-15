package dtos

import "time"

type Difficult int

const (
	Easy Difficult = iota
	Medium
	Hard
)

type Status int64

const (
	Draft Status = iota
	Published
	Archived
)

type TourDto struct {
	ID          int64           `json:"id"`
	AuthorId    int64           `json:"authorId"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Difficult   Difficult       `json:"difficult"`
	Status      Status          `json:"status"`
	Price       float64         `json:"price"`
	Tags        string          `json:"tags"`
	Distance    float64         `json:"distance"`
	Checkpoints []CheckpointDto `json:"checkpoints"`
	PublishTime *time.Time      `json:"publishTime"`
	ArchiveTime *time.Time      `json:"archiveTime"`
	Equipments  []EquipmentDto  `json:"tourEquipment"`
}
