// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewReadCloser creates a new instance of ReadCloser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReadCloser(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReadCloser {
	mock := &ReadCloser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// ReadCloser is an autogenerated mock type for the ReadCloser type
type ReadCloser struct {
	mock.Mock
}

type ReadCloser_Expecter struct {
	mock *mock.Mock
}

func (_m *ReadCloser) EXPECT() *ReadCloser_Expecter {
	return &ReadCloser_Expecter{mock: &_m.Mock}
}

// Close provides a mock function for the type ReadCloser
func (_mock *ReadCloser) Close() error {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func() error); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// ReadCloser_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type ReadCloser_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *ReadCloser_Expecter) Close() *ReadCloser_Close_Call {
	return &ReadCloser_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *ReadCloser_Close_Call) Run(run func()) *ReadCloser_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadCloser_Close_Call) Return(err error) *ReadCloser_Close_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *ReadCloser_Close_Call) RunAndReturn(run func() error) *ReadCloser_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function for the type ReadCloser
func (_mock *ReadCloser) Read(p []byte) (int, error) {
	ret := _mock.Called(p)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 int
	var r1 error
	if returnFunc, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return returnFunc(p)
	}
	if returnFunc, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = returnFunc(p)
	} else {
		r0 = ret.Get(0).(int)
	}
	if returnFunc, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = returnFunc(p)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// ReadCloser_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type ReadCloser_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - p []byte
func (_e *ReadCloser_Expecter) Read(p interface{}) *ReadCloser_Read_Call {
	return &ReadCloser_Read_Call{Call: _e.mock.On("Read", p)}
}

func (_c *ReadCloser_Read_Call) Run(run func(p []byte)) *ReadCloser_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 []byte
		if args[0] != nil {
			arg0 = args[0].([]byte)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *ReadCloser_Read_Call) Return(n int, err error) *ReadCloser_Read_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *ReadCloser_Read_Call) RunAndReturn(run func(p []byte) (int, error)) *ReadCloser_Read_Call {
	_c.Call.Return(run)
	return _c
}
