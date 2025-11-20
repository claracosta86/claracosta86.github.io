package mysql

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"strings"

	"github.com/nleof/goyesql"

	"poc2/back/domain/notification"
)

var (
	//go:embed queries/notification.sql
	notificationEmbed   []byte
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

	rows, err := r.db.QueryContext(ctx, notificationQueries["fetch-notifications-by-user-id"], userID, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notifications := make([]notification.NotificationCulturalList, 0)
	for rows.Next() {
		var n notification.NotificationCulturalList
		if err := rows.Scan(&n.ID,
			&n.Title,
			&n.CulturalType,
			&n.CulturalID,
			&n.Type,
		); err != nil {
			return nil, err
		}
		if n.Type == notification.NotificationTypeUpdated {
			notificationID, err := insertNotifications(ctx, r.db, userID, n, 1)
			if err != nil {
				return nil, err
			}
			n.ID = notificationID
		}

		notifications = append(notifications, n)
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

func insertNotifications(ctx context.Context, db *sql.DB, userID int, notification notification.NotificationCulturalList, seen int) (int, error) {
	result, err := db.ExecContext(ctx, notificationQueries["create-notification"],
		userID,
		notification.CulturalID,
		notification.CulturalType,
		notification.Type,
		seen,
	)
	if err != nil {
		return 0, fmt.Errorf("erro ao inserir notificações em massa: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erro ao obter ID da notificação inserida: %w", err)
	}

	return int(id), nil
}
