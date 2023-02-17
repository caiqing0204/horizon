// Code generated by MockGen. DO NOT EDIT.
// Source: gitrepo_cluster.go

// Package mock_gitrepo is a generated GoMock package.
package mock_gitrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitrepo "github.com/horizoncd/horizon/pkg/cluster/gitrepo"
	models "github.com/horizoncd/horizon/pkg/tag/models"
)

// MockClusterGitRepo is a mock of ClusterGitRepo interface.
type MockClusterGitRepo struct {
	ctrl     *gomock.Controller
	recorder *MockClusterGitRepoMockRecorder
}

// MockClusterGitRepoMockRecorder is the mock recorder for MockClusterGitRepo.
type MockClusterGitRepoMockRecorder struct {
	mock *MockClusterGitRepo
}

// NewMockClusterGitRepo creates a new mock instance.
func NewMockClusterGitRepo(ctrl *gomock.Controller) *MockClusterGitRepo {
	mock := &MockClusterGitRepo{ctrl: ctrl}
	mock.recorder = &MockClusterGitRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterGitRepo) EXPECT() *MockClusterGitRepoMockRecorder {
	return m.recorder
}

// CompareConfig mocks base method.
func (m *MockClusterGitRepo) CompareConfig(ctx context.Context, application, cluster string, from, to *string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompareConfig", ctx, application, cluster, from, to)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompareConfig indicates an expected call of CompareConfig.
func (mr *MockClusterGitRepoMockRecorder) CompareConfig(ctx, application, cluster, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareConfig", reflect.TypeOf((*MockClusterGitRepo)(nil).CompareConfig), ctx, application, cluster, from, to)
}

// CreateCluster mocks base method.
func (m *MockClusterGitRepo) CreateCluster(ctx context.Context, params *gitrepo.CreateClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster.
func (mr *MockClusterGitRepoMockRecorder) CreateCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).CreateCluster), ctx, params)
}

// DefaultBranch mocks base method.
func (m *MockClusterGitRepo) DefaultBranch() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultBranch")
	ret0, _ := ret[0].(string)
	return ret0
}

// DefaultBranch indicates an expected call of DefaultBranch.
func (mr *MockClusterGitRepoMockRecorder) DefaultBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultBranch", reflect.TypeOf((*MockClusterGitRepo)(nil).DefaultBranch))
}

// DeleteCluster mocks base method.
func (m *MockClusterGitRepo) DeleteCluster(ctx context.Context, application, cluster string, clusterID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, application, cluster, clusterID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockClusterGitRepoMockRecorder) DeleteCluster(ctx, application, cluster, clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).DeleteCluster), ctx, application, cluster, clusterID)
}

// GetCluster mocks base method.
func (m *MockClusterGitRepo) GetCluster(ctx context.Context, application, cluster, templateName string) (*gitrepo.ClusterFiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, application, cluster, templateName)
	ret0, _ := ret[0].(*gitrepo.ClusterFiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockClusterGitRepoMockRecorder) GetCluster(ctx, application, cluster, templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).GetCluster), ctx, application, cluster, templateName)
}

// GetClusterTemplate mocks base method.
func (m *MockClusterGitRepo) GetClusterTemplate(ctx context.Context, application, cluster string) (*gitrepo.ClusterTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterTemplate", ctx, application, cluster)
	ret0, _ := ret[0].(*gitrepo.ClusterTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterTemplate indicates an expected call of GetClusterTemplate.
func (mr *MockClusterGitRepoMockRecorder) GetClusterTemplate(ctx, application, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterTemplate", reflect.TypeOf((*MockClusterGitRepo)(nil).GetClusterTemplate), ctx, application, cluster)
}

// GetClusterValueFiles mocks base method.
func (m *MockClusterGitRepo) GetClusterValueFiles(ctx context.Context, application, cluster string) ([]gitrepo.ClusterValueFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterValueFiles", ctx, application, cluster)
	ret0, _ := ret[0].([]gitrepo.ClusterValueFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterValueFiles indicates an expected call of GetClusterValueFiles.
func (mr *MockClusterGitRepoMockRecorder) GetClusterValueFiles(ctx, application, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterValueFiles", reflect.TypeOf((*MockClusterGitRepo)(nil).GetClusterValueFiles), ctx, application, cluster)
}

// GetConfigCommit mocks base method.
func (m *MockClusterGitRepo) GetConfigCommit(ctx context.Context, application, cluster string) (*gitrepo.ClusterCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigCommit", ctx, application, cluster)
	ret0, _ := ret[0].(*gitrepo.ClusterCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigCommit indicates an expected call of GetConfigCommit.
func (mr *MockClusterGitRepoMockRecorder) GetConfigCommit(ctx, application, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigCommit", reflect.TypeOf((*MockClusterGitRepo)(nil).GetConfigCommit), ctx, application, cluster)
}

// GetEnvValue mocks base method.
func (m *MockClusterGitRepo) GetEnvValue(ctx context.Context, application, cluster, templateName string) (*gitrepo.EnvValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvValue", ctx, application, cluster, templateName)
	ret0, _ := ret[0].(*gitrepo.EnvValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvValue indicates an expected call of GetEnvValue.
func (mr *MockClusterGitRepoMockRecorder) GetEnvValue(ctx, application, cluster, templateName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvValue", reflect.TypeOf((*MockClusterGitRepo)(nil).GetEnvValue), ctx, application, cluster, templateName)
}

// GetPipelineOutput mocks base method.
func (m *MockClusterGitRepo) GetPipelineOutput(ctx context.Context, application, cluster, template string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineOutput", ctx, application, cluster, template)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineOutput indicates an expected call of GetPipelineOutput.
func (mr *MockClusterGitRepoMockRecorder) GetPipelineOutput(ctx, application, cluster, template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineOutput", reflect.TypeOf((*MockClusterGitRepo)(nil).GetPipelineOutput), ctx, application, cluster, template)
}

// GetRepoInfo mocks base method.
func (m *MockClusterGitRepo) GetRepoInfo(ctx context.Context, application, cluster string) *gitrepo.RepoInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepoInfo", ctx, application, cluster)
	ret0, _ := ret[0].(*gitrepo.RepoInfo)
	return ret0
}

// GetRepoInfo indicates an expected call of GetRepoInfo.
func (mr *MockClusterGitRepoMockRecorder) GetRepoInfo(ctx, application, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepoInfo", reflect.TypeOf((*MockClusterGitRepo)(nil).GetRepoInfo), ctx, application, cluster)
}

// GetRestartTime mocks base method.
func (m *MockClusterGitRepo) GetRestartTime(ctx context.Context, application, cluster, template string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestartTime", ctx, application, cluster, template)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestartTime indicates an expected call of GetRestartTime.
func (mr *MockClusterGitRepoMockRecorder) GetRestartTime(ctx, application, cluster, template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestartTime", reflect.TypeOf((*MockClusterGitRepo)(nil).GetRestartTime), ctx, application, cluster, template)
}

// HardDeleteCluster mocks base method.
func (m *MockClusterGitRepo) HardDeleteCluster(ctx context.Context, application, cluster string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDeleteCluster", ctx, application, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDeleteCluster indicates an expected call of HardDeleteCluster.
func (mr *MockClusterGitRepoMockRecorder) HardDeleteCluster(ctx, application, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDeleteCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).HardDeleteCluster), ctx, application, cluster)
}

// MergeBranch mocks base method.
func (m *MockClusterGitRepo) MergeBranch(ctx context.Context, application, cluster, sourceBranch, targetBranch string, prID *uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MergeBranch", ctx, application, cluster, sourceBranch, targetBranch, prID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeBranch indicates an expected call of MergeBranch.
func (mr *MockClusterGitRepoMockRecorder) MergeBranch(ctx, application, cluster, sourceBranch, targetBranch, prID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeBranch", reflect.TypeOf((*MockClusterGitRepo)(nil).MergeBranch), ctx, application, cluster, sourceBranch, targetBranch, prID)
}

// Rollback mocks base method.
func (m *MockClusterGitRepo) Rollback(ctx context.Context, application, cluster, commit string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx, application, cluster, commit)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rollback indicates an expected call of Rollback.
func (mr *MockClusterGitRepoMockRecorder) Rollback(ctx, application, cluster, commit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockClusterGitRepo)(nil).Rollback), ctx, application, cluster, commit)
}

// UpdateCluster mocks base method.
func (m *MockClusterGitRepo) UpdateCluster(ctx context.Context, params *gitrepo.UpdateClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCluster indicates an expected call of UpdateCluster.
func (mr *MockClusterGitRepoMockRecorder) UpdateCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).UpdateCluster), ctx, params)
}

// UpdatePipelineOutput mocks base method.
func (m *MockClusterGitRepo) UpdatePipelineOutput(ctx context.Context, application, cluster, template string, pipelineOutput interface{}) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineOutput", ctx, application, cluster, template, pipelineOutput)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipelineOutput indicates an expected call of UpdatePipelineOutput.
func (mr *MockClusterGitRepoMockRecorder) UpdatePipelineOutput(ctx, application, cluster, template, pipelineOutput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineOutput", reflect.TypeOf((*MockClusterGitRepo)(nil).UpdatePipelineOutput), ctx, application, cluster, template, pipelineOutput)
}

// UpdateRestartTime mocks base method.
func (m *MockClusterGitRepo) UpdateRestartTime(ctx context.Context, application, cluster, template string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRestartTime", ctx, application, cluster, template)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRestartTime indicates an expected call of UpdateRestartTime.
func (mr *MockClusterGitRepoMockRecorder) UpdateRestartTime(ctx, application, cluster, template interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRestartTime", reflect.TypeOf((*MockClusterGitRepo)(nil).UpdateRestartTime), ctx, application, cluster, template)
}

// UpdateTags mocks base method.
func (m *MockClusterGitRepo) UpdateTags(ctx context.Context, application, cluster, templateName string, tags []*models.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", ctx, application, cluster, templateName, tags)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTags indicates an expected call of UpdateTags.
func (mr *MockClusterGitRepoMockRecorder) UpdateTags(ctx, application, cluster, templateName, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockClusterGitRepo)(nil).UpdateTags), ctx, application, cluster, templateName, tags)
}

// UpgradeCluster mocks base method.
func (m *MockClusterGitRepo) UpgradeCluster(ctx context.Context, param *gitrepo.UpgradeValuesParam) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeCluster", ctx, param)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeCluster indicates an expected call of UpgradeCluster.
func (mr *MockClusterGitRepoMockRecorder) UpgradeCluster(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCluster", reflect.TypeOf((*MockClusterGitRepo)(nil).UpgradeCluster), ctx, param)
}
