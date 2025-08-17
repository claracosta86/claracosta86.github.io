package repository

import (
	"errors"

	"poc2/back/model"

)

type EventRepository interface {
	SaveEventData(event model.Event) error
	FetchUserFavoritesByID(userID int) ([]model.Event, error)
	FetchAllEvents() ([]model.Event, error)
	FetchEventByID(eventID int) (model.Event, error)
	UpdateEventData(event model.Event) error
	DeleteEventByID(eventID int) error
}

type eventsRepository struct{
	userRepository UserRepository
}

func NewEventRepository() EventRepository {
	return &eventsRepository{}
}

func (r *eventsRepository) SaveEventData(event model.Event) error {
	return errors.New("not implemented")
}

func (r *eventsRepository) FetchUserFavoritesByID(userID int) ([]model.Event, error) {
	return nil, errors.New("not implemented")
}

func (r *eventsRepository) FetchAllEvents() ([]model.Event, error) {
	return nil, errors.New("not implemented")
}

func (r *eventsRepository) FetchEventByID(eventID int) (model.Event, error) {
	return model.Event{}, errors.New("not implemented")
}

func (r *eventsRepository) UpdateEventData(event model.Event) error {
	return errors.New("not implemented")
}

func (r *eventsRepository) DeleteEventByID(eventID int) error {
	return errors.New("not implemented")
}