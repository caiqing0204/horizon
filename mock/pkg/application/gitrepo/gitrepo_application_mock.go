// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/application/gitrepo/gitrepo_application.go

// Package mock_gitrepo is a generated GoMock package.
package mock_gitrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockApplicationGitRepo is a mock of ApplicationGitRepo interface.
type MockApplicationGitRepo struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationGitRepoMockRecorder
}

// MockApplicationGitRepoMockRecorder is the mock recorder for MockApplicationGitRepo.
type MockApplicationGitRepoMockRecorder struct {
	mock *MockApplicationGitRepo
}

// NewMockApplicationGitRepo creates a new mock instance.
func NewMockApplicationGitRepo(ctrl *gomock.Controller) *MockApplicationGitRepo {
	mock := &MockApplicationGitRepo{ctrl: ctrl}
	mock.recorder = &MockApplicationGitRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationGitRepo) EXPECT() *MockApplicationGitRepoMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method.
func (m *MockApplicationGitRepo) CreateApplication(ctx context.Context, application string, pipelineJSONBlob, applicationJSONBlob map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, application, pipelineJSONBlob, applicationJSONBlob)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationGitRepoMockRecorder) CreateApplication(ctx, application, pipelineJSONBlob, applicationJSONBlob interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationGitRepo)(nil).CreateApplication), ctx, application, pipelineJSONBlob, applicationJSONBlob)
}

// DeleteApplication mocks base method.
func (m *MockApplicationGitRepo) DeleteApplication(ctx context.Context, application string, applicationID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", ctx, application, applicationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockApplicationGitRepoMockRecorder) DeleteApplication(ctx, application, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockApplicationGitRepo)(nil).DeleteApplication), ctx, application, applicationID)
}

// GetApplication mocks base method.
func (m *MockApplicationGitRepo) GetApplication(ctx context.Context, application string) (map[string]interface{}, map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", ctx, application)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockApplicationGitRepoMockRecorder) GetApplication(ctx, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockApplicationGitRepo)(nil).GetApplication), ctx, application)
}

// GetApplicationEnvTemplate mocks base method.
func (m *MockApplicationGitRepo) GetApplicationEnvTemplate(ctx context.Context, application, env string) (map[string]interface{}, map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationEnvTemplate", ctx, application, env)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApplicationEnvTemplate indicates an expected call of GetApplicationEnvTemplate.
func (mr *MockApplicationGitRepoMockRecorder) GetApplicationEnvTemplate(ctx, application, env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationEnvTemplate", reflect.TypeOf((*MockApplicationGitRepo)(nil).GetApplicationEnvTemplate), ctx, application, env)
}

// HardDeleteApplication mocks base method.
func (m *MockApplicationGitRepo) HardDeleteApplication(ctx context.Context, application string, applicationID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDeleteApplication", ctx, application, applicationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDeleteApplication indicates an expected call of HardDeleteApplication.
func (mr *MockApplicationGitRepoMockRecorder) HardDeleteApplication(ctx, application, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDeleteApplication", reflect.TypeOf((*MockApplicationGitRepo)(nil).HardDeleteApplication), ctx, application, applicationID)
}

// UpdateApplication mocks base method.
func (m *MockApplicationGitRepo) UpdateApplication(ctx context.Context, application string, pipelineJSONBlob, applicationJSONBlob map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", ctx, application, pipelineJSONBlob, applicationJSONBlob)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockApplicationGitRepoMockRecorder) UpdateApplication(ctx, application, pipelineJSONBlob, applicationJSONBlob interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockApplicationGitRepo)(nil).UpdateApplication), ctx, application, pipelineJSONBlob, applicationJSONBlob)
}

// UpdateApplicationEnvTemplate mocks base method.
func (m *MockApplicationGitRepo) UpdateApplicationEnvTemplate(ctx context.Context, application, env string, pipelineJSONBlob, applicationJSONBlob map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplicationEnvTemplate", ctx, application, env, pipelineJSONBlob, applicationJSONBlob)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApplicationEnvTemplate indicates an expected call of UpdateApplicationEnvTemplate.
func (mr *MockApplicationGitRepoMockRecorder) UpdateApplicationEnvTemplate(ctx, application, env, pipelineJSONBlob, applicationJSONBlob interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplicationEnvTemplate", reflect.TypeOf((*MockApplicationGitRepo)(nil).UpdateApplicationEnvTemplate), ctx, application, env, pipelineJSONBlob, applicationJSONBlob)
}
