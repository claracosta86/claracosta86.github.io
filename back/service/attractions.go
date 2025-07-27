package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type AttractionService interface {
	RegisterAttraction(attraction model.TouristAttraction) error
	GetAllAttractions() ([]model.TouristAttraction, error)
	GetAttractionByID(attractionID int) (model.TouristAttraction, error)
	UpdateAttraction(attraction model.TouristAttraction) error
	DeleteAttractionByID(attractionID int) error
}

type attractionService struct {
	attractionRepository repository.AttractionRepository
}

func NewAttractionService() *attractionService {
	return &attractionService{
		attractionRepository: repository.NewAttractionRepository(),
	}
}

func (s *attractionService) RegisterAttraction(attraction model.TouristAttraction) error {
	return s.attractionRepository.SaveAttractionData(attraction)
}

func (s *attractionService) GetAllAttractions() ([]model.TouristAttraction, error) {
	return s.attractionRepository.FetchAllAttractions()
}

func (s *attractionService) GetAttractionByID(attractionID int) (model.TouristAttraction, error) {
	return s.attractionRepository.FetchAttractionByID(attractionID)
}

func (s *attractionService) UpdateAttraction(attraction model.TouristAttraction) error {
	return s.attractionRepository.UpdateAttractionData(attraction)
}

func (s *attractionService) DeleteAttractionByID(attractionID int) error {
	return s.attractionRepository.DeleteAttractionByID(attractionID)
}

