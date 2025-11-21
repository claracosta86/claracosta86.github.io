package mocks

import (
	context "context"
	reflect "reflect"

	comment "poc2/back/domain/comment"

	gomock "go.uber.org/mock/gomock"
)

// MockCommentService is a mock of Service interface.
type MockCommentService struct {
	ctrl     *gomock.Controller
	recorder *MockCommentServiceMockRecorder
}

// MockCommentServiceMockRecorder is the mock recorder for MockCommentService.
type MockCommentServiceMockRecorder struct {
	mock *MockCommentService
}

// NewMockCommentService creates a new mock instance.
func NewMockCommentService(ctrl *gomock.Controller) *MockCommentService {
	mock := &MockCommentService{ctrl: ctrl}
	mock.recorder = &MockCommentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentService) EXPECT() *MockCommentServiceMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockCommentService) CreateComment(ctx context.Context, culturalID int, culturalType string, userID int, commentVal comment.CommentContent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, culturalID, culturalType, userID, commentVal)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockCommentServiceMockRecorder) CreateComment(ctx, culturalID, culturalType, userID, commentVal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentService)(nil).CreateComment), ctx, culturalID, culturalType, userID, commentVal)
}

// GetComments mocks base method.
func (m *MockCommentService) GetComments(ctx context.Context, culturalID int, culturalType string) (comment.Comments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", ctx, culturalID, culturalType)
	ret0, _ := ret[0].(comment.Comments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockCommentServiceMockRecorder) GetComments(ctx, culturalID, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockCommentService)(nil).GetComments), ctx, culturalID, culturalType)
}
