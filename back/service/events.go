package service

import (
	"poc2/back/model"
	"poc2/back/repository"

)

type EventService interface {
	RegisterEvent(event model.Event) error
	GetAllEvents() ([]model.Event, error)
	GetEventByID(eventID int) (model.Event, error)
	UpdateEvent(event model.Event) error
	DeleteEventByID(eventID int) error
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

func (s *eventService) GetAllEvents() ([]model.Event, error) {
	return s.eventRepository.FetchAllEvents()
}

func (s *eventService) GetEventByID(eventID int) (model.Event, error) {
	return s.eventRepository.FetchEventByID(eventID)
}

func (s *eventService) UpdateEvent(event model.Event) error {
	return s.eventRepository.UpdateEventData(event)
}

func (s *eventService) DeleteEventByID(eventID int) error {
	return s.eventRepository.DeleteEventByID(eventID)
}