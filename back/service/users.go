package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type UserService interface {
	RegisterUser(user model.User) error
	FetchUserDataByID(userID int) (model.User, error)
	FetchUserFavorites(userID string) (*model.UserFavorites, error)
	UpdateUserData(user model.User) error
	AddEventToFavorites(user model.User) error
	AddAttractionToFavorites(user model.User) error
	DeleteUserByID(userID string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser(user model.User) error {
	return s.userRepository.SaveUserData(user)
}

func (s *userService) FetchUserDataByID(userID int) (model.User, error) {
	return s.userRepository.FetchUserDataByID(userID)
}

func (s *userService) FetchUserFavorites(userID string) (*model.UserFavorites, error) {
	return s.userRepository.FetchUserFavorites(userID)
}

func (s *userService) UpdateUserData(user model.User) error {
	return s.userRepository.UpdateUserData(user)
}

func (s *userService) AddEventToFavorites(user model.User) error {
	return s.userRepository.AddEventToFavorites(user)
}

func (s *userService) AddAttractionToFavorites(user model.User) error {
	return s.userRepository.AddAttractionToFavorites(user)
}

func (s *userService) DeleteUserByID(userID string) error {
	return s.userRepository.DeleteUserByID(userID)
}