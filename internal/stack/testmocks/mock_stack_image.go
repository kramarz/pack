// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpack/pack/internal/stack (interfaces: Image)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	imgutil "github.com/buildpack/imgutil"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
	time "time"
)

// MockImage is a mock of Image interface
type MockImage struct {
	ctrl     *gomock.Controller
	recorder *MockImageMockRecorder
}

// MockImageMockRecorder is the mock recorder for MockImage
type MockImageMockRecorder struct {
	mock *MockImage
}

// NewMockImage creates a new mock instance
func NewMockImage(ctrl *gomock.Controller) *MockImage {
	mock := &MockImage{ctrl: ctrl}
	mock.recorder = &MockImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImage) EXPECT() *MockImageMockRecorder {
	return m.recorder
}

// AddLayer mocks base method
func (m *MockImage) AddLayer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLayer indicates an expected call of AddLayer
func (mr *MockImageMockRecorder) AddLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLayer", reflect.TypeOf((*MockImage)(nil).AddLayer), arg0)
}

// CommonMixins mocks base method
func (m *MockImage) CommonMixins() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommonMixins")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CommonMixins indicates an expected call of CommonMixins
func (mr *MockImageMockRecorder) CommonMixins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonMixins", reflect.TypeOf((*MockImage)(nil).CommonMixins))
}

// CreatedAt mocks base method
func (m *MockImage) CreatedAt() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockImageMockRecorder) CreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockImage)(nil).CreatedAt))
}

// Delete mocks base method
func (m *MockImage) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockImageMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImage)(nil).Delete))
}

// Env mocks base method
func (m *MockImage) Env(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Env", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Env indicates an expected call of Env
func (mr *MockImageMockRecorder) Env(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*MockImage)(nil).Env), arg0)
}

// Found mocks base method
func (m *MockImage) Found() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Found")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Found indicates an expected call of Found
func (mr *MockImageMockRecorder) Found() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Found", reflect.TypeOf((*MockImage)(nil).Found))
}

// GetLayer mocks base method
func (m *MockImage) GetLayer(arg0 string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayer", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayer indicates an expected call of GetLayer
func (mr *MockImageMockRecorder) GetLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayer", reflect.TypeOf((*MockImage)(nil).GetLayer), arg0)
}

// Identifier mocks base method
func (m *MockImage) Identifier() (imgutil.Identifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(imgutil.Identifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identifier indicates an expected call of Identifier
func (mr *MockImageMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockImage)(nil).Identifier))
}

// Label mocks base method
func (m *MockImage) Label(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Label", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Label indicates an expected call of Label
func (mr *MockImageMockRecorder) Label(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Label", reflect.TypeOf((*MockImage)(nil).Label), arg0)
}

// Mixins mocks base method
func (m *MockImage) Mixins() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mixins")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Mixins indicates an expected call of Mixins
func (mr *MockImageMockRecorder) Mixins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mixins", reflect.TypeOf((*MockImage)(nil).Mixins))
}

// Name mocks base method
func (m *MockImage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockImageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockImage)(nil).Name))
}

// Rebase mocks base method
func (m *MockImage) Rebase(arg0 string, arg1 imgutil.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rebase indicates an expected call of Rebase
func (mr *MockImageMockRecorder) Rebase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebase", reflect.TypeOf((*MockImage)(nil).Rebase), arg0, arg1)
}

// Rename mocks base method
func (m *MockImage) Rename(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rename", arg0)
}

// Rename indicates an expected call of Rename
func (mr *MockImageMockRecorder) Rename(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockImage)(nil).Rename), arg0)
}

// ReuseLayer mocks base method
func (m *MockImage) ReuseLayer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReuseLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReuseLayer indicates an expected call of ReuseLayer
func (mr *MockImageMockRecorder) ReuseLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReuseLayer", reflect.TypeOf((*MockImage)(nil).ReuseLayer), arg0)
}

// Save mocks base method
func (m *MockImage) Save(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockImageMockRecorder) Save(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockImage)(nil).Save), arg0...)
}

// SetCmd mocks base method
func (m *MockImage) SetCmd(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetCmd", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCmd indicates an expected call of SetCmd
func (mr *MockImageMockRecorder) SetCmd(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCmd", reflect.TypeOf((*MockImage)(nil).SetCmd), arg0...)
}

// SetEntrypoint mocks base method
func (m *MockImage) SetEntrypoint(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetEntrypoint", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEntrypoint indicates an expected call of SetEntrypoint
func (mr *MockImageMockRecorder) SetEntrypoint(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEntrypoint", reflect.TypeOf((*MockImage)(nil).SetEntrypoint), arg0...)
}

// SetEnv mocks base method
func (m *MockImage) SetEnv(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEnv", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnv indicates an expected call of SetEnv
func (mr *MockImageMockRecorder) SetEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockImage)(nil).SetEnv), arg0, arg1)
}

// SetLabel mocks base method
func (m *MockImage) SetLabel(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLabel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLabel indicates an expected call of SetLabel
func (mr *MockImageMockRecorder) SetLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabel", reflect.TypeOf((*MockImage)(nil).SetLabel), arg0, arg1)
}

// SetWorkingDir mocks base method
func (m *MockImage) SetWorkingDir(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkingDir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkingDir indicates an expected call of SetWorkingDir
func (mr *MockImageMockRecorder) SetWorkingDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkingDir", reflect.TypeOf((*MockImage)(nil).SetWorkingDir), arg0)
}

// StackID mocks base method
func (m *MockImage) StackID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// StackID indicates an expected call of StackID
func (mr *MockImageMockRecorder) StackID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackID", reflect.TypeOf((*MockImage)(nil).StackID))
}

// TopLayer mocks base method
func (m *MockImage) TopLayer() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopLayer")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopLayer indicates an expected call of TopLayer
func (mr *MockImageMockRecorder) TopLayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopLayer", reflect.TypeOf((*MockImage)(nil).TopLayer))
}
