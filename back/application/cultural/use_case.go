package cultural

import (
	"context"
	"errors"

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
	GetCultural(ctx context.Context, id int, culturalType string) (model.GetCulturalResponse, error)

	// UpdateCultural updates a cultural entry by ID
	UpdateCultural(ctx context.Context, data model.UpdateCulturalRequest) error

	// DeleteCultural deletes a cultural entry by ID
	DeleteCultural(ctx context.Context, id int, culturalType string) error

	// GetAllCulturais retrieves all cultural events and attractions
	GetAllCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error)

	// GetHomeCulturais retrieves cultural events and attractions for home display
	GetHomeCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error)
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
	price := cultural.NewPrice(data.Price)
	location, err := cultural.NewLocation(data.Location)
	if err != nil {
		return model.CreateCulturalResponse{}, err
	}

	switch data.Type {
	case CulturalTypeEvent:
		id, err := uc.culturalService.CreateEvent(ctx, data.Title, data.Description, location,
			data.Event.StartDate, data.Event.EndDate, data.Event.DurationHours, price, data.IsAccessible, data.OrganizerID, data.Image)
		if err != nil {
			return model.CreateCulturalResponse{}, err
		}
		return model.CreateCulturalResponse{
			ID:   id,
			Type: CulturalTypeEvent,
		}, nil
	case CulturalTypeTouristAttraction:
		id, err := uc.culturalService.CreateTouristAttraction(ctx, data.Title, data.Description, location,
			data.TouristAttraction.WorkingHours, price, data.IsAccessible, data.OrganizerID, data.Image)
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

func (uc *culturalUseCase) GetCultural(ctx context.Context, id int, culturalType string) (model.GetCulturalResponse, error) {
	switch culturalType {
	case CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, id)
		if err != nil {
			return model.GetCulturalResponse{}, err
		}

		return model.GetCulturalResponse{
			ID:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			Location:     event.Location.String(),
			Price:        event.Price.String(),
			IsAccessible: event.IsAccessible,
			Organizer: model.Organizer{
				ID:    event.OrganizerID,
				Email: event.OrganizerEmail,
			},
			Image: event.Image,
			Event: model.EventHours{
				StartDate:     event.StartDate,
				EndDate:       event.EndDate,
				DurationHours: event.DurationHours,
			},
		}, err
	case CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, id)
		if err != nil {
			return model.GetCulturalResponse{}, err
		}

		return model.GetCulturalResponse{
			ID:           attraction.ID,
			Title:        attraction.Title,
			Description:  attraction.Description,
			Location:     attraction.Location.String(),
			Price:        attraction.Price.String(),
			IsAccessible: attraction.IsAccessible,
			Organizer: model.Organizer{
				ID:    attraction.OrganizerID,
				Email: attraction.OrganizerEmail,
			},
			Image: attraction.Image,
			TouristAttraction: model.TouristAttractionHours{
				WorkingHours: attraction.WorkingHours,
			},
		}, err
	}
	return model.GetCulturalResponse{}, errors.New("invalid cultural type")
}

func (uc *culturalUseCase) UpdateCultural(ctx context.Context, data model.UpdateCulturalRequest) error {
	price := cultural.NewPrice(data.Price)
	location, err := cultural.NewLocation(data.Location)
	if err != nil {
		return err
	}

	switch data.Type {
	case CulturalTypeEvent:
		return uc.culturalService.UpdateEventByID(ctx, data.ID, data.Title, data.Description, location,
			data.Event.StartDate, data.Event.EndDate, data.Event.DurationHours, price, data.IsAccessible, data.OrganizerID, data.Image)
	case CulturalTypeTouristAttraction:
		return uc.culturalService.UpdateTouristAttractionByID(ctx, data.ID, data.Title, data.Description, location,
			data.TouristAttraction.WorkingHours, price, data.IsAccessible, data.OrganizerID, data.Image)
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

func (uc *culturalUseCase) GetAllCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error) {
	events, err := uc.culturalService.GetAllEvents(ctx)
	if err != nil {
		return model.GetAllCulturaisResponse{}, err
	}

	attractions, err := uc.culturalService.GetAllTouristAttractions(ctx)
	if err != nil {
		return model.GetAllCulturaisResponse{}, err
	}

	return model.GetAllCulturaisResponse{
		Events:             convertEventsToModel(events),
		TouristAttractions: convertAttractionsToModel(attractions),
	}, nil
}

func (uc *culturalUseCase) GetHomeCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error) {
	events, err := uc.culturalService.GetAllEvents(ctx)
	if err != nil {
		return model.GetAllCulturaisResponse{}, err
	}

	attractions, err := uc.culturalService.GetAllTouristAttractions(ctx)
	if err != nil {
		return model.GetAllCulturaisResponse{}, err
	}

	return model.GetAllCulturaisResponse{
		Events:             convertEventsToModel(events[:5]),
		TouristAttractions: convertAttractionsToModel(attractions[:5]),
	}, nil
}

func convertEventsToModel(events []cultural.Event) []model.Event {
	var eventModels []model.Event
	for _, event := range events {
		eventModels = append(eventModels, model.Event{
			ID:            event.ID,
			Title:         event.Title,
			Description:   event.Description,
			Location:      event.Location.String(),
			Price:         event.Price.String(),
			IsAccessible:  event.IsAccessible,
			OrganizerID:   event.OrganizerID,
			Image:         event.Image,
			StartDate:     event.StartDate,
			EndDate:       event.EndDate,
			DurationHours: event.DurationHours,
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
			Location:     attraction.Location.String(),
			Price:        attraction.Price.String(),
			IsAccessible: attraction.IsAccessible,
			OrganizerID:  attraction.OrganizerID,
			Image:        attraction.Image,
			WorkingHours: attraction.WorkingHours,
		})
	}
	return attractionModels
}
