// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/group/manager/manager.go

// Package mock_manager is a generated GoMock package.
package mock_manager

import (
	context "context"
	reflect "reflect"

	models "g.hz.netease.com/horizon/pkg/group/models"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockManager) Create(ctx context.Context, group *models.Group) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, group)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockManagerMockRecorder) Create(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, group)
}

// Delete mocks base method.
func (m *MockManager) Delete(ctx context.Context, id uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockManagerMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockManager)(nil).Delete), ctx, id)
}

// GetByID mocks base method.
func (m *MockManager) GetByID(ctx context.Context, id uint) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockManagerMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockManager)(nil).GetByID), ctx, id)
}

// GetByIDNameFuzzily mocks base method.
func (m *MockManager) GetByIDNameFuzzily(ctx context.Context, id uint, name string) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDNameFuzzily", ctx, id, name)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDNameFuzzily indicates an expected call of GetByIDNameFuzzily.
func (mr *MockManagerMockRecorder) GetByIDNameFuzzily(ctx, id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDNameFuzzily", reflect.TypeOf((*MockManager)(nil).GetByIDNameFuzzily), ctx, id, name)
}

// GetByIDs mocks base method.
func (m *MockManager) GetByIDs(ctx context.Context, ids []uint) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockManagerMockRecorder) GetByIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockManager)(nil).GetByIDs), ctx, ids)
}

// GetByNameFuzzily mocks base method.
func (m *MockManager) GetByNameFuzzily(ctx context.Context, name string) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameFuzzily", ctx, name)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameFuzzily indicates an expected call of GetByNameFuzzily.
func (mr *MockManagerMockRecorder) GetByNameFuzzily(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameFuzzily", reflect.TypeOf((*MockManager)(nil).GetByNameFuzzily), ctx, name)
}

// GetByNameOrPathUnderParent mocks base method.
func (m *MockManager) GetByNameOrPathUnderParent(ctx context.Context, name, path string, parentID uint) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNameOrPathUnderParent", ctx, name, path, parentID)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNameOrPathUnderParent indicates an expected call of GetByNameOrPathUnderParent.
func (mr *MockManagerMockRecorder) GetByNameOrPathUnderParent(ctx, name, path, parentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNameOrPathUnderParent", reflect.TypeOf((*MockManager)(nil).GetByNameOrPathUnderParent), ctx, name, path, parentID)
}

// GetByPaths mocks base method.
func (m *MockManager) GetByPaths(ctx context.Context, paths []string) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPaths", ctx, paths)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPaths indicates an expected call of GetByPaths.
func (mr *MockManagerMockRecorder) GetByPaths(ctx, paths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPaths", reflect.TypeOf((*MockManager)(nil).GetByPaths), ctx, paths)
}

// GetChildren mocks base method.
func (m *MockManager) GetChildren(ctx context.Context, parentID uint, pageNumber, pageSize int) ([]*models.GroupOrApplication, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChildren", ctx, parentID, pageNumber, pageSize)
	ret0, _ := ret[0].([]*models.GroupOrApplication)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetChildren indicates an expected call of GetChildren.
func (mr *MockManagerMockRecorder) GetChildren(ctx, parentID, pageNumber, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChildren", reflect.TypeOf((*MockManager)(nil).GetChildren), ctx, parentID, pageNumber, pageSize)
}

// GetSubGroups mocks base method.
func (m *MockManager) GetSubGroups(ctx context.Context, id uint, pageNumber, pageSize int) ([]*models.Group, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubGroups", ctx, id, pageNumber, pageSize)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSubGroups indicates an expected call of GetSubGroups.
func (mr *MockManagerMockRecorder) GetSubGroups(ctx, id, pageNumber, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubGroups", reflect.TypeOf((*MockManager)(nil).GetSubGroups), ctx, id, pageNumber, pageSize)
}

// GetSubGroupsUnderParentIDs mocks base method.
func (m *MockManager) GetSubGroupsUnderParentIDs(ctx context.Context, parentIDs []uint) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubGroupsUnderParentIDs", ctx, parentIDs)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubGroupsUnderParentIDs indicates an expected call of GetSubGroupsUnderParentIDs.
func (mr *MockManagerMockRecorder) GetSubGroupsUnderParentIDs(ctx, parentIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubGroupsUnderParentIDs", reflect.TypeOf((*MockManager)(nil).GetSubGroupsUnderParentIDs), ctx, parentIDs)
}

// Transfer mocks base method.
func (m *MockManager) Transfer(ctx context.Context, id, newParentID, userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, id, newParentID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockManagerMockRecorder) Transfer(ctx, id, newParentID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockManager)(nil).Transfer), ctx, id, newParentID, userID)
}

// UpdateBasic mocks base method.
func (m *MockManager) UpdateBasic(ctx context.Context, group *models.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBasic", ctx, group)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBasic indicates an expected call of UpdateBasic.
func (mr *MockManagerMockRecorder) UpdateBasic(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBasic", reflect.TypeOf((*MockManager)(nil).UpdateBasic), ctx, group)
}
