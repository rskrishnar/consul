// Code generated by mockery v1.0.0. DO NOT EDIT.

package autopilotevents

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockTimeProvider is an autogenerated mock type for the timeProvider type
type mockTimeProvider struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *mockTimeProvider) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}