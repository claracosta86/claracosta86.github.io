package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type UserService interface {
	RegisterUser(user model.User) error
	GetUserDataByID(userID int) (model.User, error)
	GetUserFavoritesByID(userID int) (*model.UserFavorites, error)
	UpdateUserData(user model.User) error
	AddEventToFavorites(userID int, event model.Event) error
	AddAttractionToFavorites(userID int, attraction model.TouristAttraction) error
	DeleteEventFromFavorites(userID int, eventID int) error
	DeleteAttractionFromFavorites(userID int, attractionID int) error
	DeleteUserByID(userID int) error
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

func (s *userService) RegisterUser(user model.User) error {
	return s.userRepository.SaveUserData(user)
}

func (s *userService) GetUserDataByID(userID int) (model.User, error) {
	return s.userRepository.FetchUserDataByID(userID)
}

func (s *userService) GetUserFavoritesByID(userID int) (*model.UserFavorites, error) {
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

func (s *userService) UpdateUserData(user model.User) error {
	return s.userRepository.UpdateUserData(user)
}

func (s *userService) AddEventToFavorites(userID int, event model.Event) error {
	return s.userRepository.AddEventToFavorites(userID, event)
}

func (s *userService) AddAttractionToFavorites(userID int, attraction model.TouristAttraction) error {
	return s.userRepository.AddAttractionToFavorites(userID, attraction)
}

func (s *userService) DeleteEventFromFavorites(userID, eventID int) error {
	return s.userRepository.RemoveEventFromFavorites(userID, eventID)
}

func (s *userService) DeleteAttractionFromFavorites(userID, attractionID int) error {
	return s.userRepository.RemoveAttractionFromFavorites(userID, attractionID)
}

func (s *userService) DeleteUserByID(userID int) error {
	return s.userRepository.DeleteUserByID(userID)
}