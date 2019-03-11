// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/cloudwatchlogs (interfaces: LogClientFactory)

// Package mock_cloudwatchlogs is a generated GoMock package.
package mock_cloudwatchlogs

import (
	reflect "reflect"

	cloudwatchlogs "github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

// MockLogClientFactory is a mock of LogClientFactory interface
type MockLogClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLogClientFactoryMockRecorder
}

// MockLogClientFactoryMockRecorder is the mock recorder for MockLogClientFactory
type MockLogClientFactoryMockRecorder struct {
	mock *MockLogClientFactory
}

// NewMockLogClientFactory creates a new mock instance
func NewMockLogClientFactory(ctrl *gomock.Controller) *MockLogClientFactory {
	mock := &MockLogClientFactory{ctrl: ctrl}
	mock.recorder = &MockLogClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogClientFactory) EXPECT() *MockLogClientFactoryMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockLogClientFactory) Get(arg0 string) cloudwatchlogs.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(cloudwatchlogs.Client)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockLogClientFactoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLogClientFactory)(nil).Get), arg0)
}
