package cultural

import (
	"context"
	"errors"

	"poc2/back/interface/model"
	"poc2/back/domain/cultural"
	"poc2/back/domain/user"

)

const (
	CulturalTypeEvent          = "event"
	CulturalTypeTouristAttraction = "attraction"
)

type UseCase interface {
	// CreateCultural creates a new cultural entry
	CreateCultural(ctx context.Context, data model.CreateCulturalRequest) error

	// GetCultural retrieves a cultural entry by ID
	GetCultural(ctx context.Context, id int, culturalType string) (model.CulturalResponse, error)

	// UpdateCultural updates a cultural entry by ID
	UpdateCultural(ctx context.Context, id int, data model.UpdateCulturalRequest) error

	// DeleteCultural deletes a cultural entry by ID
	DeleteCultural(ctx context.Context, id int, culturalType string) error
}

type culturalUseCase struct {
	culturalService cultural.Service
	userService      user.Service
}

func NewUseCase(culturalService cultural.Service, userService user.Service) UseCase {
	return &culturalUseCase{
		culturalService: culturalService,
		userService:     userService,
	}
}

func (uc *culturalUseCase) CreateCultural(ctx context.Context, data model.CreateCulturalRequest) error {
	switch data.Type {
	case CulturalTypeEvent:
		return uc.culturalService.CreateEvent(ctx, data.Title, data.Description, data.Location,
			data.Event.StartDate, data.Event.EndDate, data.Event.DurationTime, data.Price, data.IsAccessible, data.Organizer.ID, data.Image)
	case CulturalTypeTouristAttraction:
		return uc.culturalService.CreateTouristAttraction(ctx, data.Title, data.Description, data.Location,
			data.TouristAttraction.OpenDays, data.TouristAttraction.OpenTime, data.Price, data.IsAccessible, data.Organizer.ID, data.Image)
	}
	return errors.New("invalid cultural type")
}

func (uc *culturalUseCase) GetCultural(ctx context.Context, id int, culturalType string) (model.CulturalResponse, error) {
	switch culturalType {
	case CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, id)
		return model.CulturalResponse{
			ID:    event.ID,
			Title: event.Title,
			Description: event.Description,
			Location: event.Location,
			Price: event.Price,
			IsAccessible: event.IsAccessible,
			Organizer: model.Organizer{
				ID:    event.OrganizerID,
				Email: event.OrganizerEmail,
			},
			Image: event.Image,
			Event: model.EventDateInformation{	
				StartDate:   event.StartDate,
				EndDate:     event.EndDate,
				DurationTime: event.DurationTime,
			},
		}, err
	case CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, id)
		return model.CulturalResponse{
			ID:    attraction.ID,
			Title: attraction.Title,
			Description: attraction.Description,
			Location: attraction.Location,
			Price: attraction.Price,
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

func (uc *culturalUseCase) UpdateCultural(ctx context.Context, id int, data model.UpdateCulturalRequest) error {
	switch data.Type {
	case CulturalTypeEvent:
		return uc.culturalService.UpdateEventByID(ctx, id, data.Title, data.Description, data.Location,
			data.Event.StartDate, data.Event.EndDate, data.Event.DurationTime, data.Price, data.IsAccessible, data.OrganizerID, data.Image)
	case CulturalTypeTouristAttraction:
		return uc.culturalService.UpdateTouristAttractionByID(ctx, id, data.Title, data.Description, data.Location,
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
