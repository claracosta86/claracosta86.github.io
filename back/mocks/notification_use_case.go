package mocks

import (
	"context"
	"reflect"

	gomock "go.uber.org/mock/gomock"

	model "poc2/back/interface/model"
)

// MockNotificationUseCase is a mock of UseCase interface.
type MockNotificationUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationUseCaseMockRecorder
}

// MockNotificationUseCaseMockRecorder is the mock recorder for MockNotificationUseCase.
type MockNotificationUseCaseMockRecorder struct {
	mock *MockNotificationUseCase
}

// NewMockNotificationUseCase creates a new mock instance.
func NewMockNotificationUseCase(ctrl *gomock.Controller) *MockNotificationUseCase {
	mock := &MockNotificationUseCase{ctrl: ctrl}
	mock.recorder = &MockNotificationUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationUseCase) EXPECT() *MockNotificationUseCaseMockRecorder {
	return m.recorder
}

// GetNotifications mocks base method.
func (m *MockNotificationUseCase) GetNotifications(ctx context.Context, userID int) (*model.GetNotificationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", ctx, userID)
	ret0, _ := ret[0].(*model.GetNotificationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockNotificationUseCaseMockRecorder) GetNotifications(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockNotificationUseCase)(nil).GetNotifications), ctx, userID)
}

// MarkNotificationsAsSeen mocks base method.
func (m *MockNotificationUseCase) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkNotificationsAsSeen", ctx, userID, notificationIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkNotificationsAsSeen indicates an expected call of MarkNotificationsAsSeen.
func (mr *MockNotificationUseCaseMockRecorder) MarkNotificationsAsSeen(ctx, userID, notificationIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkNotificationsAsSeen", reflect.TypeOf((*MockNotificationUseCase)(nil).MarkNotificationsAsSeen), ctx, userID, notificationIDs)
}
