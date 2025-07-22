package repository

import (
	"errors"

	"poc2/back/model"

)

type UserRepository interface {
	SaveUserData(user model.User) error
	FetchUserFavorites(userID int) (model.UserFavorites, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) SaveUserData(user model.User) error {
	return errors.New("not implemented")
}

func (r *UserRepository) FetchUserFavorites(userID int) (model.UserFavorites, error) {
	return nil, errors.New("not implemented")
}	