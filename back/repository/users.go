package repository

import (
	"errors"

	"poc2/back/model"

)

type UserRepository interface {
	SaveUserData(user model.User) error
	FetchUserDataByID(userID int) (model.User, error)
	FetchUserFavorites(userID string) (*model.UserFavorites, error)
	UpdateUserData(user model.User) error
	AddEventToFavorites(user model.User) error
	AddAttractionToFavorites(user model.User) error
	DeleteUserByID(userID string) error
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

func (r *userRepository) FetchUserFavorites(userID string) (*model.UserFavorites, error) {
	return nil, errors.New("not implemented")
}	

func (r *userRepository) UpdateUserData(user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddEventToFavorites(user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) AddAttractionToFavorites(user model.User) error {
	return errors.New("not implemented")
}

func (r *userRepository) DeleteUserByID(userID string) error {
	return errors.New("not implemented")
}