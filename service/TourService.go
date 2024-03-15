package service

import (
	"database-example/dtos"
	"database-example/model"
	"database-example/repo"
	"fmt"

	"github.com/rafiulgits/go-automapper"
)

type TourService struct {
	TourRepo repo.ITourRepository
}

func (service *TourService) FindTour(id int64) (*dtos.TourDto, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %d not found", id))
	}

	tourDto := dtos.TourDto{}
	automapper.Map(tour, &tourDto)

	return &tourDto, nil
}

func (service *TourService) FindByAuthorId(authorId int64) ([]*dtos.TourDto, error) {
	tours, err := service.TourRepo.FindByAuthorId(authorId)
	if err != nil {
		return nil, fmt.Errorf("failed to find tours for author with ID %d: %w", authorId, err)
	}

	var tourDtos []*dtos.TourDto
	for _, tour := range tours {
		var tourDto dtos.TourDto
		automapper.Map(&tour, &tourDto, func(src *model.Tour, dst *dtos.TourDto) {
			dst.Status = dtos.Status(src.Status)
			dst.Difficult = dtos.Difficult(src.Difficult)
		})
		tourDtos = append(tourDtos, &tourDto)
	}

	return tourDtos, nil
}

func (service *TourService) Create(tourDto *dtos.TourDto) (*dtos.TourDto, error) {
	var tour model.Tour
	automapper.Map(tourDto, &tour)

	createdTour, err := service.TourRepo.CreateTour(&tour)
	if err != nil {
		return nil, err
	}

	// Map the created tour back to DTO
	createdTourDto := dtos.TourDto{}
	automapper.Map(&createdTour, &createdTourDto)

	return &createdTourDto, nil
}

func (service *TourService) PublishTour(tourId int64) error {
	err := service.TourRepo.PublishTour(tourId)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) ArchiveTour(tourId int64) error {
	err := service.TourRepo.ArchiveTour(tourId)
	if err != nil {
		return err
	}
	return nil
}
