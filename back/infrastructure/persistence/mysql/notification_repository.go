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

func (r *notificationRepository) FindByUserID(ctx context.Context, userID int) ([]notification.NotificationCulturalList, error) {
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
		notifications = append(notifications, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for i, n := range notifications {
		if n.ID == 0 && n.Type == notification.NotificationTypeUpdated {
			notificationID, err := insertNotifications(ctx, r.db, userID, n, 1)
			if err != nil {
				return nil, err
			}
			notifications[i].ID = notificationID
		}
	}

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
