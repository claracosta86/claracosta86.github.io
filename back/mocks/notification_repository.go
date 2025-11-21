package mocks

import (
	context "context"
	reflect "reflect"

	notification "poc2/back/domain/notification"

	gomock "go.uber.org/mock/gomock"
)

// MockNotificationRepository is a mock of Repository interface.
type MockNotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRepositoryMockRecorder
}

// MockNotificationRepositoryMockRecorder is the mock recorder for MockNotificationRepository.
type MockNotificationRepositoryMockRecorder struct {
	mock *MockNotificationRepository
}

// NewMockNotificationRepository creates a new mock instance.
func NewMockNotificationRepository(ctrl *gomock.Controller) *MockNotificationRepository {
	mock := &MockNotificationRepository{ctrl: ctrl}
	mock.recorder = &MockNotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRepository) EXPECT() *MockNotificationRepositoryMockRecorder {
	return m.recorder
}

// FindByUserID mocks base method.
func (m *MockNotificationRepository) FindByUserID(ctx context.Context, userID int) ([]notification.NotificationCulturalList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].([]notification.NotificationCulturalList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockNotificationRepositoryMockRecorder) FindByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockNotificationRepository)(nil).FindByUserID), ctx, userID)
}

// MarkAsSeen mocks base method.
func (m *MockNotificationRepository) MarkAsSeen(ctx context.Context, userID int, notificationIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsSeen", ctx, userID, notificationIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsSeen indicates an expected call of MarkAsSeen.
func (mr *MockNotificationRepositoryMockRecorder) MarkAsSeen(ctx, userID, notificationIDs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsSeen", reflect.TypeOf((*MockNotificationRepository)(nil).MarkAsSeen), ctx, userID, notificationIDs)
}
