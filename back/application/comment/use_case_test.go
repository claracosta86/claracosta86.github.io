package comment_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	applicationComment "poc2/back/application/comment"
	domainComment "poc2/back/domain/comment"
	domainCultural "poc2/back/domain/cultural"
	"poc2/back/interface/model"
	"poc2/back/mocks"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentService := mocks.NewMockCommentService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationComment.NewUseCase(mockCommentService, mockCulturalService)

	ctx := context.Background()
	req := model.CreateCommentRequest{
		CulturalID:   1,
		CulturalType: "event",
		UserID:       1,
		Comment:      "Great event!",
	}

	t.Run("success", func(t *testing.T) {
		mockCommentService.EXPECT().CreateComment(
			ctx,
			req.CulturalID,
			req.CulturalType,
			req.UserID,
			gomock.Any(), // CommentContent
		).Return(nil)

		err := useCase.CreateComment(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("invalid comment content", func(t *testing.T) {
		invalidReq := req
		invalidReq.Comment = "" // Empty comment might be invalid

		err := useCase.CreateComment(ctx, invalidReq)
		assert.Error(t, err)
	})

	t.Run("service error", func(t *testing.T) {
		mockCommentService.EXPECT().CreateComment(
			ctx,
			req.CulturalID,
			req.CulturalType,
			req.UserID,
			gomock.Any(),
		).Return(errors.New("service error"))

		err := useCase.CreateComment(ctx, req)
		assert.Error(t, err)
	})
}

func TestGetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentService := mocks.NewMockCommentService(ctrl)
	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	useCase := applicationComment.NewUseCase(mockCommentService, mockCulturalService)

	ctx := context.Background()
	culturalID := 1

	t.Run("get event comments success", func(t *testing.T) {
		culturalType := domainComment.CulturalTypeEvent
		event := domainCultural.Event{ID: culturalID}
		content, _ := domainComment.NewCommentContent("Nice!")
		comments := []domainComment.Comment{
			{
				ID:           1,
				CulturalID:   culturalID,
				CulturalType: culturalType,
				UserName:     "User",
				Comment:      content,
				CreatedAt:    time.Now().Format(time.RFC3339),
			},
		}

		mockCulturalService.EXPECT().GetEventByID(ctx, culturalID).Return(event, nil)
		mockCommentService.EXPECT().GetComments(ctx, culturalID, culturalType).Return(comments, nil)

		resp, err := useCase.GetComments(ctx, culturalID, culturalType)
		assert.NoError(t, err)
		assert.Len(t, resp.Comments, 1)
		assert.Equal(t, "Nice!", resp.Comments[0].Comment)
	})

	t.Run("get attraction comments success", func(t *testing.T) {
		culturalType := domainComment.CulturalTypeTouristAttraction
		attraction := domainCultural.TouristAttraction{ID: culturalID}
		comments := []domainComment.Comment{}

		mockCulturalService.EXPECT().GetTouristAttractionByID(ctx, culturalID).Return(attraction, nil)
		mockCommentService.EXPECT().GetComments(ctx, culturalID, culturalType).Return(comments, nil)

		resp, err := useCase.GetComments(ctx, culturalID, culturalType)
		assert.NoError(t, err)
		assert.Len(t, resp.Comments, 0)
	})

	t.Run("cultural not found (event)", func(t *testing.T) {
		culturalType := domainComment.CulturalTypeEvent
		mockCulturalService.EXPECT().GetEventByID(ctx, culturalID).Return(domainCultural.Event{}, errors.New("not found"))

		_, err := useCase.GetComments(ctx, culturalID, culturalType)
		assert.Error(t, err)
		assert.Equal(t, "cultural not found", err.Error())
	})

	t.Run("cultural not found (attraction)", func(t *testing.T) {
		culturalType := domainComment.CulturalTypeTouristAttraction
		mockCulturalService.EXPECT().GetTouristAttractionByID(ctx, culturalID).Return(domainCultural.TouristAttraction{}, errors.New("not found"))

		_, err := useCase.GetComments(ctx, culturalID, culturalType)
		assert.Error(t, err)
		assert.Equal(t, "cultural not found", err.Error())
	})

	t.Run("invalid cultural type", func(t *testing.T) {
		_, err := useCase.GetComments(ctx, culturalID, "invalid")
		assert.Error(t, err)
		assert.Equal(t, "invalid cultural type", err.Error())
	})
}
