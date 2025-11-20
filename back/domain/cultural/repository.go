package cultural

import (
	"context"
)

type Repository interface {

	// SaveEvent saves a new cultural event
	SaveEvent(ctx context.Context, title, description, location string,
		startDate, finishDate, durationHours string, price string, isAccessible bool, organizerID int, image string) (int, error)

	// SaveTouristAttraction saves a new cultural tourist attraction
	SaveTouristAttraction(ctx context.Context, title, description, location, workingHours string,
		price string, isAccessible bool, organizerID int, image string) (int, error)

	// FindEventByID retrieves a cultural event by its ID
	FindEventByID(ctx context.Context, id int) (Event, error)

	// FindTouristAttractionByID retrieves a cultural tourist attraction by its ID
	FindTouristAttractionByID(ctx context.Context, id int) (TouristAttraction, error)

	// UpdateEventByID updates a cultural event by its ID
	UpdateEventByID(ctx context.Context, id int, title, description, location string,
		startDate, finishDate, durationHours string, price string, isAccessible bool, organizerID int, image string) error

	// UpdateTouristAttractionByID updates a cultural tourist attraction by its ID
	UpdateTouristAttractionByID(ctx context.Context, id int, title, description, location, workingHours string,
		price string, isAccessible bool, organizerID int, image string) error

	// DeleteEventByID deletes a cultural event by its ID
	DeleteEventByID(ctx context.Context, id int) error

	// DeleteTouristAttractionByID deletes a cultural tourist attraction by its ID
	DeleteTouristAttractionByID(ctx context.Context, id int) error

	// FindEventsIDsByOrganizer retrieves all event IDs organized by a specific user
	FindEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error)

	// FindTouristAttractionsIDsByOrganizer retrieves all tourist attraction IDs organized by a specific user
	FindTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error)

	// FindAllEvents retrieves all cultural events
	FindAllEvents(ctx context.Context) ([]Event, error)

	// FindAllTouristAttractions retrieves all cultural tourist attractions
	FindAllTouristAttractions(ctx context.Context) ([]TouristAttraction, error)
}
