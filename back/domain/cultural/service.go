package cultural

import (
	"context"

)

// Service defines the business logic for cultural operations
type Service interface {
	// CreateEvent creates a new cultural event
	CreateEvent(ctx context.Context, title, description, location string,
		startDate, finishDate, workingHours string, price string, isAccessible bool, organizerID int, image string) (int, error)

	// CreateTouristAttraction creates a new cultural tourist attraction
	CreateTouristAttraction(ctx context.Context, title, description, location, openDays, openTime string,
		price string, isAccessible bool, organizerID int, image string) (int, error)

	// GetEventByID retrieves a cultural event by its ID
	GetEventByID(ctx context.Context, id int) (Event, error)

	// GetTouristAttractionByID retrieves a cultural tourist attraction by its ID
	GetTouristAttractionByID(ctx context.Context, id int) (TouristAttraction, error)

	// UpdateEventByID updates a cultural event by its ID
	UpdateEventByID(ctx context.Context, id int, title, description, location string,
		startDate, finishDate, workingHours string, price string, isAccessible bool, organizerID int, image string) error

	// UpdateTouristAttractionByID updates a cultural tourist attraction by its ID
	UpdateTouristAttractionByID(ctx context.Context, id int, title, description, location, openDays, openTime string,
		price string, isAccessible bool, organizerID int, image string) error

	// DeleteEventByID deletes a cultural event by its ID
	DeleteEventByID(ctx context.Context, id int) error

	// DeleteTouristAttractionByID deletes a cultural tourist attraction by its ID
	DeleteTouristAttractionByID(ctx context.Context, id int) error

	// GetEventsIDsByOrganizer retrieves all event IDs organized by a specific user
	GetEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error)

	// GetTouristAttractionsIDsByOrganizer retrieves all tourist attraction IDs organized by a specific user
	GetTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error)
}

type service struct {
	repository Repository
}

// NewService creates a new cultural service
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateEvent(ctx context.Context, title, description, location string,
	startDate, finishDate, workingHours string, price string, isAccessible bool, organizerID int, image string) (int, error) {

	return s.repository.SaveEvent(
		ctx,
		title, 
		description, 
		location, 
		startDate, 
		finishDate, 
		workingHours,
		price, 
		isAccessible, 
		organizerID, 
		image,
	)
}

func (s *service) CreateTouristAttraction(ctx context.Context, title, description, location, openDays, openTime string,
	price string, isAccessible bool, organizerID int, image string) (int, error) {

	return s.repository.SaveTouristAttraction(
		ctx,
		title,
		description,
		location,
		openDays,
		openTime, 
		price, 
		isAccessible, 
		organizerID, 
		image,
	)
}

func (s *service) GetEventByID(ctx context.Context, id int) (Event, error) {
	return  s.repository.FindEventByID(ctx, id)
}

func (s *service) GetTouristAttractionByID(ctx context.Context, id int) (TouristAttraction, error) {
	return s.repository.FindTouristAttractionByID(ctx, id)
}

func (s *service) UpdateEventByID(ctx context.Context, id int, title, description, location string, 
	startDate, finishDate, workingHours, price string, isAccessible bool, organizerID int, image string) error {
	return s.repository.UpdateEventByID(
		ctx, 
		id, 
		title, 
		description, 
		location, 
		startDate, 
		finishDate, 
		workingHours,
		price, 
		isAccessible, 
		organizerID, 
		image,
	)
}

func (s *service) UpdateTouristAttractionByID(ctx context.Context, id int, title, description, location, openDays, openTime string,
	price string, isAccessible bool, organizerID int, image string) error {
	return  s.repository.UpdateTouristAttractionByID(
		ctx, 
		id, 
		title, 
		description, 
		location, 
		openDays, 
		openTime, 
		price, 
		isAccessible, 
		organizerID, 
		image,
	)
}

func (s *service) DeleteEventByID(ctx context.Context, id int) error {
	return s.repository.DeleteEventByID(ctx, id)
}

func (s *service) DeleteTouristAttractionByID(ctx context.Context, id int) error {
	return s.repository.DeleteTouristAttractionByID(ctx, id)
}

func (s *service) GetEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	return s.repository.FindEventsIDsByOrganizer(ctx, organizerID)
}

func (s *service) GetTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	return s.repository.FindTouristAttractionsIDsByOrganizer(ctx, organizerID)
}
