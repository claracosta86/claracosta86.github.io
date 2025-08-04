package repository

import (
	"errors"
	"database/sql"

	"poc2/back/model"

)

type AttractionRepository interface {
	SaveAttractionData(attraction model.TouristAttraction) error
	FetchUserFavoritesByID(userID int) ([]model.TouristAttraction, error)
	FetchAllAttractions() ([]model.TouristAttraction, error)
	FetchAttractionByID(attractionID int) (model.TouristAttraction, error)
	UpdateAttractionData(attraction model.TouristAttraction) error
	DeleteAttractionByID(attractionID int) error
}

type attractionRepository struct{
	db *sql.DB
	userRepository UserRepository
}

func NewAttractionRepository(db *sql.DB, ur UserRepository) AttractionRepository {
	return &attractionRepository{
		db: db,
		userRepository: ur,
	}
}

func (r *attractionRepository) SaveAttractionData(attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *attractionRepository) FetchUserFavoritesByID(userID int) ([]model.TouristAttraction, error) {
	return nil, errors.New("not implemented")
}

func (r *attractionRepository) FetchAllAttractions() ([]model.TouristAttraction, error) {
	return nil, errors.New("not implemented")
}

func (r *attractionRepository) FetchAttractionByID(attractionID int) (model.TouristAttraction, error) {
	return model.TouristAttraction{}, errors.New("not implemented")
}

func (r *attractionRepository) UpdateAttractionData(attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *attractionRepository) DeleteAttractionByID(attractionID int) error {
	return errors.New("not implemented")
}