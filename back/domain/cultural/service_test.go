package cultural_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"poc2/back/domain/cultural"
	"poc2/back/mocks"
)

func TestCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	title := "Event Title"
	description := "Event Description"
	location, _ := cultural.NewLocation("Address")
	startDate := "2023-10-27"
	finishDate := "2023-10-28"
	durationHours := "2"
	price := cultural.NewPrice("R$10,00")
	isAccessible := true
	organizerID := 1
	image := "image.jpg"

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().SaveEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image).Return(1, nil)

		id, err := service.CreateEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)

		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().SaveEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image).Return(0, errors.New("repository error"))

		id, err := service.CreateEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)

		assert.Error(t, err)
		assert.Equal(t, 0, id)
	})
}

func TestCreateTouristAttraction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	title := "Attraction Title"
	description := "Attraction Description"
	location, _ := cultural.NewLocation("Address")
	workingHours := "09:00-18:00"
	price := cultural.NewPrice("R$20,00")
	isAccessible := true
	organizerID := 1
	image := "image.jpg"

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().SaveTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image).Return(1, nil)

		id, err := service.CreateTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)

		assert.NoError(t, err)
		assert.Equal(t, 1, id)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().SaveTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image).Return(0, errors.New("repository error"))

		id, err := service.CreateTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)

		assert.Error(t, err)
		assert.Equal(t, 0, id)
	})
}

func TestGetEventByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	eventID := 1
	expectedEvent := cultural.Event{ID: eventID, Title: "Event Title"}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindEventByID(ctx, eventID).Return(expectedEvent, nil)

		event, err := service.GetEventByID(ctx, eventID)

		assert.NoError(t, err)
		assert.Equal(t, expectedEvent, event)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindEventByID(ctx, eventID).Return(cultural.Event{}, errors.New("not found"))

		event, err := service.GetEventByID(ctx, eventID)

		assert.Error(t, err)
		assert.Equal(t, cultural.Event{}, event)
	})
}

func TestGetTouristAttractionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	attractionID := 1
	expectedAttraction := cultural.TouristAttraction{ID: attractionID, Title: "Attraction Title"}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindTouristAttractionByID(ctx, attractionID).Return(expectedAttraction, nil)

		attraction, err := service.GetTouristAttractionByID(ctx, attractionID)

		assert.NoError(t, err)
		assert.Equal(t, expectedAttraction, attraction)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindTouristAttractionByID(ctx, attractionID).Return(cultural.TouristAttraction{}, errors.New("not found"))

		attraction, err := service.GetTouristAttractionByID(ctx, attractionID)

		assert.Error(t, err)
		assert.Equal(t, cultural.TouristAttraction{}, attraction)
	})
}

func TestUpdateEventByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	id := 1
	title := "Updated Title"
	description := "Updated Description"
	location, _ := cultural.NewLocation("Address")
	startDate := "2023-10-29"
	finishDate := "2023-10-30"
	durationHours := "3"
	price := cultural.NewPrice("R$15,00")
	isAccessible := false
	organizerID := 1
	image := "updated.jpg"

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image).Return(nil)

		err := service.UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image).Return(errors.New("update error"))

		err := service.UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)

		assert.Error(t, err)
	})
}

func TestUpdateTouristAttractionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	id := 1
	title := "Updated Title"
	description := "Updated Description"
	location, _ := cultural.NewLocation("Updated Address")
	workingHours := "10:00-19:00"
	price := cultural.NewPrice("R$25,00")
	isAccessible := false
	organizerID := 1
	image := "updated.jpg"

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image).Return(nil)

		err := service.UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image).Return(errors.New("update error"))

		err := service.UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)

		assert.Error(t, err)
	})
}

func TestDeleteEventByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	id := 1

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().DeleteEventByID(ctx, id).Return(nil)

		err := service.DeleteEventByID(ctx, id)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().DeleteEventByID(ctx, id).Return(errors.New("delete error"))

		err := service.DeleteEventByID(ctx, id)

		assert.Error(t, err)
	})
}

func TestDeleteTouristAttractionByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	id := 1

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().DeleteTouristAttractionByID(ctx, id).Return(nil)

		err := service.DeleteTouristAttractionByID(ctx, id)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().DeleteTouristAttractionByID(ctx, id).Return(errors.New("delete error"))

		err := service.DeleteTouristAttractionByID(ctx, id)

		assert.Error(t, err)
	})
}

func TestGetEventsIDsByOrganizer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	organizerID := 1
	expectedIDs := []int{1, 2, 3}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindEventsIDsByOrganizer(ctx, organizerID).Return(expectedIDs, nil)

		ids, err := service.GetEventsIDsByOrganizer(ctx, organizerID)

		assert.NoError(t, err)
		assert.Equal(t, expectedIDs, ids)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindEventsIDsByOrganizer(ctx, organizerID).Return(nil, errors.New("error"))

		ids, err := service.GetEventsIDsByOrganizer(ctx, organizerID)

		assert.Error(t, err)
		assert.Nil(t, ids)
	})
}

func TestGetTouristAttractionsIDsByOrganizer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	organizerID := 1
	expectedIDs := []int{1, 2, 3}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindTouristAttractionsIDsByOrganizer(ctx, organizerID).Return(expectedIDs, nil)

		ids, err := service.GetTouristAttractionsIDsByOrganizer(ctx, organizerID)

		assert.NoError(t, err)
		assert.Equal(t, expectedIDs, ids)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindTouristAttractionsIDsByOrganizer(ctx, organizerID).Return(nil, errors.New("error"))

		ids, err := service.GetTouristAttractionsIDsByOrganizer(ctx, organizerID)

		assert.Error(t, err)
		assert.Nil(t, ids)
	})
}

func TestGetAllEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	expectedEvents := []cultural.Event{{ID: 1, Title: "Event 1"}, {ID: 2, Title: "Event 2"}}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindAllEvents(ctx).Return(expectedEvents, nil)

		events, err := service.GetAllEvents(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedEvents, events)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindAllEvents(ctx).Return(nil, errors.New("error"))

		events, err := service.GetAllEvents(ctx)

		assert.Error(t, err)
		assert.Nil(t, events)
	})
}

func TestGetAllTouristAttractions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCulturalRepository(ctrl)
	service := cultural.NewService(mockRepo)

	ctx := context.Background()
	expectedAttractions := []cultural.TouristAttraction{{ID: 1, Title: "Attraction 1"}, {ID: 2, Title: "Attraction 2"}}

	t.Run("success", func(t *testing.T) {
		mockRepo.EXPECT().FindAllTouristAttractions(ctx).Return(expectedAttractions, nil)

		attractions, err := service.GetAllTouristAttractions(ctx)

		assert.NoError(t, err)
		assert.Equal(t, expectedAttractions, attractions)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.EXPECT().FindAllTouristAttractions(ctx).Return(nil, errors.New("error"))

		attractions, err := service.GetAllTouristAttractions(ctx)

		assert.Error(t, err)
		assert.Nil(t, attractions)
	})
}
