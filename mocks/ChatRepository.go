// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "abduselam-arabianmejlis/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ChatRepository is an autogenerated mock type for the ChatRepository type
type ChatRepository struct {
	mock.Mock
}

// CreateMessage provides a mock function with given fields: c, message
func (_m *ChatRepository) CreateMessage(c context.Context, message *domain.Message) error {
	ret := _m.Called(c, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Message) error); ok {
		r0 = rf(c, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessage provides a mock function with given fields: c, id
func (_m *ChatRepository) DeleteMessage(c context.Context, id string) error {
	ret := _m.Called(c, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMessagesByID provides a mock function with given fields: c, userID, adminID
func (_m *ChatRepository) GetMessagesByID(c context.Context, userID string, adminID string) ([]*domain.Message, error) {
	ret := _m.Called(c, userID, adminID)

	var r0 []*domain.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]*domain.Message, error)); ok {
		return rf(c, userID, adminID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*domain.Message); ok {
		r0 = rf(c, userID, adminID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(c, userID, adminID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewChatRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewChatRepository creates a new instance of ChatRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChatRepository(t mockConstructorTestingTNewChatRepository) *ChatRepository {
	mock := &ChatRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
