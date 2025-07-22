package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type EventService interface {
	RegisterEvent(event model.Event) error
	GetUserEvents(userID string) ([]model.Event, error)
	GetAllEvents() ([]model.Event, error)
	GetEventByID(eventID string) (model.Event, error)
	UpdateEvent(event model.Event) error
}

type eventService struct {
	eventRepository repository.EventRepository
}

func NewEventService() *eventService {
	return &eventService{
		eventRepository: repository.NewEventRepository(),
	}
}

func (s *eventService) RegisterEvent(event model.Event) error {
	return s.eventRepository.SaveEventData(event)
}

func (s *eventService) GetUserEvents(userID string) ([]model.Event, error) {
	return s.eventRepository.FetchUserEvents(userID)
}

func (s *eventService) GetAllEvents() ([]model.Event, error) {
	return s.eventRepository.FetchAllEvents()
}

func (s *eventService) GetEventByID(eventID string) (model.Event, error) {
	return s.eventRepository.FetchEventByID(eventID)
}

func (s *eventService) UpdateEvent(event model.Event) error {
	return s.eventRepository.UpdateEventData(event)
}
