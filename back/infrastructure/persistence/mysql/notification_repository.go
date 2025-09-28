package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"strings"
	"fmt"

	"github.com/nleof/goyesql"

	"poc2/back/domain/notification"

)

var (
	//go:embed queries/notification.sql
	notificationEmbed []byte
	notificationQueries goyesql.Queries
)

type notificationRepository struct {
	db *sql.DB
}

func init() {
	notificationQueries = goyesql.MustParseBytes(notificationEmbed)
}

// NewNotificationRepository creates a new MySQL notification repository
func NewNotificationRepository(db *sql.DB) notification.Repository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) FindByUserIDAndFavorites(ctx context.Context, userID int, favoritesMap map[int]string) ([]notification.NotificationCulturalList, error) {
	eventsIDs := make([]any, 0)
	touristAttractionIDs := make([]any, 0)
	for favoriteID, favoriteType := range favoritesMap {
		if favoriteType == "event" {
			eventsIDs = append(eventsIDs, favoriteID)
		} else if favoriteType == "tourist_attraction" {
			touristAttractionIDs = append(touristAttractionIDs, favoriteID)
		}	
	}
fmt.Printf("Fetching notifications for user %d with favorite events: %v and favorite tourist attractions: %v\n", userID, eventsIDs, touristAttractionIDs)
	if len(eventsIDs) == 0 && len(touristAttractionIDs) == 0 {
        return nil, nil 
    }

	baseQuery := notificationQueries["fetch-notifications-by-user-id"]
    
    eventsPlaceholders := ""
    if len(eventsIDs) > 0 {
        eventsPlaceholders = "?" + strings.Repeat(",?", len(eventsIDs)-1)
    } else {
        eventsPlaceholders = "NULL" 
    }

    touristAttractionPlaceholders := ""
    if len(touristAttractionIDs) > 0 {
        touristAttractionPlaceholders = "?" + strings.Repeat(",?", len(touristAttractionIDs)-1)
    } else {
        touristAttractionPlaceholders = "NULL"
    }

    finalQuery := fmt.Sprintf(baseQuery, eventsPlaceholders, touristAttractionPlaceholders)

    args := make([]any, 0, 1+len(eventsIDs)+1+len(touristAttractionIDs))
    args = append(args, userID)
    args = append(args, eventsIDs...)
    args = append(args, userID)
    args = append(args, touristAttractionIDs...)

   rows, err := r.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notifications []notification.NotificationCulturalList
	for rows.Next() {
		var notification notification.NotificationCulturalList
		if err := rows.Scan(&notification.ID, 
			&notification.Title, 
			&notification.CulturalType,
			&notification.CulturalID,
			&notification.Type,
		); err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

	fmt.Printf("Retrieved %d notifications for user %d\n", len(notifications), userID)

	return notifications, nil
}
	
func (r *notificationRepository) MarkAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	if len(notificationIDs) == 0 {
		return nil 
	}

	placeholders := strings.Repeat("?,", len(notificationIDs))
	placeholders = strings.TrimRight(placeholders, ",")
	query := fmt.Sprintf(notificationQueries["mark-notifications-as-seen"], placeholders)
	args := make([]any, 0, len(notificationIDs)+1)
	for _, id := range notificationIDs {
		args = append(args, id)
	}

	args = append(args, userID)
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}