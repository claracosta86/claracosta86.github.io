package commentary

import (
	"context"
	"errors"

	"poc2/back/interface/model"
	"poc2/back/domain/commentary"
	"poc2/back/domain/cultural"

)

const (
	CulturalTypeEvent          = "event"
	CulturalTypeTouristAttraction = "tourist_attraction"
)

type UseCase interface {
	// CreateCommentary creates a new commentary entry
	CreateCommentary(ctx context.Context, data model.CreateCommentaryRequest) error
	
	// GetCommentaries retrieves commentaries for a specific cultural entry
	GetCommentaries(ctx context.Context, culturalID int, culturalType string) (model.GetCommentariesResponse, error)
}


type commentaryUseCase struct {
	commentaryService commentary.Service
	culturalService  cultural.Service
}

func NewUseCase(commentaryService commentary.Service, culturalService cultural.Service) UseCase {
	return &commentaryUseCase{
		commentaryService: commentaryService,
		culturalService:  culturalService,
	}
}

func (uc *commentaryUseCase) CreateCommentary(ctx context.Context, data model.CreateCommentaryRequest) error {
	err := uc.commentaryService.CreateCommentary(ctx, data.CulturalID, data.CulturalType, data.UserID, data.Commentary)
	if err != nil {
		return err
	}
	return nil
}

func (uc *commentaryUseCase) GetCommentaries(ctx context.Context, culturalID int, culturalType string) (model.GetCommentariesResponse, error) {
	switch culturalType {
	case CulturalTypeEvent:
		event, err := uc.culturalService.GetEventByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentariesResponse{}, errors.New("cultural not found")
		}

		if event.ID == 0 {
			return model.GetCommentariesResponse{}, errors.New("cultural not found")
		}

	case CulturalTypeTouristAttraction:
		attraction, err := uc.culturalService.GetTouristAttractionByID(ctx, culturalID)
		if err != nil {
			return model.GetCommentariesResponse{}, errors.New("cultural not found")
		}

		if attraction.ID == 0 {
			return model.GetCommentariesResponse{}, errors.New("cultural not found")
		}
	}

	commentaries, err := uc.commentaryService.GetCommentaries(ctx, culturalID, culturalType)
	if err != nil {
		return model.GetCommentariesResponse{}, errors.New("commentaries not found")
	}

	var response []model.Commentary
	for _, commentary := range commentaries {
		response = append(response, model.Commentary{
			ID:          commentary.ID,
			CulturalID:  commentary.CulturalID,
			CulturalType: commentary.CulturalType,
			UserName:    commentary.UserName,
			Commentary:  commentary.Commentary,
			CreatedAt:   commentary.CreatedAt,
			UpdatedAt:   commentary.UpdatedAt,
		})
	}

	return model.GetCommentariesResponse{Commentaries: response}, nil
}
