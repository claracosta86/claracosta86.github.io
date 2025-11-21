package mocks

import (
	context "context"
	reflect "reflect"

	cultural "poc2/back/domain/cultural"

	gomock "go.uber.org/mock/gomock"
)

// MockCulturalService is a mock of Service interface.
type MockCulturalService struct {
	ctrl     *gomock.Controller
	recorder *MockCulturalServiceMockRecorder
}

// MockCulturalServiceMockRecorder is the mock recorder for MockCulturalService.
type MockCulturalServiceMockRecorder struct {
	mock *MockCulturalService
}

// NewMockCulturalService creates a new mock instance.
func NewMockCulturalService(ctrl *gomock.Controller) *MockCulturalService {
	mock := &MockCulturalService{ctrl: ctrl}
	mock.recorder = &MockCulturalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCulturalService) EXPECT() *MockCulturalServiceMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockCulturalService) CreateEvent(ctx context.Context, title, description string, location cultural.Location, startDate, finishDate, durationHours string, price cultural.Price, isAccessible bool, organizerID int, image string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockCulturalServiceMockRecorder) CreateEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockCulturalService)(nil).CreateEvent), ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
}

// CreateTouristAttraction mocks base method.
func (m *MockCulturalService) CreateTouristAttraction(ctx context.Context, title, description string, location cultural.Location, workingHours string, price cultural.Price, isAccessible bool, organizerID int, image string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTouristAttraction", ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTouristAttraction indicates an expected call of CreateTouristAttraction.
func (mr *MockCulturalServiceMockRecorder) CreateTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTouristAttraction", reflect.TypeOf((*MockCulturalService)(nil).CreateTouristAttraction), ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)
}

// DeleteEventByID mocks base method.
func (m *MockCulturalService) DeleteEventByID(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventByID indicates an expected call of DeleteEventByID.
func (mr *MockCulturalServiceMockRecorder) DeleteEventByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventByID", reflect.TypeOf((*MockCulturalService)(nil).DeleteEventByID), ctx, id)
}

// DeleteTouristAttractionByID mocks base method.
func (m *MockCulturalService) DeleteTouristAttractionByID(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTouristAttractionByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTouristAttractionByID indicates an expected call of DeleteTouristAttractionByID.
func (mr *MockCulturalServiceMockRecorder) DeleteTouristAttractionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTouristAttractionByID", reflect.TypeOf((*MockCulturalService)(nil).DeleteTouristAttractionByID), ctx, id)
}

// GetAllEvents mocks base method.
func (m *MockCulturalService) GetAllEvents(ctx context.Context) ([]cultural.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents", ctx)
	ret0, _ := ret[0].([]cultural.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockCulturalServiceMockRecorder) GetAllEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockCulturalService)(nil).GetAllEvents), ctx)
}

// GetAllTouristAttractions mocks base method.
func (m *MockCulturalService) GetAllTouristAttractions(ctx context.Context) ([]cultural.TouristAttraction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTouristAttractions", ctx)
	ret0, _ := ret[0].([]cultural.TouristAttraction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTouristAttractions indicates an expected call of GetAllTouristAttractions.
func (mr *MockCulturalServiceMockRecorder) GetAllTouristAttractions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTouristAttractions", reflect.TypeOf((*MockCulturalService)(nil).GetAllTouristAttractions), ctx)
}

// GetEventByID mocks base method.
func (m *MockCulturalService) GetEventByID(ctx context.Context, id int) (cultural.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventByID", ctx, id)
	ret0, _ := ret[0].(cultural.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventByID indicates an expected call of GetEventByID.
func (mr *MockCulturalServiceMockRecorder) GetEventByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventByID", reflect.TypeOf((*MockCulturalService)(nil).GetEventByID), ctx, id)
}

// GetEventsIDsByOrganizer mocks base method.
func (m *MockCulturalService) GetEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsIDsByOrganizer", ctx, organizerID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsIDsByOrganizer indicates an expected call of GetEventsIDsByOrganizer.
func (mr *MockCulturalServiceMockRecorder) GetEventsIDsByOrganizer(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsIDsByOrganizer", reflect.TypeOf((*MockCulturalService)(nil).GetEventsIDsByOrganizer), ctx, organizerID)
}

// GetTouristAttractionByID mocks base method.
func (m *MockCulturalService) GetTouristAttractionByID(ctx context.Context, id int) (cultural.TouristAttraction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTouristAttractionByID", ctx, id)
	ret0, _ := ret[0].(cultural.TouristAttraction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTouristAttractionByID indicates an expected call of GetTouristAttractionByID.
func (mr *MockCulturalServiceMockRecorder) GetTouristAttractionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTouristAttractionByID", reflect.TypeOf((*MockCulturalService)(nil).GetTouristAttractionByID), ctx, id)
}

// GetTouristAttractionsIDsByOrganizer mocks base method.
func (m *MockCulturalService) GetTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTouristAttractionsIDsByOrganizer", ctx, organizerID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTouristAttractionsIDsByOrganizer indicates an expected call of GetTouristAttractionsIDsByOrganizer.
func (mr *MockCulturalServiceMockRecorder) GetTouristAttractionsIDsByOrganizer(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTouristAttractionsIDsByOrganizer", reflect.TypeOf((*MockCulturalService)(nil).GetTouristAttractionsIDsByOrganizer), ctx, organizerID)
}

// UpdateEventByID mocks base method.
func (m *MockCulturalService) UpdateEventByID(ctx context.Context, id int, title, description string, location cultural.Location, startDate, finishDate, durationHours string, price cultural.Price, isAccessible bool, organizerID int, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventByID", ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventByID indicates an expected call of UpdateEventByID.
func (mr *MockCulturalServiceMockRecorder) UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventByID", reflect.TypeOf((*MockCulturalService)(nil).UpdateEventByID), ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
}

// UpdateTouristAttractionByID mocks base method.
func (m *MockCulturalService) UpdateTouristAttractionByID(ctx context.Context, id int, title, description string, location cultural.Location, workingHours string, price cultural.Price, isAccessible bool, organizerID int, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTouristAttractionByID", ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTouristAttractionByID indicates an expected call of UpdateTouristAttractionByID.
func (mr *MockCulturalServiceMockRecorder) UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTouristAttractionByID", reflect.TypeOf((*MockCulturalService)(nil).UpdateTouristAttractionByID), ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)
}
