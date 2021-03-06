// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// ValidatorInterface is an autogenerated mock type for the ValidatorInterface type
type ValidatorInterface struct {
	mock.Mock
}

// ValidateRequest provides a mock function with given fields: s
func (_m *ValidatorInterface) ValidateRequest(s interface{}) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateRequestWithGetBody provides a mock function with given fields: c, s
func (_m *ValidatorInterface) ValidateRequestWithGetBody(c *gin.Context, s interface{}) error {
	ret := _m.Called(c, s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context, interface{}) error); ok {
		r0 = rf(c, s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewValidatorInterfaceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewValidatorInterface creates a new instance of ValidatorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewValidatorInterface(t NewValidatorInterfaceT) *ValidatorInterface {
	mock := &ValidatorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
