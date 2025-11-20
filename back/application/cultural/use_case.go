package cultural

import (
	"context"
	"errors"
	"fmt"

	"poc2/back/domain/cultural"
	"poc2/back/domain/user"
	"poc2/back/interface/model"
)

const (
	CulturalTypeEvent             = "event"
	CulturalTypeTouristAttraction = "tourist_attraction"
)

type UseCase interface {
	// CreateCultural creates a new cultural entry
	CreateCultural(ctx context.Context, data model.CreateCulturalRequest) (model.CreateCulturalResponse, error)

	// GetCultural retrieves a cultural entry by ID
	GetCultural(ctx context.Context, id int, culturalType string) (model.CulturalResponse, error)

	// UpdateCultural updates a cultural entry by ID
	UpdateCultural(ctx context.Context, data model.UpdateCulturalRequest) error

	// DeleteCultural deletes a cultural entry by ID
	DeleteCultural(ctx context.Context, id int, culturalType string) error

	// GetAllCulturais retrieves all cultural events and attractions
	GetAllCulturais(ctx context.Context) (model.AllCulturaisResponse, error)

	// GetHomeCulturais retrieves cultural events and attractions for home display
	GetHomeCulturais(ctx context.Context) (model.AllCulturaisResponse, error)
}

type culturalUseCase struct {
	culturalService cultural.Service
	userService     user.Service
}

func NewUseCase(culturalService cultural.Service, userService user.Service) UseCase {
	return &culturalUseCase{
		culturalService: culturalService,
		userService:     userService,
	}
}

func (uc *culturalUseCase) CreateCultural(ctx context.Context, data model.CreateCulturalRequest) (model.CreateCulturalResponse, error) {
	switch data.Type {
	case CulturalTypeEvent:
		id, err := uc.culturalService.CreateEvent(ctx, data.Title, data.Description, data.Location,
			data.Event.StartDate, data.Event.EndDate, data.Event.WorkingHours, data.Price, data.IsAccessible, data.OrganizerID, data.Image)
		if err != nil {
			return model.CreateCulturalResponse{}, err
		}
		return model.CreateCulturalResponse{
			ID:   id,
			Type: CulturalTypeEvent,
		}, nil
	case CulturalTypeTouristAttraction:
		id, err := uc.culturalService.CreateTouristAttraction(ctx, data.Title, data.Description, data.Location,
			data.TouristAttraction.OpenDays, data.TouristAttraction.OpenTime, data.Price, data.IsAccessible, data.OrganizerID, data.Image)
		if err != nil {
			return model.CreateCulturalResponse{}, err
		}
		return model.CreateCulturalResponse{
			ID:   id,
			Type: CulturalTypeTouristAttraction,
		}, nil
	}
	return model.CreateCulturalResponse{}, errors.New("invalid cultural type")
}

func (uc *culturalUseCase) GetCultural(ctx context.Context, id int, culturalType string) (model.CulturalResponse, error) {
	switch culturalType {
	case CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, id)
		return model.CulturalResponse{
			ID:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			Location:     event.Location,
			Price:        event.Price,
			IsAccessible: event.IsAccessible,
			Organizer: model.Organizer{
				ID:    event.OrganizerID,
				Email: event.OrganizerEmail,
			},
			Image: event.Image,
			Event: model.EventDateInformation{
				StartDate:    event.StartDate,
				EndDate:      event.EndDate,
				WorkingHours: event.WorkingHours,
			},
		}, err
	case CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, id)
		fmt.Println(attraction)
		return model.CulturalResponse{
			ID:           attraction.ID,
			Title:        attraction.Title,
			Description:  attraction.Description,
			Location:     attraction.Location,
			Price:        attraction.Price,
			IsAccessible: attraction.IsAccessible,
			Organizer: model.Organizer{
				ID:    attraction.OrganizerID,
				Email: attraction.OrganizerEmail,
			},
			Image: attraction.Image,
			TouristAttraction: model.TouristAttractionHours{
				OpenDays: attraction.OpenDays,
				OpenTime: attraction.OpenTime,
			},
		}, err
	}
	return model.CulturalResponse{}, errors.New("invalid cultural type")
}

func (uc *culturalUseCase) UpdateCultural(ctx context.Context, data model.UpdateCulturalRequest) error {
	switch data.Type {
	case CulturalTypeEvent:
		return uc.culturalService.UpdateEventByID(ctx, data.ID, data.Title, data.Description, data.Location,
			data.Event.StartDate, data.Event.EndDate, data.Event.WorkingHours, data.Price, data.IsAccessible, data.OrganizerID, data.Image)
	case CulturalTypeTouristAttraction:
		return uc.culturalService.UpdateTouristAttractionByID(ctx, data.ID, data.Title, data.Description, data.Location,
			data.TouristAttraction.OpenDays, data.TouristAttraction.OpenTime, data.Price, data.IsAccessible, data.OrganizerID, data.Image)
	}
	return errors.New("invalid cultural type")
}

func (uc *culturalUseCase) DeleteCultural(ctx context.Context, id int, culturalType string) error {
	switch culturalType {
	case CulturalTypeEvent:
		if err := uc.culturalService.DeleteEventByID(ctx, id); err != nil {
			return err
		}

		return uc.userService.RemoveEventFromAllUsers(ctx, []int{id})

	case CulturalTypeTouristAttraction:
		if err := uc.culturalService.DeleteTouristAttractionByID(ctx, id); err != nil {
			return err
		}
		return uc.userService.RemoveTouristAttractionFromAllUsers(ctx, []int{id})
	}

	return errors.New("invalid cultural type")
}

func (uc *culturalUseCase) GetAllCulturais(ctx context.Context) (model.AllCulturaisResponse, error) {
	events, err := uc.culturalService.GetAllEvents(ctx)
	if err != nil {
		return model.AllCulturaisResponse{}, err
	}

	attractions, err := uc.culturalService.GetAllTouristAttractions(ctx)
	if err != nil {
		return model.AllCulturaisResponse{}, err
	}

	return model.AllCulturaisResponse{
		Events:             convertEventsToModel(events),
		TouristAttractions: convertAttractionsToModel(attractions),
	}, nil
}

func (uc *culturalUseCase) GetHomeCulturais(ctx context.Context) (model.AllCulturaisResponse, error) {
	events, err := uc.culturalService.GetAllEvents(ctx)
	if err != nil {
		return model.AllCulturaisResponse{}, err
	}

	attractions, err := uc.culturalService.GetAllTouristAttractions(ctx)
	if err != nil {
		return model.AllCulturaisResponse{}, err
	}

	return model.AllCulturaisResponse{
		Events:             convertEventsToModel(events[:5]),
		TouristAttractions: convertAttractionsToModel(attractions[:5]),
	}, nil
}

func convertEventsToModel(events []cultural.Event) []model.Event {
	var eventModels []model.Event
	for _, event := range events {
		eventModels = append(eventModels, model.Event{
			ID:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			Location:     event.Location,
			Price:        event.Price,
			IsAccessible: event.IsAccessible,
			OrganizerID:  event.OrganizerID,
			Image:        event.Image,
			StartDate:    event.StartDate,
			EndDate:      event.EndDate,
			WorkingHours: event.WorkingHours,
		})
	}
	return eventModels
}

func convertAttractionsToModel(attractions []cultural.TouristAttraction) []model.TouristAttraction {
	var attractionModels []model.TouristAttraction
	for _, attraction := range attractions {
		attractionModels = append(attractionModels, model.TouristAttraction{
			ID:           attraction.ID,
			Title:        attraction.Title,
			Description:  attraction.Description,
			Location:     attraction.Location,
			Price:        attraction.Price,
			IsAccessible: attraction.IsAccessible,
			OrganizerID:  attraction.OrganizerID,
			Image:        attraction.Image,
			OpenDays:     attraction.OpenDays,
			OpenTime:     attraction.OpenTime,
		})
	}
	return attractionModels
}
