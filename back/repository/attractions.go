package repository

import (
	"errors"

	"poc2/back/model"

)

type AttractionRepository interface {
	SaveAttractionData(attraction model.TouristAttraction) error
	FetchUserAttractions(userID string) ([]model.TouristAttraction, error)
	FetchAllAttractions() ([]model.TouristAttraction, error)
	FetchAttractionByID(attractionID string) (model.TouristAttraction, error)
	UpdateAttractionData(attraction model.TouristAttraction) error
	DeleteAttractionByID(attractionID string) error
}

type attractionRepository struct{
	userRepository UserRepository
}

func NewAttractionRepository() AttractionRepository {
	return &attractionRepository{
		userRepository: NewUserRepository(),
	}
}

func (r *attractionRepository) SaveAttractionData(attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *attractionRepository) FetchUserAttractions(userID string) ([]model.TouristAttraction, error) {
	return nil, errors.New("not implemented")
}

func (r *attractionRepository) FetchAllAttractions() ([]model.TouristAttraction, error) {
	return nil, errors.New("not implemented")
}

func (r *attractionRepository) FetchAttractionByID(attractionID string) (model.TouristAttraction, error) {
	return model.TouristAttraction{}, errors.New("not implemented")
}

func (r *attractionRepository) UpdateAttractionData(attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *attractionRepository) DeleteAttractionByID(attractionID string) error {
	return errors.New("not implemented")
}