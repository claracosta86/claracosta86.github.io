package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	user "poc2/back/domain/user"
)

// MockUserRepository is a mock of Repository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddFavorite mocks base method.
func (m *MockUserRepository) AddFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFavorite", ctx, userID, culturalType, culturalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFavorite indicates an expected call of AddFavorite.
func (mr *MockUserRepositoryMockRecorder) AddFavorite(ctx, userID, culturalType, culturalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFavorite", reflect.TypeOf((*MockUserRepository)(nil).AddFavorite), ctx, userID, culturalType, culturalID)
}

// CheckPassword mocks base method.
func (m *MockUserRepository) CheckPassword(ctx context.Context, userID int, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", ctx, userID, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockUserRepositoryMockRecorder) CheckPassword(ctx, userID, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockUserRepository)(nil).CheckPassword), ctx, userID, password)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), ctx, id)
}

// DeleteUserFavorites mocks base method.
func (m *MockUserRepository) DeleteUserFavorites(ctx context.Context, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserFavorites", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserFavorites indicates an expected call of DeleteUserFavorites.
func (mr *MockUserRepositoryMockRecorder) DeleteUserFavorites(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserFavorites", reflect.TypeOf((*MockUserRepository)(nil).DeleteUserFavorites), ctx, userID)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), ctx, email)
}

// FindByID mocks base method.
func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), ctx, id)
}

// GetCulturaisByOrganizerID mocks base method.
func (m *MockUserRepository) GetCulturaisByOrganizerID(ctx context.Context, organizerID int) ([]user.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCulturaisByOrganizerID", ctx, organizerID)
	ret0, _ := ret[0].([]user.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCulturaisByOrganizerID indicates an expected call of GetCulturaisByOrganizerID.
func (mr *MockUserRepositoryMockRecorder) GetCulturaisByOrganizerID(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCulturaisByOrganizerID", reflect.TypeOf((*MockUserRepository)(nil).GetCulturaisByOrganizerID), ctx, organizerID)
}

// GetFavoritesByUserID mocks base method.
func (m *MockUserRepository) GetFavoritesByUserID(ctx context.Context, userID int) ([]user.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavoritesByUserID", ctx, userID)
	ret0, _ := ret[0].([]user.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavoritesByUserID indicates an expected call of GetFavoritesByUserID.
func (mr *MockUserRepositoryMockRecorder) GetFavoritesByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavoritesByUserID", reflect.TypeOf((*MockUserRepository)(nil).GetFavoritesByUserID), ctx, userID)
}

// RemoveEvent mocks base method.
func (m *MockUserRepository) RemoveEvent(ctx context.Context, eventIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEvent", ctx, eventIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEvent indicates an expected call of RemoveEvent.
func (mr *MockUserRepositoryMockRecorder) RemoveEvent(ctx, eventIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEvent", reflect.TypeOf((*MockUserRepository)(nil).RemoveEvent), ctx, eventIDs)
}

// RemoveFavorite mocks base method.
func (m *MockUserRepository) RemoveFavorite(ctx context.Context, userID int, culturalType string, culturalID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFavorite", ctx, userID, culturalType, culturalID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFavorite indicates an expected call of RemoveFavorite.
func (mr *MockUserRepositoryMockRecorder) RemoveFavorite(ctx, userID, culturalType, culturalID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFavorite", reflect.TypeOf((*MockUserRepository)(nil).RemoveFavorite), ctx, userID, culturalType, culturalID)
}

// RemoveTouristAttraction mocks base method.
func (m *MockUserRepository) RemoveTouristAttraction(ctx context.Context, touristAttractionIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTouristAttraction", ctx, touristAttractionIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTouristAttraction indicates an expected call of RemoveTouristAttraction.
func (mr *MockUserRepositoryMockRecorder) RemoveTouristAttraction(ctx, touristAttractionIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTouristAttraction", reflect.TypeOf((*MockUserRepository)(nil).RemoveTouristAttraction), ctx, touristAttractionIDs)
}

// Save mocks base method.
func (m *MockUserRepository) Save(ctx context.Context, user *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), ctx, user)
}

// Update mocks base method.
func (m *MockUserRepository) Update(ctx context.Context, user *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), ctx, user)
}

// UpdateLastSeenFavorite mocks base method.
func (m *MockUserRepository) UpdateLastSeenFavorite(ctx context.Context, userID, culturalID int, culturalType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastSeenFavorite", ctx, userID, culturalID, culturalType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastSeenFavorite indicates an expected call of UpdateLastSeenFavorite.
func (mr *MockUserRepositoryMockRecorder) UpdateLastSeenFavorite(ctx, userID, culturalID, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastSeenFavorite", reflect.TypeOf((*MockUserRepository)(nil).UpdateLastSeenFavorite), ctx, userID, culturalID, culturalType)
}

// UpdatePassword mocks base method.
func (m *MockUserRepository) UpdatePassword(ctx context.Context, userID int, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, userID, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(ctx, userID, newPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), ctx, userID, newPassword)
}
