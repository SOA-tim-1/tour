package model

import (
	"time"
)

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

type Tour struct {
	ID          int64        `json:"id" gorm:"primaryKey:autoIncrement"`
	AuthorId    int64        `json:"authorId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Difficult   Difficult    `json:"difficult"`
	Status      Status       `json:"status"`
	Price       float64      `json:"price"`
	Tags        string       `json:"tags"`
	Distance    float64      `json:"distance"`
	Checkpoints []Checkpoint `json:"checkpoints"`
	PublishTime *time.Time   `json:"publishTime"`
	ArchiveTime *time.Time   `json:"archiveTime"`
	Equipments  []Equipment  `json:"tourEquipment" gorm:"many2many:tourequipment;"`
}
