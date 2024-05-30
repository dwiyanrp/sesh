// Code generated by mockery v2.43.2. DO NOT EDIT.

package dir

import mock "github.com/stretchr/testify/mock"

// MockDir is an autogenerated mock type for the Dir type
type MockDir struct {
	mock.Mock
}

type MockDir_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDir) EXPECT() *MockDir_Expecter {
	return &MockDir_Expecter{mock: &_m.Mock}
}

// Dir provides a mock function with given fields: name
func (_m *MockDir) Dir(name string) (bool, string) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Dir")
	}

	var r0 bool
	var r1 string
	if rf, ok := ret.Get(0).(func(string) (bool, string)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MockDir_Dir_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dir'
type MockDir_Dir_Call struct {
	*mock.Call
}

// Dir is a helper method to define mock.On call
//   - name string
func (_e *MockDir_Expecter) Dir(name interface{}) *MockDir_Dir_Call {
	return &MockDir_Dir_Call{Call: _e.mock.On("Dir", name)}
}

func (_c *MockDir_Dir_Call) Run(run func(name string)) *MockDir_Dir_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDir_Dir_Call) Return(isDir bool, absPath string) *MockDir_Dir_Call {
	_c.Call.Return(isDir, absPath)
	return _c
}

func (_c *MockDir_Dir_Call) RunAndReturn(run func(string) (bool, string)) *MockDir_Dir_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDir creates a new instance of MockDir. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDir(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDir {
	mock := &MockDir{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
