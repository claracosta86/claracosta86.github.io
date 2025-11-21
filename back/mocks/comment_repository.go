package mocks

import (
	context "context"
	reflect "reflect"

	comment "poc2/back/domain/comment"

	gomock "go.uber.org/mock/gomock"
)

// MockCommentRepository is a mock of Repository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// FindCommentsByCultural mocks base method.
func (m *MockCommentRepository) FindCommentsByCultural(ctx context.Context, culturalID int, culturalType string) (comment.Comments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCommentsByCultural", ctx, culturalID, culturalType)
	ret0, _ := ret[0].(comment.Comments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCommentsByCultural indicates an expected call of FindCommentsByCultural.
func (mr *MockCommentRepositoryMockRecorder) FindCommentsByCultural(ctx, culturalID, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCommentsByCultural", reflect.TypeOf((*MockCommentRepository)(nil).FindCommentsByCultural), ctx, culturalID, culturalType)
}

// SaveComment mocks base method.
func (m *MockCommentRepository) SaveComment(ctx context.Context, culturalID int, culturalType string, userID int, commentArg comment.CommentContent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveComment", ctx, culturalID, culturalType, userID, commentArg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveComment indicates an expected call of SaveComment.
func (mr *MockCommentRepositoryMockRecorder) SaveComment(ctx, culturalID, culturalType, userID, commentArg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveComment", reflect.TypeOf((*MockCommentRepository)(nil).SaveComment), ctx, culturalID, culturalType, userID, commentArg)
}
