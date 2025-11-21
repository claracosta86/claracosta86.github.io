package mysql

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"poc2/back/domain/notification"
)

func TestNotificationRepository_FindByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewNotificationRepository(db)
	ctx := context.Background()
	userID := 1

	t.Run("success with existing notifications", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "title", "cultural_type", "cultural_id", "type"}).
			AddRow(1, "Event Title", "event", 10, "commented")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT n.id, COALESCE(e.title, ta.title) AS title, n.cultural_type, n.cultural_id, n.type FROM notifications n")).
			WithArgs(userID, userID).
			WillReturnRows(rows)

		notifications, err := repo.FindByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Len(t, notifications, 1)
		assert.Equal(t, 1, notifications[0].ID)
	})

	t.Run("success with new updated notification (needs insert)", func(t *testing.T) {
		// First query returns a row with ID=0 and Type='updated'
		rows := sqlmock.NewRows([]string{"id", "title", "cultural_type", "cultural_id", "type"}).
			AddRow(0, "Updated Event", "event", 20, "updated")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT n.id, COALESCE(e.title, ta.title) AS title, n.cultural_type, n.cultural_id, n.type FROM notifications n")).
			WithArgs(userID, userID).
			WillReturnRows(rows)

		// Expect insert
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO notifications")).
			WithArgs(userID, 20, "event", notification.NotificationTypeUpdated, 1).
			WillReturnResult(sqlmock.NewResult(100, 1))

		notifications, err := repo.FindByUserID(ctx, userID)
		assert.NoError(t, err)
		assert.Len(t, notifications, 1)
		assert.Equal(t, 100, notifications[0].ID) // ID should be updated from insert
	})
}

func TestNotificationRepository_MarkAsSeen(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewNotificationRepository(db)
	ctx := context.Background()
	userID := 1
	notificationIDs := []int{1, 2}

	t.Run("success", func(t *testing.T) {
		// Query: UPDATE notifications SET seen = 1 WHERE user_id = ? AND id IN ( ?,? )
		// Args: userID, id1, id2
		mock.ExpectExec(regexp.QuoteMeta("UPDATE notifications SET seen = 1 WHERE user_id = ? AND id IN ( ?,? )")).
			WithArgs(userID, 1, 2).
			WillReturnResult(sqlmock.NewResult(0, 2))

		err := repo.MarkAsSeen(ctx, userID, notificationIDs)
		assert.NoError(t, err)
	})
}
