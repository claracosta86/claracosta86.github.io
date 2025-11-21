package mocks

import (
	context "context"
	reflect "reflect"

	cultural "poc2/back/domain/cultural"

	gomock "go.uber.org/mock/gomock"
)

// MockCulturalRepository is a mock of Repository interface.
type MockCulturalRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCulturalRepositoryMockRecorder
}

// MockCulturalRepositoryMockRecorder is the mock recorder for MockCulturalRepository.
type MockCulturalRepositoryMockRecorder struct {
	mock *MockCulturalRepository
}

// NewMockCulturalRepository creates a new mock instance.
func NewMockCulturalRepository(ctrl *gomock.Controller) *MockCulturalRepository {
	mock := &MockCulturalRepository{ctrl: ctrl}
	mock.recorder = &MockCulturalRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCulturalRepository) EXPECT() *MockCulturalRepositoryMockRecorder {
	return m.recorder
}

// DeleteEventByID mocks base method.
func (m *MockCulturalRepository) DeleteEventByID(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEventByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEventByID indicates an expected call of DeleteEventByID.
func (mr *MockCulturalRepositoryMockRecorder) DeleteEventByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEventByID", reflect.TypeOf((*MockCulturalRepository)(nil).DeleteEventByID), ctx, id)
}

// DeleteTouristAttractionByID mocks base method.
func (m *MockCulturalRepository) DeleteTouristAttractionByID(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTouristAttractionByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTouristAttractionByID indicates an expected call of DeleteTouristAttractionByID.
func (mr *MockCulturalRepositoryMockRecorder) DeleteTouristAttractionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTouristAttractionByID", reflect.TypeOf((*MockCulturalRepository)(nil).DeleteTouristAttractionByID), ctx, id)
}

// FindAllEvents mocks base method.
func (m *MockCulturalRepository) FindAllEvents(ctx context.Context) ([]cultural.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllEvents", ctx)
	ret0, _ := ret[0].([]cultural.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllEvents indicates an expected call of FindAllEvents.
func (mr *MockCulturalRepositoryMockRecorder) FindAllEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllEvents", reflect.TypeOf((*MockCulturalRepository)(nil).FindAllEvents), ctx)
}

// FindAllTouristAttractions mocks base method.
func (m *MockCulturalRepository) FindAllTouristAttractions(ctx context.Context) ([]cultural.TouristAttraction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTouristAttractions", ctx)
	ret0, _ := ret[0].([]cultural.TouristAttraction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTouristAttractions indicates an expected call of FindAllTouristAttractions.
func (mr *MockCulturalRepositoryMockRecorder) FindAllTouristAttractions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTouristAttractions", reflect.TypeOf((*MockCulturalRepository)(nil).FindAllTouristAttractions), ctx)
}

// FindEventByID mocks base method.
func (m *MockCulturalRepository) FindEventByID(ctx context.Context, id int) (cultural.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventByID", ctx, id)
	ret0, _ := ret[0].(cultural.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventByID indicates an expected call of FindEventByID.
func (mr *MockCulturalRepositoryMockRecorder) FindEventByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventByID", reflect.TypeOf((*MockCulturalRepository)(nil).FindEventByID), ctx, id)
}

// FindEventsIDsByOrganizer mocks base method.
func (m *MockCulturalRepository) FindEventsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventsIDsByOrganizer", ctx, organizerID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventsIDsByOrganizer indicates an expected call of FindEventsIDsByOrganizer.
func (mr *MockCulturalRepositoryMockRecorder) FindEventsIDsByOrganizer(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventsIDsByOrganizer", reflect.TypeOf((*MockCulturalRepository)(nil).FindEventsIDsByOrganizer), ctx, organizerID)
}

// FindTouristAttractionByID mocks base method.
func (m *MockCulturalRepository) FindTouristAttractionByID(ctx context.Context, id int) (cultural.TouristAttraction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTouristAttractionByID", ctx, id)
	ret0, _ := ret[0].(cultural.TouristAttraction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTouristAttractionByID indicates an expected call of FindTouristAttractionByID.
func (mr *MockCulturalRepositoryMockRecorder) FindTouristAttractionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTouristAttractionByID", reflect.TypeOf((*MockCulturalRepository)(nil).FindTouristAttractionByID), ctx, id)
}

// FindTouristAttractionsIDsByOrganizer mocks base method.
func (m *MockCulturalRepository) FindTouristAttractionsIDsByOrganizer(ctx context.Context, organizerID int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTouristAttractionsIDsByOrganizer", ctx, organizerID)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTouristAttractionsIDsByOrganizer indicates an expected call of FindTouristAttractionsIDsByOrganizer.
func (mr *MockCulturalRepositoryMockRecorder) FindTouristAttractionsIDsByOrganizer(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTouristAttractionsIDsByOrganizer", reflect.TypeOf((*MockCulturalRepository)(nil).FindTouristAttractionsIDsByOrganizer), ctx, organizerID)
}

// SaveEvent mocks base method.
func (m *MockCulturalRepository) SaveEvent(ctx context.Context, title, description string, location cultural.Location, startDate, finishDate, durationHours string, price cultural.Price, isAccessible bool, organizerID int, image string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEvent", ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveEvent indicates an expected call of SaveEvent.
func (mr *MockCulturalRepositoryMockRecorder) SaveEvent(ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEvent", reflect.TypeOf((*MockCulturalRepository)(nil).SaveEvent), ctx, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
}

// SaveTouristAttraction mocks base method.
func (m *MockCulturalRepository) SaveTouristAttraction(ctx context.Context, title, description string, location cultural.Location, workingHours string, price cultural.Price, isAccessible bool, organizerID int, image string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTouristAttraction", ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTouristAttraction indicates an expected call of SaveTouristAttraction.
func (mr *MockCulturalRepositoryMockRecorder) SaveTouristAttraction(ctx, title, description, location, workingHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTouristAttraction", reflect.TypeOf((*MockCulturalRepository)(nil).SaveTouristAttraction), ctx, title, description, location, workingHours, price, isAccessible, organizerID, image)
}

// UpdateEventByID mocks base method.
func (m *MockCulturalRepository) UpdateEventByID(ctx context.Context, id int, title, description string, location cultural.Location, startDate, finishDate, durationHours string, price cultural.Price, isAccessible bool, organizerID int, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventByID", ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventByID indicates an expected call of UpdateEventByID.
func (mr *MockCulturalRepositoryMockRecorder) UpdateEventByID(ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventByID", reflect.TypeOf((*MockCulturalRepository)(nil).UpdateEventByID), ctx, id, title, description, location, startDate, finishDate, durationHours, price, isAccessible, organizerID, image)
}

// UpdateTouristAttractionByID mocks base method.
func (m *MockCulturalRepository) UpdateTouristAttractionByID(ctx context.Context, id int, title, description string, location cultural.Location, workingHours string, price cultural.Price, isAccessible bool, organizerID int, image string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTouristAttractionByID", ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTouristAttractionByID indicates an expected call of UpdateTouristAttractionByID.
func (mr *MockCulturalRepositoryMockRecorder) UpdateTouristAttractionByID(ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTouristAttractionByID", reflect.TypeOf((*MockCulturalRepository)(nil).UpdateTouristAttractionByID), ctx, id, title, description, location, workingHours, price, isAccessible, organizerID, image)
}
