package comment

import (
	"context"
	"errors"
	"fmt"

	"poc2/back/domain/comment"
	"poc2/back/domain/cultural"
	"poc2/back/interface/model"
)

const (
	CulturalTypeEvent             = "event"
	CulturalTypeTouristAttraction = "tourist_attraction"
)

type UseCase interface {
	// CreateComment creates a new comment entry
	CreateComment(ctx context.Context, data model.CreateCommentRequest) error

	// GetComments retrieves commentaries for a specific cultural entry
	GetComments(ctx context.Context, culturalID int, culturalType string) (model.GetCommentsResponse, error)
}

type commentaryUseCase struct {
	commentaryService comment.Service
	culturalService   cultural.Service
}

func NewUseCase(commentaryService comment.Service, culturalService cultural.Service) UseCase {
	return &commentaryUseCase{
		commentaryService: commentaryService,
		culturalService:   culturalService,
	}
}

func (uc *commentaryUseCase) CreateComment(ctx context.Context, data model.CreateCommentRequest) error {
	err := uc.commentaryService.CreateComment(ctx, data.CulturalID, data.CulturalType, data.UserID, data.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *commentaryUseCase) GetComments(ctx context.Context, culturalID int, culturalType string) (model.GetCommentsResponse, error) {
	switch culturalType {
	case CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

		if event.ID == 0 {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

	case CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

		if attraction.ID == 0 {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}
	}

	comments, err := uc.commentaryService.GetComments(ctx, culturalID, culturalType)
	if err != nil {
		fmt.Println(err)
		return model.GetCommentsResponse{}, errors.New("comments not found")
	}

	var response []model.Comment
	for _, comment := range comments {
		response = append(response, model.Comment{
			ID:           comment.ID,
			CulturalID:   comment.CulturalID,
			CulturalType: comment.CulturalType,
			UserName:     comment.UserName,
			Comment:      comment.Comment,
			CreatedAt:    comment.CreatedAt,
		})
	}

	return model.GetCommentsResponse{Comments: response}, nil
}
