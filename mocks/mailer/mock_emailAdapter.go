// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks_mailer

import (
	context "context"

	"braces.dev/errtrace"
	mailer "github.com/dmitrymomot/mailer"
	mock "github.com/stretchr/testify/mock"
)

// MockemailAdapter is an autogenerated mock type for the emailAdapter type
type MockemailAdapter struct {
	mock.Mock
}

type MockemailAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockemailAdapter) EXPECT() *MockemailAdapter_Expecter {
	return &MockemailAdapter_Expecter{mock: &_m.Mock}
}

// SendEmail provides a mock function with given fields: ctx, payload
func (_m *MockemailAdapter) SendEmail(ctx context.Context, payload mailer.SendEmailPayload) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for SendEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, mailer.SendEmailPayload) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return errtrace.Wrap(r0)
}

// MockemailAdapter_SendEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmail'
type MockemailAdapter_SendEmail_Call struct {
	*mock.Call
}

// SendEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - payload mailer.SendEmailPayload
func (_e *MockemailAdapter_Expecter) SendEmail(ctx interface{}, payload interface{}) *MockemailAdapter_SendEmail_Call {
	return &MockemailAdapter_SendEmail_Call{Call: _e.mock.On("SendEmail", ctx, payload)}
}

func (_c *MockemailAdapter_SendEmail_Call) Run(run func(ctx context.Context, payload mailer.SendEmailPayload)) *MockemailAdapter_SendEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(mailer.SendEmailPayload))
	})
	return _c
}

func (_c *MockemailAdapter_SendEmail_Call) Return(_a0 error) *MockemailAdapter_SendEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockemailAdapter_SendEmail_Call) RunAndReturn(run func(context.Context, mailer.SendEmailPayload) error) *MockemailAdapter_SendEmail_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockemailAdapter creates a new instance of MockemailAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockemailAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockemailAdapter {
	mock := &MockemailAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
