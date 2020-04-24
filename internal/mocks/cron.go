// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cron "github.com/robfig/cron/v3"
	mock "github.com/stretchr/testify/mock"
)

// Cron is an autogenerated mock type for the Cron type
type Cron struct {
	mock.Mock
}

// AddFunc provides a mock function with given fields: _a0, _a1
func (_m *Cron) AddFunc(_a0 string, _a1 func()) (cron.EntryID, error) {
	ret := _m.Called(_a0, _a1)

	var r0 cron.EntryID
	if rf, ok := ret.Get(0).(func(string, func()) cron.EntryID); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(cron.EntryID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, func()) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Cron) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *Cron) Stop() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}