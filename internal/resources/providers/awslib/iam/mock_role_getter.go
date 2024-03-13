// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.37.1. DO NOT EDIT.

package iam

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRoleGetter is an autogenerated mock type for the RoleGetter type
type MockRoleGetter struct {
	mock.Mock
}

type MockRoleGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRoleGetter) EXPECT() *MockRoleGetter_Expecter {
	return &MockRoleGetter_Expecter{mock: &_m.Mock}
}

// GetRole provides a mock function with given fields: ctx, roleName
func (_m *MockRoleGetter) GetRole(ctx context.Context, roleName string) (*Role, error) {
	ret := _m.Called(ctx, roleName)

	var r0 *Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Role, error)); ok {
		return rf(ctx, roleName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Role); ok {
		r0 = rf(ctx, roleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, roleName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRoleGetter_GetRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRole'
type MockRoleGetter_GetRole_Call struct {
	*mock.Call
}

// GetRole is a helper method to define mock.On call
//   - ctx context.Context
//   - roleName string
func (_e *MockRoleGetter_Expecter) GetRole(ctx interface{}, roleName interface{}) *MockRoleGetter_GetRole_Call {
	return &MockRoleGetter_GetRole_Call{Call: _e.mock.On("GetRole", ctx, roleName)}
}

func (_c *MockRoleGetter_GetRole_Call) Run(run func(ctx context.Context, roleName string)) *MockRoleGetter_GetRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRoleGetter_GetRole_Call) Return(_a0 *Role, _a1 error) *MockRoleGetter_GetRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRoleGetter_GetRole_Call) RunAndReturn(run func(context.Context, string) (*Role, error)) *MockRoleGetter_GetRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRoleGetter creates a new instance of MockRoleGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRoleGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRoleGetter {
	mock := &MockRoleGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}