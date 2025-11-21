package mocks

import (
	"context"
	"reflect"

	gomock "go.uber.org/mock/gomock"

	model "poc2/back/interface/model"
)

// MockUserUseCase is a mock of UseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockUserUseCase) RegisterUser(ctx context.Context, request model.RegisterUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserUseCaseMockRecorder) RegisterUser(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserUseCase)(nil).RegisterUser), ctx, request)
}

// LoginUser mocks base method.
func (m *MockUserUseCase) LoginUser(ctx context.Context, request model.LoginUserRequest) (*model.LoginUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, request)
	ret0, _ := ret[0].(*model.LoginUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockUserUseCaseMockRecorder) LoginUser(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockUserUseCase)(nil).LoginUser), ctx, request)
}

// GetUserProfile mocks base method.
func (m *MockUserUseCase) GetUserProfile(ctx context.Context, userID int) (*model.GetUserProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", ctx, userID)
	ret0, _ := ret[0].(*model.GetUserProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockUserUseCaseMockRecorder) GetUserProfile(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockUserUseCase)(nil).GetUserProfile), ctx, userID)
}

// UpdateUserProfile mocks base method.
func (m *MockUserUseCase) UpdateUserProfile(ctx context.Context, userID int, request model.UpdateUserProfileRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockUserUseCaseMockRecorder) UpdateUserProfile(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUserProfile), ctx, userID, request)
}

// ChangePassword mocks base method.
func (m *MockUserUseCase) ChangePassword(ctx context.Context, userID int, request model.ChangePasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserUseCaseMockRecorder) ChangePassword(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserUseCase)(nil).ChangePassword), ctx, userID, request)
}

// DeleteUser mocks base method.
func (m *MockUserUseCase) DeleteUser(ctx context.Context, userID int, userType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID, userType)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserUseCaseMockRecorder) DeleteUser(ctx, userID, userType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserUseCase)(nil).DeleteUser), ctx, userID, userType)
}

// UpdateFavorites mocks base method.
func (m *MockUserUseCase) UpdateFavorites(ctx context.Context, userID int, request model.FavoriteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFavorites", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFavorites indicates an expected call of UpdateFavorites.
func (mr *MockUserUseCaseMockRecorder) UpdateFavorites(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFavorites", reflect.TypeOf((*MockUserUseCase)(nil).UpdateFavorites), ctx, userID, request)
}

// GetUserFavorites mocks base method.
func (m *MockUserUseCase) GetUserFavorites(ctx context.Context, userID int) ([]model.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavorites", ctx, userID)
	ret0, _ := ret[0].([]model.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFavorites indicates an expected call of GetUserFavorites.
func (mr *MockUserUseCaseMockRecorder) GetUserFavorites(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavorites", reflect.TypeOf((*MockUserUseCase)(nil).GetUserFavorites), ctx, userID)
}

// UpdateLastSeenFavorite mocks base method.
func (m *MockUserUseCase) UpdateLastSeenFavorite(ctx context.Context, userID int, request model.FavoriteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastSeenFavorite", ctx, userID, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastSeenFavorite indicates an expected call of UpdateLastSeenFavorite.
func (mr *MockUserUseCaseMockRecorder) UpdateLastSeenFavorite(ctx, userID, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastSeenFavorite", reflect.TypeOf((*MockUserUseCase)(nil).UpdateLastSeenFavorite), ctx, userID, request)
}

// GetOrganizerCulturais mocks base method.
func (m *MockUserUseCase) GetOrganizerCulturais(ctx context.Context, organizerID int) ([]model.CulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizerCulturais", ctx, organizerID)
	ret0, _ := ret[0].([]model.CulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizerCulturais indicates an expected call of GetOrganizerCulturais.
func (mr *MockUserUseCaseMockRecorder) GetOrganizerCulturais(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizerCulturais", reflect.TypeOf((*MockUserUseCase)(nil).GetOrganizerCulturais), ctx, organizerID)
}

// GetOrganizerInfo mocks base method.
func (m *MockUserUseCase) GetOrganizerInfo(ctx context.Context, organizerID int) (*model.GetOrganizerInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizerInfo", ctx, organizerID)
	ret0, _ := ret[0].(*model.GetOrganizerInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizerInfo indicates an expected call of GetOrganizerInfo.
func (mr *MockUserUseCaseMockRecorder) GetOrganizerInfo(ctx, organizerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizerInfo", reflect.TypeOf((*MockUserUseCase)(nil).GetOrganizerInfo), ctx, organizerID)
}
