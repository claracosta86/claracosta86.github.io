package mysql

import (
	"context"
	"errors"
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/nleof/goyesql"

	"poc2/back/domain/cultural"

)

var (
	//go:embed queries/cultural.sql
	culturalEmbed []byte
	culturalQueries goyesql.Queries
)

type culturalRepository struct {
	db *sql.DB
}

func init() {
	culturalQueries = goyesql.MustParseBytes(culturalEmbed)
}

// NewCulturalRepository creates a new MySQL cultural repository
func NewCulturalRepository(db *sql.DB) cultural.Repository {
	return &culturalRepository{
		db: db,
	}
}


func (r *culturalRepository) SaveEvent(ctx context.Context, title, description, location string,
	startDate, finishDate, duration, price string, isAccessible bool, organizerID int, image string) (int, error) {
	result, err := r.db.ExecContext(ctx, culturalQueries["create-event"],
		title,
		description,
		location,
		startDate, 
		finishDate, 
		price,
		isAccessible,
		organizerID,
		image,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to save event: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return int(id), nil
}

func (r *culturalRepository) SaveTouristAttraction(ctx context.Context, title, description, location, openDays, openTime string,
	price string, isAccessible bool, organizerID int, image string) (int, error) {

	result, err := r.db.ExecContext(ctx, culturalQueries["create-tourist-attraction"],
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
	if err != nil {
		return 0, fmt.Errorf("failed to save tourist attraction: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	fmt.Printf("Inserted tourist attraction with ID: %d\n", id)
	return int(id), nil
}

func (r *culturalRepository) FindEventByID(ctx context.Context, id int) (cultural.Event, error) {
	var event cultural.Event
	err := r.db.QueryRowContext(ctx, culturalQueries["fetch-event-by-id"], id).Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.StartDate,
		&event.EndDate,
		&event.DurationTime,
		&event.Price,
		&event.IsAccessible,
		&event.OrganizerID,
		&event.OrganizerEmail,
		&event.Image,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found for event ID:", id)
			return cultural.Event{}, errors.New("cultural event not found")
		}
					fmt.Println(err)

		return cultural.Event{}, err
	}

	return event, nil
}

func (r *culturalRepository) FindTouristAttractionByID(ctx context.Context, id int) (cultural.TouristAttraction, error) {
	var attraction cultural.TouristAttraction
	err := r.db.QueryRowContext(ctx, culturalQueries["fetch-tourist-attraction-by-id"], id).Scan(
		&attraction.ID,
		&attraction.Title,
		&attraction.Description,
		&attraction.Location,
		&attraction.OpenDays,
		&attraction.OpenTime,
		&attraction.Price,
		&attraction.IsAccessible,
		&attraction.OrganizerID,
		&attraction.OrganizerEmail,
		&attraction.Image,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return cultural.TouristAttraction{}, errors.New("cultural tourist attraction not found")
		}
		return cultural.TouristAttraction{}, err
	}

	return attraction, nil
}

func (r *culturalRepository) UpdateEventByID(ctx context.Context, id int, title, description, location string,
	startDate, finishDate, duration string, price string, isAccessible bool, organizerID int, image string) error {

	_, err := r.db.ExecContext(ctx, culturalQueries["update-event"], id,
		title,
		description,
		location,
		startDate,
		finishDate,
		price,
		isAccessible,
		organizerID,
		image,
	)
	return fmt.Errorf("failed to update event: %w", err)
}

func (r *culturalRepository) UpdateTouristAttractionByID(ctx context.Context, id int, title, description, location, openDays, openTime string,
	price string, isAccessible bool, organizerID int, image string) error {

	_, err := r.db.ExecContext(ctx, culturalQueries["update-tourist-attraction"], id,
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
	return fmt.Errorf("failed to update tourist attraction: %w", err)
}

func (r *culturalRepository) DeleteEventByID(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, culturalQueries["delete-event"], id)

	return fmt.Errorf("failed to delete event: %w", err)
}

func (r *culturalRepository) DeleteTouristAttractionByID(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, culturalQueries["delete-tourist-attraction"], id)
	return fmt.Errorf("failed to delete tourist attraction: %w", err)
}

func (r *culturalRepository) FindEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	var eventIDs []int
	rows, err := r.db.QueryContext(ctx, culturalQueries["fetch-events-ids-by-organizer"], organizerID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events IDs by organizer: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan event ID: %w", err)
		}
		eventIDs = append(eventIDs, id)
	}

	return eventIDs, nil
}

func (r *culturalRepository) FindTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	var attractionIDs []int
	rows, err := r.db.QueryContext(ctx, culturalQueries["fetch-tourist-attractions-ids-by-organizer"], organizerID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tourist attractions IDs by organizer: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan tourist attraction ID: %w", err)
		}
		attractionIDs = append(attractionIDs, id)
	}

	return attractionIDs, nil
}