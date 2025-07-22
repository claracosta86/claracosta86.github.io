package service

import (
	"errors"

	"poc2/back/model"
	"poc2/back/lib"
	"poc2/back/repository"

)

type UserService interface {
	RegisterUser(user model.User) error
	FetchUserFavorites(userID int) (model.UserFavorites, error)
}

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(user model.User) error {
	return s.userRepository.SaveUserData(user)
}

func (s *UserService) FetchUserFavorites(userID int) (model.UserFavorites, error) {
	return s.userRepository.FetchUserFavorites(userID)
}

func (s *UserService) AddEventToFavorites(user model.User) error {
	return s.userRepository.AddEventToFavorites(user)
}

func (s *UserService) AddAttractionToFavorites(user model.User) error {
	return s.userRepository.AddAttractionToFavorites(user)
}