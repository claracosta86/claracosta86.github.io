package mysql

import (
	"context"
	"database/sql"
	_ "embed"

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

func (r *notificationRepository) FindByUserID(ctx context.Context, id int) ([]notification.NotificationCulturalList, error) {
	rows, err := r.db.QueryContext(ctx, notificationQueries["fetch-notifications-by-user-id"], id, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notifications []notification.NotificationCulturalList
	for rows.Next() {
		var notification notification.NotificationCulturalList
		if err := rows.Scan(&notification.ID, &notification.Title); err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

	return notifications, nil
}
	
