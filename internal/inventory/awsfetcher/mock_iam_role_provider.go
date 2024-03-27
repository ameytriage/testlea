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

package awsfetcher

import (
	context "context"

	iam "github.com/elastic/cloudbeat/internal/resources/providers/awslib/iam"
	mock "github.com/stretchr/testify/mock"
)

// mockIamRoleProvider is an autogenerated mock type for the iamRoleProvider type
type mockIamRoleProvider struct {
	mock.Mock
}

type mockIamRoleProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIamRoleProvider) EXPECT() *mockIamRoleProvider_Expecter {
	return &mockIamRoleProvider_Expecter{mock: &_m.Mock}
}

// ListRoles provides a mock function with given fields: ctx
func (_m *mockIamRoleProvider) ListRoles(ctx context.Context) ([]*iam.Role, error) {
	ret := _m.Called(ctx)

	var r0 []*iam.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*iam.Role, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*iam.Role); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*iam.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockIamRoleProvider_ListRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListRoles'
type mockIamRoleProvider_ListRoles_Call struct {
	*mock.Call
}

// ListRoles is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockIamRoleProvider_Expecter) ListRoles(ctx interface{}) *mockIamRoleProvider_ListRoles_Call {
	return &mockIamRoleProvider_ListRoles_Call{Call: _e.mock.On("ListRoles", ctx)}
}

func (_c *mockIamRoleProvider_ListRoles_Call) Run(run func(ctx context.Context)) *mockIamRoleProvider_ListRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockIamRoleProvider_ListRoles_Call) Return(_a0 []*iam.Role, _a1 error) *mockIamRoleProvider_ListRoles_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockIamRoleProvider_ListRoles_Call) RunAndReturn(run func(context.Context) ([]*iam.Role, error)) *mockIamRoleProvider_ListRoles_Call {
	_c.Call.Return(run)
	return _c
}

// newMockIamRoleProvider creates a new instance of mockIamRoleProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockIamRoleProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockIamRoleProvider {
	mock := &mockIamRoleProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
