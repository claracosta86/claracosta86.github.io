package repository

import (
	"errors"

	"poc2/back/model"

)

type UserRepository interface {
	SaveUserData(user model.User) error
	FetchUserDataByID(userID int) (model.User, error)
	UpdateUserData(user model.User) error
	AddEventToFavorites(userID int, event model.Event) error
	AddAttractionToFavorites(userID int, attraction model.TouristAttraction) error
	RemoveEventFromFavorites(userID int, eventID int) error
	RemoveAttractionFromFavorites(userID int, attractionID int) error
	DeleteUserByID(userID int) error
}

type userRepository struct {}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) SaveUserData(user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) FetchUserDataByID(userID int) (model.User, error) {
	return model.User{}, errors.New("not implemented")
}

func (r *userRepository) UpdateUserData(user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddEventToFavorites(userID int, event model.Event) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddAttractionToFavorites(userID int, attraction model.TouristAttraction) error {
	return errors.New("not implemented")
}

func (r *userRepository) RemoveEventFromFavorites(userID, eventID int) error {
	return errors.New("not implemented")
}

func (r *userRepository) RemoveAttractionFromFavorites(userID, attractionID int) error {
	return errors.New("not implemented")
}

func (r *userRepository) DeleteUserByID(userID int) error {
	return errors.New("not implemented")
}