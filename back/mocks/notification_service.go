package mocks

import (
	context "context"
	reflect "reflect"

	notification "poc2/back/domain/notification"

	gomock "go.uber.org/mock/gomock"
)

// MockNotificationService is a mock of Service interface.
type MockNotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationServiceMockRecorder
}

// MockNotificationServiceMockRecorder is the mock recorder for MockNotificationService.
type MockNotificationServiceMockRecorder struct {
	mock *MockNotificationService
}

// NewMockNotificationService creates a new mock instance.
func NewMockNotificationService(ctrl *gomock.Controller) *MockNotificationService {
	mock := &MockNotificationService{ctrl: ctrl}
	mock.recorder = &MockNotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationService) EXPECT() *MockNotificationServiceMockRecorder {
	return m.recorder
}

// GetNotifications mocks base method.
func (m *MockNotificationService) GetNotifications(ctx context.Context, userID int) ([]notification.NotificationCulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", ctx, userID)
	ret0, _ := ret[0].([]notification.NotificationCulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockNotificationServiceMockRecorder) GetNotifications(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockNotificationService)(nil).GetNotifications), ctx, userID)
}

// MarkNotificationsAsSeen mocks base method.
func (m *MockNotificationService) MarkNotificationsAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkNotificationsAsSeen", ctx, userID, notificationIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkNotificationsAsSeen indicates an expected call of MarkNotificationsAsSeen.
func (mr *MockNotificationServiceMockRecorder) MarkNotificationsAsSeen(ctx, userID, notificationIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkNotificationsAsSeen", reflect.TypeOf((*MockNotificationService)(nil).MarkNotificationsAsSeen), ctx, userID, notificationIDs)
}
