// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/app/factory (interfaces: StateManager,SaveableOption)

// Package mock_factory is a generated GoMock package.
package mock_factory

import (
	reflect "reflect"

	config "github.com/aws/amazon-ecs-agent/agent/config"
	statemanager "github.com/aws/amazon-ecs-agent/agent/statemanager"
	gomock "github.com/golang/mock/gomock"
)

// MockStateManager is a mock of StateManager interface
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerMockRecorder
}

// MockStateManagerMockRecorder is the mock recorder for MockStateManager
type MockStateManagerMockRecorder struct {
	mock *MockStateManager
}

// NewMockStateManager creates a new mock instance
func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &MockStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStateManager) EXPECT() *MockStateManagerMockRecorder {
	return m.recorder
}

// NewStateManager mocks base method
func (m *MockStateManager) NewStateManager(arg0 *config.Config, arg1 ...statemanager.Option) (statemanager.StateManager, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewStateManager", varargs...)
	ret0, _ := ret[0].(statemanager.StateManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStateManager indicates an expected call of NewStateManager
func (mr *MockStateManagerMockRecorder) NewStateManager(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStateManager", reflect.TypeOf((*MockStateManager)(nil).NewStateManager), varargs...)
}

// MockSaveableOption is a mock of SaveableOption interface
type MockSaveableOption struct {
	ctrl     *gomock.Controller
	recorder *MockSaveableOptionMockRecorder
}

// MockSaveableOptionMockRecorder is the mock recorder for MockSaveableOption
type MockSaveableOptionMockRecorder struct {
	mock *MockSaveableOption
}

// NewMockSaveableOption creates a new mock instance
func NewMockSaveableOption(ctrl *gomock.Controller) *MockSaveableOption {
	mock := &MockSaveableOption{ctrl: ctrl}
	mock.recorder = &MockSaveableOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSaveableOption) EXPECT() *MockSaveableOptionMockRecorder {
	return m.recorder
}

// AddSaveable mocks base method
func (m *MockSaveableOption) AddSaveable(arg0 string, arg1 statemanager.Saveable) statemanager.Option {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSaveable", arg0, arg1)
	ret0, _ := ret[0].(statemanager.Option)
	return ret0
}

// AddSaveable indicates an expected call of AddSaveable
func (mr *MockSaveableOptionMockRecorder) AddSaveable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSaveable", reflect.TypeOf((*MockSaveableOption)(nil).AddSaveable), arg0, arg1)
}
