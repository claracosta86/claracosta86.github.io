package service

import (
	"context"

	"poc2/back/model"
	"poc2/back/repository"
	"poc2/back/lib/errors"

)

type UserService interface {
	RegisterUser(ctx context.Context, user model.User) error
	GetUserDataByID(ctx context.Context, userID int) (model.User, error)
	GetUserIDByEmail(ctx context.Context, email string) (int, error)
	VerifyUserPassword(ctx context.Context, userID int, password string) (bool, error)
	GetUserFavoritesByID(ctx context.Context, userID int) (*model.UserFavorites, error)
	UpdateUserProfile(ctx context.Context, user model.User) error
	UpdateUserPassword(ctx context.Context, userID int, passwordUpdate model.PasswordUpdate) error
	AddToFavorites(ctx context.Context, userID int, favorite model.UserFavorite) error
	DeleteFromFavorites(ctx context.Context, userID int, favoriteID int) error
	DeleteUserByID(ctx context.Context, userID int) error
}

type userService struct {
	userRepository repository.UserRepository
	eventRepository repository.EventRepository
	attractionRepository repository.AttractionRepository
}

func NewUserService(userRepository repository.UserRepository, eventRepository repository.EventRepository, attractionRepository repository.AttractionRepository) *userService {
	return &userService{
		userRepository:    userRepository,
		eventRepository:   eventRepository,
		attractionRepository: attractionRepository,
	}
}

func (s *userService) RegisterUser(ctx context.Context, user model.User) error {
	return s.userRepository.SaveUserData(ctx, user)
}

func (s *userService) GetUserDataByID(ctx context.Context, userID int) (model.User, error) {
	return s.userRepository.FetchUserDataByID(ctx, userID)
}

func (s *userService) GetUserIDByEmail(ctx context.Context, email string) (int, error) {
	return s.userRepository.FetchUserIDByEmail(ctx, email)
}

func (s *userService) VerifyUserPassword(ctx context.Context, userID int, password string) (bool, error) {
	return s.userRepository.CheckUserPassword(ctx, userID, password)
}

func (s *userService) GetUserFavoritesByID(ctx context.Context, userID int) (*model.UserFavorites, error) {
	events, err := s.eventRepository.FetchUserFavoritesByID(userID)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		events = []model.Event{}
	}

	attractions, err := s.attractionRepository.FetchUserFavoritesByID(userID)
	if err != nil {
		return nil, err
	}
	if len(attractions) == 0 {
		attractions = []model.TouristAttraction{}
	}

	return &model.UserFavorites{
		Events:             events,
		TouristAttractions: attractions,
	}, nil
}

func (s *userService) UpdateUserProfile(ctx context.Context, user model.User) error {
	return s.userRepository.UpdateUserProfile(ctx, user)
}

func (s *userService) UpdateUserPassword(ctx context.Context, userID int, passwordUpdate model.PasswordUpdate) error {
	verifyPassword, err := s.userRepository.CheckUserPassword(ctx, userID, passwordUpdate.CurrentPassword)
	if err != nil {
		return err
	}

    if !verifyPassword {
        return errors.ErrIncorrectPassword
    }

	return s.userRepository.UpdateUserPassword(ctx, userID, passwordUpdate.NewPassword)
}

func (s *userService) AddToFavorites(ctx context.Context, userID int, favorite model.UserFavorite) error {
	return s.userRepository.AddToFavorites(ctx, userID, favorite)
}

func (s *userService) DeleteFromFavorites(ctx context.Context, userID, favoriteID int) error {
	return s.userRepository.RemoveFromFavorites(ctx, userID, favoriteID)
}

func (s *userService) DeleteUserByID(ctx context.Context, userID int) error {
	return s.userRepository.DeleteUserByID(ctx, userID)
}