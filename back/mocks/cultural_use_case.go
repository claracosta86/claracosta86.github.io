package mocks

import (
	"context"
	"reflect"

	gomock "go.uber.org/mock/gomock"

	model "poc2/back/interface/model"
)

// MockCommentUseCase is a mock of UseCase interface.
type MockCulturalUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCulturalUseCaseMockRecorder
}

// MockCommentUseCaseMockRecorder is the mock recorder for MockCommentUseCase.
type MockCulturalUseCaseMockRecorder struct {
	mock *MockCulturalUseCase
}

// NewMockCommentUseCase creates a new mock instance.
func NewMockCulturalUseCase(ctrl *gomock.Controller) *MockCulturalUseCase {
	mock := &MockCulturalUseCase{ctrl: ctrl}
	mock.recorder = &MockCulturalUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCulturalUseCase) EXPECT() *MockCulturalUseCaseMockRecorder {
	return m.recorder
}

// CreateCultural mocks base method.
func (m *MockCulturalUseCase) CreateCultural(ctx context.Context, data model.CreateCulturalRequest) (model.CreateCulturalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCultural", ctx, data)
	ret0, _ := ret[0].(model.CreateCulturalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCultural indicates an expected call of CreateCultural.
func (mr *MockCulturalUseCaseMockRecorder) CreateCultural(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCultural", reflect.TypeOf((*MockCulturalUseCase)(nil).CreateCultural), ctx, data)
}

// GetCultural mocks base method.
func (m *MockCulturalUseCase) GetCultural(ctx context.Context, id int, culturalType string) (model.GetCulturalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCultural", ctx, id, culturalType)
	ret0, _ := ret[0].(model.GetCulturalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCultural indicates an expected call of GetCultural.
func (mr *MockCulturalUseCaseMockRecorder) GetCultural(ctx, id, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCultural", reflect.TypeOf((*MockCulturalUseCase)(nil).GetCultural), ctx, id, culturalType)
}

// GetAllCulturais mocks base method.
func (m *MockCulturalUseCase) GetAllCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCulturais", ctx)
	ret0, _ := ret[0].(model.GetAllCulturaisResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCulturais indicates an expected call of GetAllCulturais.
func (mr *MockCulturalUseCaseMockRecorder) GetAllCulturais(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCulturais", reflect.TypeOf((*MockCulturalUseCase)(nil).GetAllCulturais), ctx)
}

// GetHomeCulturais mocks base method.
func (m *MockCulturalUseCase) GetHomeCulturais(ctx context.Context) (model.GetAllCulturaisResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHomeCulturais", ctx)
	ret0, _ := ret[0].(model.GetAllCulturaisResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHomeCulturais indicates an expected call of GetHomeCulturais.
func (mr *MockCulturalUseCaseMockRecorder) GetHomeCulturais(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHomeCulturais", reflect.TypeOf((*MockCulturalUseCase)(nil).GetHomeCulturais), ctx)
}

// UpdateCultural mocks base method.
func (m *MockCulturalUseCase) UpdateCultural(ctx context.Context, data model.UpdateCulturalRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCultural", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCultural indicates an expected call of UpdateCultural.
func (mr *MockCulturalUseCaseMockRecorder) UpdateCultural(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCultural", reflect.TypeOf((*MockCulturalUseCase)(nil).UpdateCultural), ctx, data)
}

// DeleteCultural mocks base method.
func (m *MockCulturalUseCase) DeleteCultural(ctx context.Context, id int, culturalType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCultural", ctx, id, culturalType)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCultural indicates an expected call of DeleteCultural.
func (mr *MockCulturalUseCaseMockRecorder) DeleteCultural(ctx, id, culturalType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCultural", reflect.TypeOf((*MockCulturalUseCase)(nil).DeleteCultural), ctx, id, culturalType)
}
