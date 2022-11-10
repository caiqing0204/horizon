// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.go

// Package mock_gitlab is a generated GoMock package.
package mock_gitlab

import (
	context "context"
	reflect "reflect"

	gitlab "g.hz.netease.com/horizon/lib/gitlab"
	gomock "github.com/golang/mock/gomock"
	gitlab0 "github.com/xanzy/go-gitlab"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AcceptMR mocks base method.
func (m *MockInterface) AcceptMR(ctx context.Context, pid interface{}, mrID int, mergeCommitMsg *string, shouldRemoveSourceBranch *bool) (*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptMR", ctx, pid, mrID, mergeCommitMsg, shouldRemoveSourceBranch)
	ret0, _ := ret[0].(*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptMR indicates an expected call of AcceptMR.
func (mr *MockInterfaceMockRecorder) AcceptMR(ctx, pid, mrID, mergeCommitMsg, shouldRemoveSourceBranch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptMR", reflect.TypeOf((*MockInterface)(nil).AcceptMR), ctx, pid, mrID, mergeCommitMsg, shouldRemoveSourceBranch)
}

// CloseMR mocks base method.
func (m *MockInterface) CloseMR(ctx context.Context, pid interface{}, mrID int) (*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseMR", ctx, pid, mrID)
	ret0, _ := ret[0].(*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseMR indicates an expected call of CloseMR.
func (mr *MockInterfaceMockRecorder) CloseMR(ctx, pid, mrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseMR", reflect.TypeOf((*MockInterface)(nil).CloseMR), ctx, pid, mrID)
}

// Compare mocks base method.
func (m *MockInterface) Compare(ctx context.Context, pid interface{}, from, to string, straight *bool) (*gitlab0.Compare, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", ctx, pid, from, to, straight)
	ret0, _ := ret[0].(*gitlab0.Compare)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare.
func (mr *MockInterfaceMockRecorder) Compare(ctx, pid, from, to, straight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockInterface)(nil).Compare), ctx, pid, from, to, straight)
}

// CreateBranch mocks base method.
func (m *MockInterface) CreateBranch(ctx context.Context, pid interface{}, branch, fromRef string) (*gitlab0.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBranch", ctx, pid, branch, fromRef)
	ret0, _ := ret[0].(*gitlab0.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBranch indicates an expected call of CreateBranch.
func (mr *MockInterfaceMockRecorder) CreateBranch(ctx, pid, branch, fromRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBranch", reflect.TypeOf((*MockInterface)(nil).CreateBranch), ctx, pid, branch, fromRef)
}

// CreateGroup mocks base method.
func (m *MockInterface) CreateGroup(ctx context.Context, name, path string, parentID *int) (*gitlab0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, name, path, parentID)
	ret0, _ := ret[0].(*gitlab0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockInterfaceMockRecorder) CreateGroup(ctx, name, path, parentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockInterface)(nil).CreateGroup), ctx, name, path, parentID)
}

// CreateMR mocks base method.
func (m *MockInterface) CreateMR(ctx context.Context, pid interface{}, source, target, title string) (*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMR", ctx, pid, source, target, title)
	ret0, _ := ret[0].(*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMR indicates an expected call of CreateMR.
func (mr *MockInterfaceMockRecorder) CreateMR(ctx, pid, source, target, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMR", reflect.TypeOf((*MockInterface)(nil).CreateMR), ctx, pid, source, target, title)
}

// CreateProject mocks base method.
func (m *MockInterface) CreateProject(ctx context.Context, name string, groupID int) (*gitlab0.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, name, groupID)
	ret0, _ := ret[0].(*gitlab0.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockInterfaceMockRecorder) CreateProject(ctx, name, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockInterface)(nil).CreateProject), ctx, name, groupID)
}

// DeleteBranch mocks base method.
func (m *MockInterface) DeleteBranch(ctx context.Context, pid interface{}, branch string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBranch", ctx, pid, branch)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBranch indicates an expected call of DeleteBranch.
func (mr *MockInterfaceMockRecorder) DeleteBranch(ctx, pid, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBranch", reflect.TypeOf((*MockInterface)(nil).DeleteBranch), ctx, pid, branch)
}

// DeleteGroup mocks base method.
func (m *MockInterface) DeleteGroup(ctx context.Context, gid interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", ctx, gid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup.
func (mr *MockInterfaceMockRecorder) DeleteGroup(ctx, gid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockInterface)(nil).DeleteGroup), ctx, gid)
}

// DeleteProject mocks base method.
func (m *MockInterface) DeleteProject(ctx context.Context, pid interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, pid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockInterfaceMockRecorder) DeleteProject(ctx, pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockInterface)(nil).DeleteProject), ctx, pid)
}

// EditNameAndPathForProject mocks base method.
func (m *MockInterface) EditNameAndPathForProject(ctx context.Context, pid interface{}, newName, newPath *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditNameAndPathForProject", ctx, pid, newName, newPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditNameAndPathForProject indicates an expected call of EditNameAndPathForProject.
func (mr *MockInterfaceMockRecorder) EditNameAndPathForProject(ctx, pid, newName, newPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditNameAndPathForProject", reflect.TypeOf((*MockInterface)(nil).EditNameAndPathForProject), ctx, pid, newName, newPath)
}

// GetBranch mocks base method.
func (m *MockInterface) GetBranch(ctx context.Context, pid interface{}, branch string) (*gitlab0.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, pid, branch)
	ret0, _ := ret[0].(*gitlab0.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockInterfaceMockRecorder) GetBranch(ctx, pid, branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockInterface)(nil).GetBranch), ctx, pid, branch)
}

// GetCommit mocks base method.
func (m *MockInterface) GetCommit(ctx context.Context, pid interface{}, commit string) (*gitlab0.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, pid, commit)
	ret0, _ := ret[0].(*gitlab0.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockInterfaceMockRecorder) GetCommit(ctx, pid, commit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockInterface)(nil).GetCommit), ctx, pid, commit)
}

// GetCreatedGroup mocks base method.
func (m *MockInterface) GetCreatedGroup(ctx context.Context, parentID int, parentPath, name string) (*gitlab0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedGroup", ctx, parentID, parentPath, name)
	ret0, _ := ret[0].(*gitlab0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreatedGroup indicates an expected call of GetCreatedGroup.
func (mr *MockInterfaceMockRecorder) GetCreatedGroup(ctx, parentID, parentPath, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedGroup", reflect.TypeOf((*MockInterface)(nil).GetCreatedGroup), ctx, parentID, parentPath, name)
}

// GetFile mocks base method.
func (m *MockInterface) GetFile(ctx context.Context, pid interface{}, ref, filepath string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, pid, ref, filepath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockInterfaceMockRecorder) GetFile(ctx, pid, ref, filepath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockInterface)(nil).GetFile), ctx, pid, ref, filepath)
}

// GetGroup mocks base method.
func (m *MockInterface) GetGroup(ctx context.Context, gid interface{}) (*gitlab0.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, gid)
	ret0, _ := ret[0].(*gitlab0.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockInterfaceMockRecorder) GetGroup(ctx, gid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockInterface)(nil).GetGroup), ctx, gid)
}

// GetHTTPURL mocks base method.
func (m *MockInterface) GetHTTPURL(ctx context.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPURL", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHTTPURL indicates an expected call of GetHTTPURL.
func (mr *MockInterfaceMockRecorder) GetHTTPURL(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPURL", reflect.TypeOf((*MockInterface)(nil).GetHTTPURL), ctx)
}

// GetProject mocks base method.
func (m *MockInterface) GetProject(ctx context.Context, pid interface{}) (*gitlab0.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", ctx, pid)
	ret0, _ := ret[0].(*gitlab0.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockInterfaceMockRecorder) GetProject(ctx, pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockInterface)(nil).GetProject), ctx, pid)
}

// GetRepositoryArchive mocks base method.
func (m *MockInterface) GetRepositoryArchive(ctx context.Context, pid interface{}, sha string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepositoryArchive", ctx, pid, sha)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryArchive indicates an expected call of GetRepositoryArchive.
func (mr *MockInterfaceMockRecorder) GetRepositoryArchive(ctx, pid, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryArchive", reflect.TypeOf((*MockInterface)(nil).GetRepositoryArchive), ctx, pid, sha)
}

// GetSSHURL mocks base method.
func (m *MockInterface) GetSSHURL(ctx context.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSHURL", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSSHURL indicates an expected call of GetSSHURL.
func (mr *MockInterfaceMockRecorder) GetSSHURL(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSHURL", reflect.TypeOf((*MockInterface)(nil).GetSSHURL), ctx)
}

// GetTag mocks base method.
func (m *MockInterface) GetTag(ctx context.Context, pid interface{}, tag string) (*gitlab0.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", ctx, pid, tag)
	ret0, _ := ret[0].(*gitlab0.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag.
func (mr *MockInterfaceMockRecorder) GetTag(ctx, pid, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockInterface)(nil).GetTag), ctx, pid, tag)
}

// ListBranch mocks base method.
func (m *MockInterface) ListBranch(ctx context.Context, pid interface{}, listBranchOptions *gitlab0.ListBranchesOptions) ([]*gitlab0.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranch", ctx, pid, listBranchOptions)
	ret0, _ := ret[0].([]*gitlab0.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranch indicates an expected call of ListBranch.
func (mr *MockInterfaceMockRecorder) ListBranch(ctx, pid, listBranchOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranch", reflect.TypeOf((*MockInterface)(nil).ListBranch), ctx, pid, listBranchOptions)
}

// ListGroupProjects mocks base method.
func (m *MockInterface) ListGroupProjects(ctx context.Context, gid interface{}, page, perPage int) ([]*gitlab0.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupProjects", ctx, gid, page, perPage)
	ret0, _ := ret[0].([]*gitlab0.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupProjects indicates an expected call of ListGroupProjects.
func (mr *MockInterfaceMockRecorder) ListGroupProjects(ctx, gid, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupProjects", reflect.TypeOf((*MockInterface)(nil).ListGroupProjects), ctx, gid, page, perPage)
}

// ListMRs mocks base method.
func (m *MockInterface) ListMRs(ctx context.Context, pid interface{}, source, target, state string) ([]*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMRs", ctx, pid, source, target, state)
	ret0, _ := ret[0].([]*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMRs indicates an expected call of ListMRs.
func (mr *MockInterfaceMockRecorder) ListMRs(ctx, pid, source, target, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMRs", reflect.TypeOf((*MockInterface)(nil).ListMRs), ctx, pid, source, target, state)
}

// ListTag mocks base method.
func (m *MockInterface) ListTag(ctx context.Context, pid interface{}, listTagOptions *gitlab0.ListTagsOptions) ([]*gitlab0.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTag", ctx, pid, listTagOptions)
	ret0, _ := ret[0].([]*gitlab0.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTag indicates an expected call of ListTag.
func (mr *MockInterfaceMockRecorder) ListTag(ctx, pid, listTagOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTag", reflect.TypeOf((*MockInterface)(nil).ListTag), ctx, pid, listTagOptions)
}

// TransferProject mocks base method.
func (m *MockInterface) TransferProject(ctx context.Context, pid, gid interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferProject", ctx, pid, gid)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferProject indicates an expected call of TransferProject.
func (mr *MockInterfaceMockRecorder) TransferProject(ctx, pid, gid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferProject", reflect.TypeOf((*MockInterface)(nil).TransferProject), ctx, pid, gid)
}

// WriteFiles mocks base method.
func (m *MockInterface) WriteFiles(ctx context.Context, pid interface{}, branch, commitMsg string, startBranch *string, actions []gitlab.CommitAction) (*gitlab0.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFiles", ctx, pid, branch, commitMsg, startBranch, actions)
	ret0, _ := ret[0].(*gitlab0.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteFiles indicates an expected call of WriteFiles.
func (mr *MockInterfaceMockRecorder) WriteFiles(ctx, pid, branch, commitMsg, startBranch, actions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFiles", reflect.TypeOf((*MockInterface)(nil).WriteFiles), ctx, pid, branch, commitMsg, startBranch, actions)
}
