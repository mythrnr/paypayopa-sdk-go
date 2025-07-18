// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewMarshaler creates a new instance of Marshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMarshaler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Marshaler {
	mock := &Marshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Marshaler is an autogenerated mock type for the Marshaler type
type Marshaler struct {
	mock.Mock
}

type Marshaler_Expecter struct {
	mock *mock.Mock
}

func (_m *Marshaler) EXPECT() *Marshaler_Expecter {
	return &Marshaler_Expecter{mock: &_m.Mock}
}

// MarshalJSON provides a mock function for the type Marshaler
func (_mock *Marshaler) MarshalJSON() ([]byte, error) {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalJSON")
	}

	var r0 []byte
	var r1 error
	if returnFunc, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return returnFunc()
	}
	if returnFunc, ok := ret.Get(0).(func() []byte); ok {
		r0 = returnFunc()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}
	if returnFunc, ok := ret.Get(1).(func() error); ok {
		r1 = returnFunc()
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Marshaler_MarshalJSON_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalJSON'
type Marshaler_MarshalJSON_Call struct {
	*mock.Call
}

// MarshalJSON is a helper method to define mock.On call
func (_e *Marshaler_Expecter) MarshalJSON() *Marshaler_MarshalJSON_Call {
	return &Marshaler_MarshalJSON_Call{Call: _e.mock.On("MarshalJSON")}
}

func (_c *Marshaler_MarshalJSON_Call) Run(run func()) *Marshaler_MarshalJSON_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Marshaler_MarshalJSON_Call) Return(bytes []byte, err error) *Marshaler_MarshalJSON_Call {
	_c.Call.Return(bytes, err)
	return _c
}

func (_c *Marshaler_MarshalJSON_Call) RunAndReturn(run func() ([]byte, error)) *Marshaler_MarshalJSON_Call {
	_c.Call.Return(run)
	return _c
}
