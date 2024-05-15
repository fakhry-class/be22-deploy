// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	user "be22/clean-arch/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the DataInterface type
type UserData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *UserData) Delete(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: input
func (_m *UserData) Insert(input user.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *UserData) SelectAll() ([]user.Core, error) {
	ret := _m.Called()

	var r0 []user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]user.Core, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []user.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByEmail provides a mock function with given fields: email
func (_m *UserData) SelectByEmail(email string) (*user.Core, error) {
	ret := _m.Called(email)

	var r0 *user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*user.Core, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *user.Core); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *UserData) SelectById(id uint) (*user.Core, error) {
	ret := _m.Called(id)

	var r0 *user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*user.Core, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *user.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
