package service

import (
	"database-example/dtos"
)

type ITourService interface {
	FindTour(id int64) (*dtos.TourDto, error)
	FindByAuthorId(authorId int64) ([]*dtos.TourDto, error)
	Create(tourDto *dtos.TourDto) (*dtos.TourDto, error)
	PublishTour(tourId int64) error
	ArchiveTour(tourId int64) error
}
