package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type AttractionService interface {
	RegisterAttraction(attraction model.Attraction) error
	GetUserAttractions(userID string) ([]model.Attraction, error)
	GetAllAttractions() ([]model.Attraction, error)
	GetAttractionByID(attractionID string) (model.Attraction, error)
	UpdateAttraction(attraction model.Attraction) error
}

type attractionService struct {
	attractionRepository repository.AttractionRepository
}

func NewAttractionService() *attractionService {
	return &attractionService{
		attractionRepository: repository.NewAttractionRepository(),
	}
}

func (s *attractionService) RegisterAttraction(attraction model.Attraction) error {
	return s.attractionRepository.SaveAttractionData(attraction)
}

func (s *attractionService) GetUserAttractions(userID string) ([]model.Attraction, error) {
	return s.attractionRepository.FetchUserAttractions(userID)
}

func (s *attractionService) GetAllAttractions() ([]model.Attraction, error) {
	return s.attractionRepository.FetchAllAttractions()
}

func (s *attractionService) GetAttractionByID(attractionID string) (model.Attraction, error) {
	return s.attractionRepository.FetchAttractionByID(attractionID)
}

func (s *attractionService) UpdateAttraction(attraction model.Attraction) error {
	return s.attractionRepository.UpdateAttractionData(attraction)
}

