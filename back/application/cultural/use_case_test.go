package cultural_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	applicationCultural "poc2/back/application/cultural"
	domainCultural "poc2/back/domain/cultural"
	"poc2/back/interface/model"
	"poc2/back/mocks"
)

func TestCreateCultural(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()
	req := model.CreateCulturalRequest{
		Title:        "Test Event",
		Description:  "Description",
		Location:     "Location",
		Price:        "10.0",
		IsAccessible: true,
		OrganizerID:  1,
		Image:        "image.jpg",
		Type:         applicationCultural.CulturalTypeEvent,
		Event: model.EventHours{
			StartDate:     "2023-01-01",
			EndDate:       "2023-01-02",
			DurationHours: "2h",
		},
	}

	t.Run("create event success", func(t *testing.T) {
		mockCulturalService.EXPECT().CreateEvent(
			ctx,
			req.Title,
			req.Description,
			gomock.Any(), // Location
			req.Event.StartDate,
			req.Event.EndDate,
			req.Event.DurationHours,
			gomock.Any(), // Price
			req.IsAccessible,
			req.OrganizerID,
			req.Image,
		).Return(1, nil)

		resp, err := useCase.CreateCultural(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, 1, resp.ID)
		assert.Equal(t, applicationCultural.CulturalTypeEvent, resp.Type)
	})

	t.Run("create tourist attraction success", func(t *testing.T) {
		reqAttraction := req
		reqAttraction.Type = applicationCultural.CulturalTypeTouristAttraction
		reqAttraction.TouristAttraction = model.TouristAttractionHours{
			WorkingHours: "9am-5pm",
		}

		mockCulturalService.EXPECT().CreateTouristAttraction(
			ctx,
			reqAttraction.Title,
			reqAttraction.Description,
			gomock.Any(), // Location
			reqAttraction.TouristAttraction.WorkingHours,
			gomock.Any(), // Price
			reqAttraction.IsAccessible,
			reqAttraction.OrganizerID,
			reqAttraction.Image,
		).Return(2, nil)

		resp, err := useCase.CreateCultural(ctx, reqAttraction)
		assert.NoError(t, err)
		assert.Equal(t, 2, resp.ID)
		assert.Equal(t, applicationCultural.CulturalTypeTouristAttraction, resp.Type)
	})

	t.Run("invalid type", func(t *testing.T) {
		invalidReq := req
		invalidReq.Type = "invalid"

		_, err := useCase.CreateCultural(ctx, invalidReq)
		assert.Error(t, err)
	})
}

func TestGetCultural(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()
	id := 1

	t.Run("get event success", func(t *testing.T) {
		event := domainCultural.Event{
			ID:          id,
			Title:       "Event",
			Description: "Desc",
			OrganizerID: 1,
		}
		mockCulturalService.EXPECT().GetEventByID(ctx, id).Return(event, nil)

		resp, err := useCase.GetCultural(ctx, id, applicationCultural.CulturalTypeEvent)
		assert.NoError(t, err)
		assert.Equal(t, event.Title, resp.Title)
	})

	t.Run("get tourist attraction success", func(t *testing.T) {
		attraction := domainCultural.TouristAttraction{
			ID:          id,
			Title:       "Attraction",
			Description: "Desc",
			OrganizerID: 1,
		}
		mockCulturalService.EXPECT().GetTouristAttractionByID(ctx, id).Return(attraction, nil)

		resp, err := useCase.GetCultural(ctx, id, applicationCultural.CulturalTypeTouristAttraction)
		assert.NoError(t, err)
		assert.Equal(t, attraction.Title, resp.Title)
	})

	t.Run("invalid type", func(t *testing.T) {
		_, err := useCase.GetCultural(ctx, id, "invalid")
		assert.Error(t, err)
	})
}

func TestUpdateCultural(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()
	req := model.UpdateCulturalRequest{
		ID:           1,
		Title:        "Updated Title",
		Description:  "Updated Desc",
		Location:     "Updated Loc",
		Price:        "20.0",
		IsAccessible: false,
		OrganizerID:  1,
		Image:        "updated.jpg",
		Type:         applicationCultural.CulturalTypeEvent,
		Event: model.EventHours{
			StartDate:     "2023-02-01",
			EndDate:       "2023-02-02",
			DurationHours: "3h",
		},
	}

	t.Run("update event success", func(t *testing.T) {
		mockCulturalService.EXPECT().UpdateEventByID(
			ctx,
			req.ID,
			req.Title,
			req.Description,
			gomock.Any(),
			req.Event.StartDate,
			req.Event.EndDate,
			req.Event.DurationHours,
			gomock.Any(),
			req.IsAccessible,
			req.OrganizerID,
			req.Image,
		).Return(nil)

		err := useCase.UpdateCultural(ctx, req)
		assert.NoError(t, err)
	})

	t.Run("update tourist attraction success", func(t *testing.T) {
		reqAttraction := req
		reqAttraction.Type = applicationCultural.CulturalTypeTouristAttraction
		reqAttraction.TouristAttraction = model.TouristAttractionHours{
			WorkingHours: "10am-6pm",
		}

		mockCulturalService.EXPECT().UpdateTouristAttractionByID(
			ctx,
			reqAttraction.ID,
			reqAttraction.Title,
			reqAttraction.Description,
			gomock.Any(),
			reqAttraction.TouristAttraction.WorkingHours,
			gomock.Any(),
			reqAttraction.IsAccessible,
			reqAttraction.OrganizerID,
			reqAttraction.Image,
		).Return(nil)

		err := useCase.UpdateCultural(ctx, reqAttraction)
		assert.NoError(t, err)
	})
}

func TestDeleteCultural(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()
	id := 1

	t.Run("delete event success", func(t *testing.T) {
		mockCulturalService.EXPECT().DeleteEventByID(ctx, id).Return(nil)
		mockUserService.EXPECT().RemoveEventFromAllUsers(ctx, []int{id}).Return(nil)

		err := useCase.DeleteCultural(ctx, id, applicationCultural.CulturalTypeEvent)
		assert.NoError(t, err)
	})

	t.Run("delete tourist attraction success", func(t *testing.T) {
		mockCulturalService.EXPECT().DeleteTouristAttractionByID(ctx, id).Return(nil)
		mockUserService.EXPECT().RemoveTouristAttractionFromAllUsers(ctx, []int{id}).Return(nil)

		err := useCase.DeleteCultural(ctx, id, applicationCultural.CulturalTypeTouristAttraction)
		assert.NoError(t, err)
	})
}

func TestGetAllCulturais(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		events := []domainCultural.Event{{ID: 1, Title: "Event 1"}}
		attractions := []domainCultural.TouristAttraction{{ID: 2, Title: "Attraction 1"}}

		mockCulturalService.EXPECT().GetAllEvents(ctx).Return(events, nil)
		mockCulturalService.EXPECT().GetAllTouristAttractions(ctx).Return(attractions, nil)

		resp, err := useCase.GetAllCulturais(ctx)
		assert.NoError(t, err)
		assert.Len(t, resp.Events, 1)
		assert.Len(t, resp.TouristAttractions, 1)
	})

	t.Run("error getting events", func(t *testing.T) {
		mockCulturalService.EXPECT().GetAllEvents(ctx).Return(nil, errors.New("error"))

		_, err := useCase.GetAllCulturais(ctx)
		assert.Error(t, err)
	})
}

func TestGetHomeCulturais(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCulturalService := mocks.NewMockCulturalService(ctrl)
	mockUserService := mocks.NewMockUserService(ctrl)
	useCase := applicationCultural.NewUseCase(mockCulturalService, mockUserService)

	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		events := make([]domainCultural.Event, 6) // Create 6 events to test slicing
		attractions := make([]domainCultural.TouristAttraction, 6)

		mockCulturalService.EXPECT().GetAllEvents(ctx).Return(events, nil)
		mockCulturalService.EXPECT().GetAllTouristAttractions(ctx).Return(attractions, nil)

		resp, err := useCase.GetHomeCulturais(ctx)
		assert.NoError(t, err)
		assert.Len(t, resp.Events, 5)             // Should be sliced to 5
		assert.Len(t, resp.TouristAttractions, 5) // Should be sliced to 5
	})
}
