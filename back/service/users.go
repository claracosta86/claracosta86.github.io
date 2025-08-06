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
	AddEventToFavorites(ctx context.Context, userID int, event model.Event) error
	AddAttractionToFavorites(ctx context.Context, userID int, attraction model.TouristAttraction) error
	DeleteEventFromFavorites(ctx context.Context, userID int, eventID int) error
	DeleteAttractionFromFavorites(ctx context.Context, userID int, attractionID int) error
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

func (s *userService) AddEventToFavorites(ctx context.Context, userID int, event model.Event) error {
	return s.userRepository.AddEventToFavorites(ctx, userID, event)
}

func (s *userService) AddAttractionToFavorites(ctx context.Context, userID int, attraction model.TouristAttraction) error {
	return s.userRepository.AddAttractionToFavorites(ctx, userID, attraction)
}

func (s *userService) DeleteEventFromFavorites(ctx context.Context, userID, eventID int) error {
	return s.userRepository.RemoveEventFromFavorites(ctx, userID, eventID)
}

func (s *userService) DeleteAttractionFromFavorites(ctx context.Context, userID, attractionID int) error {
	return s.userRepository.RemoveAttractionFromFavorites(ctx, userID, attractionID)
}

func (s *userService) DeleteUserByID(ctx context.Context, userID int) error {
	return s.userRepository.DeleteUserByID(ctx, userID)
}