// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PromoteUsecase is an autogenerated mock type for the PromoteUsecase type
type PromoteUsecase struct {
	mock.Mock
}

// DemoteUser provides a mock function with given fields: c, userID
func (_m *PromoteUsecase) DemoteUser(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PromoteUser provides a mock function with given fields: c, userID
func (_m *PromoteUsecase) PromoteUser(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPromoteUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromoteUsecase creates a new instance of PromoteUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromoteUsecase(t mockConstructorTestingTNewPromoteUsecase) *PromoteUsecase {
	mock := &PromoteUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
