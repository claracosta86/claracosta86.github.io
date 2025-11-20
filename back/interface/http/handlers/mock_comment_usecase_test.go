package handlers_test

import (
	"context"
	"reflect"

	model "poc2/back/interface/model"

	gomock "go.uber.org/mock/gomock"
)

// MockCommentUseCase is a mock of UseCase interface.
type MockCommentUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCommentUseCaseMockRecorder
}

// MockCommentUseCaseMockRecorder is the mock recorder for MockCommentUseCase.
type MockCommentUseCaseMockRecorder struct {
	mock *MockCommentUseCase
}

// NewMockCommentUseCase creates a new mock instance.
func NewMockCommentUseCase(ctrl *gomock.Controller) *MockCommentUseCase {
	mock := &MockCommentUseCase{ctrl: ctrl}
	mock.recorder = &MockCommentUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentUseCase) EXPECT() *MockCommentUseCaseMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentUseCase) CreateComment(ctx context.Context, data model.CreateCommentRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentUseCaseMockRecorder) CreateComment(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentUseCase)(nil).CreateComment), ctx, data)
}

// GetComments mocks base method.
func (m *MockCommentUseCase) GetComments(ctx context.Context, culturalID int, culturalType string) (model.GetCommentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", ctx, culturalID, culturalType)
	ret0, _ := ret[0].(model.GetCommentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockCommentUseCaseMockRecorder) GetComments(ctx, culturalID, culturalType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockCommentUseCase)(nil).GetComments), ctx, culturalID, culturalType)
}
