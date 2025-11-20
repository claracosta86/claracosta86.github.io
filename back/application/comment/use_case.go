package comment

import (
	"context"
	"errors"
	"fmt"

	"poc2/back/domain/comment"
	"poc2/back/domain/cultural"
	"poc2/back/interface/model"
)

type UseCase interface {
	// CreateComment creates a new comment entry
	CreateComment(ctx context.Context, data model.CreateCommentRequest) error

	// GetComments retrieves comments for a specific cultural entry
	GetComments(ctx context.Context, culturalID int, culturalType string) (model.GetCommentsResponse, error)
}

type commentUseCase struct {
	commentService  comment.Service
	culturalService cultural.Service
}

func NewUseCase(commentService comment.Service, culturalService cultural.Service) UseCase {
	return &commentUseCase{
		commentService:  commentService,
		culturalService: culturalService,
	}
}

func (uc *commentUseCase) CreateComment(ctx context.Context, data model.CreateCommentRequest) error {
	commentVO, err := comment.NewCommentContent(data.Comment)
	if err != nil {
		return err
	}

	err = uc.commentService.CreateComment(ctx, data.CulturalID, data.CulturalType, data.UserID, commentVO)
	if err != nil {
		return err
	}
	return nil
}

func (uc *commentUseCase) GetComments(ctx context.Context, culturalID int, culturalType string) (model.GetCommentsResponse, error) {
	switch culturalType {
	case comment.CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

		if event.ID == 0 {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

	case comment.CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}

		if attraction.ID == 0 {
			return model.GetCommentsResponse{}, errors.New("cultural not found")
		}
	default:
		return model.GetCommentsResponse{}, errors.New("invalid cultural type")
	}

	comments, err := uc.commentService.GetComments(ctx, culturalID, culturalType)
	if err != nil {
		fmt.Println(err)
		return model.GetCommentsResponse{}, errors.New("comments not found")
	}

	var response []model.Comment
	for _, c := range comments {
		commentContent, err := comment.NewCommentContent(c.Comment.String())
		if err != nil {
			return model.GetCommentsResponse{}, err
		}

		response = append(response, model.Comment{
			ID:           c.ID,
			CulturalID:   c.CulturalID,
			CulturalType: c.CulturalType,
			UserName:     c.UserName,
			Comment:      commentContent.String(),
			CreatedAt:    c.CreatedAt,
		})
	}

	return model.GetCommentsResponse{Comments: response}, nil
}
