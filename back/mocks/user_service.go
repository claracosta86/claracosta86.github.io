package mocks

import (
	context "context"
	reflect "reflect"

	user "poc2/back/domain/user"

	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of Service interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockUserService) AuthenticateUser(ctx context.Context, email, password string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", ctx, email, password)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockUserServiceMockRecorder) AuthenticateUser(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockUserService)(nil).AuthenticateUser), ctx, email, password)
}

// ChangeUserPassword mocks base method.
func (m *MockUserService) ChangeUserPassword(ctx context.Context, id int, currentPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", ctx, id, currentPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockUserServiceMockRecorder) ChangeUserPassword(ctx, id, currentPassword, newPassword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserService)(nil).ChangeUserPassword), ctx, id, currentPassword, newPassword)
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), ctx, id)
}

// GetCulturaisByOrganizerID mocks base method.
func (m *MockUserService) GetCulturaisByOrganizerID(ctx context.Context, organizerID int) ([]user.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCulturaisByOrganizerID", ctx, organizerID)
	ret0, _ := ret[0].([]user.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCulturaisByOrganizerID indicates an expected call of GetCulturaisByOrganizerID.
func (mr *MockUserServiceMockRecorder) GetCulturaisByOrganizerID(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCulturaisByOrganizerID", reflect.TypeOf((*MockUserService)(nil).GetCulturaisByOrganizerID), ctx, organizerID)
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(ctx context.Context, id int) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), ctx, id)
}

// GetUserFavorites mocks base method.
func (m *MockUserService) GetUserFavorites(ctx context.Context, userID int) ([]user.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavorites", ctx, userID)
	ret0, _ := ret[0].([]user.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFavorites indicates an expected call of GetUserFavorites.
func (mr *MockUserServiceMockRecorder) GetUserFavorites(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavorites", reflect.TypeOf((*MockUserService)(nil).GetUserFavorites), ctx, userID)
}

// RegisterUser mocks base method.
func (m *MockUserService) RegisterUser(ctx context.Context, name, email, document, companyName, password string, userType user.UserType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, name, email, document, companyName, password, userType)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserServiceMockRecorder) RegisterUser(ctx, name, email, document, companyName, password, userType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserService)(nil).RegisterUser), ctx, name, email, document, companyName, password, userType)
}

// RemoveEventFromAllUsers mocks base method.
func (m *MockUserService) RemoveEventFromAllUsers(ctx context.Context, eventIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEventFromAllUsers", ctx, eventIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEventFromAllUsers indicates an expected call of RemoveEventFromAllUsers.
func (mr *MockUserServiceMockRecorder) RemoveEventFromAllUsers(ctx, eventIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEventFromAllUsers", reflect.TypeOf((*MockUserService)(nil).RemoveEventFromAllUsers), ctx, eventIDs)
}

// RemoveTouristAttractionFromAllUsers mocks base method.
func (m *MockUserService) RemoveTouristAttractionFromAllUsers(ctx context.Context, attractionIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTouristAttractionFromAllUsers", ctx, attractionIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTouristAttractionFromAllUsers indicates an expected call of RemoveTouristAttractionFromAllUsers.
func (mr *MockUserServiceMockRecorder) RemoveTouristAttractionFromAllUsers(ctx, attractionIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTouristAttractionFromAllUsers", reflect.TypeOf((*MockUserService)(nil).RemoveTouristAttractionFromAllUsers), ctx, attractionIDs)
}

// UpdateFavorites mocks base method.
func (m *MockUserService) UpdateFavorites(ctx context.Context, userID int, culturalType string, culturalID int, isFavorite bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFavorites", ctx, userID, culturalType, culturalID, isFavorite)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFavorites indicates an expected call of UpdateFavorites.
func (mr *MockUserServiceMockRecorder) UpdateFavorites(ctx, userID, culturalType, culturalID, isFavorite any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFavorites", reflect.TypeOf((*MockUserService)(nil).UpdateFavorites), ctx, userID, culturalType, culturalID, isFavorite)
}

// UpdateLastSeenFavorite mocks base method.
func (m *MockUserService) UpdateLastSeenFavorite(ctx context.Context, userID, culturalID int, culturalType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastSeenFavorite", ctx, userID, culturalID, culturalType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastSeenFavorite indicates an expected call of UpdateLastSeenFavorite.
func (mr *MockUserServiceMockRecorder) UpdateLastSeenFavorite(ctx, userID, culturalID, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastSeenFavorite", reflect.TypeOf((*MockUserService)(nil).UpdateLastSeenFavorite), ctx, userID, culturalID, culturalType)
}

// UpdateUserProfile mocks base method.
func (m *MockUserService) UpdateUserProfile(ctx context.Context, id int, name, email, companyName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", ctx, id, name, email, companyName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockUserServiceMockRecorder) UpdateUserProfile(ctx, id, name, email, companyName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockUserService)(nil).UpdateUserProfile), ctx, id, name, email, companyName)
}
