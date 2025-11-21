package comment_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"poc2/back/domain/comment"
	"poc2/back/mocks"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCommentRepository(ctrl)
	service := comment.NewService(mockRepo)

	ctx := context.Background()
	culturalID := 1
	culturalType := "event"
	userID := 1
	commentContent, _ := comment.NewCommentContent("Great event!")
	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().SaveComment(ctx, culturalID, culturalType, userID, commentContent).Return(nil)

		err := service.CreateComment(ctx, culturalID, culturalType, userID, commentContent)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().SaveComment(ctx, culturalID, culturalType, userID, commentContent).Return(errors.New("repository error"))

		err := service.CreateComment(ctx, culturalID, culturalType, userID, commentContent)

		assert.Error(t, err)
	})
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCommentRepository(ctrl)
	service := comment.NewService(mockRepo)

	ctx := context.Background()
	culturalID := 1
	culturalType := "event"
	comments, _ := comment.NewCommentContent("Perfeito!")
	expectedComments := comment.Comments{
		{ID: 1, CulturalID: 1, CulturalType: "event", UserName: "User1", Comment: comments},
		{ID: 2, CulturalID: 1, CulturalType: "event", UserName: "User2", Comment: comments},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindCommentsByCultural(ctx, culturalID, culturalType).Return(expectedComments, nil)

		comments, err := service.GetComments(ctx, culturalID, culturalType)

		assert.NoError(t, err)
		assert.Equal(t, expectedComments, comments)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindCommentsByCultural(ctx, culturalID, culturalType).Return(comment.Comments{}, errors.New("repository error"))

		comments, err := service.GetComments(ctx, culturalID, culturalType)

		assert.Error(t, err)
		assert.Equal(t, comment.Comments{}, comments)
	})
}
