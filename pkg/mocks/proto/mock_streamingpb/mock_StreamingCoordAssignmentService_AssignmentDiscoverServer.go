// Code generated by mockery v2.53.3. DO NOT EDIT.

package mock_streamingpb

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	streamingpb "github.com/milvus-io/milvus/pkg/v2/proto/streamingpb"
)

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer is an autogenerated mock type for the StreamingCoordAssignmentService_AssignmentDiscoverServer type
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer struct {
	mock.Mock
}

type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) EXPECT() *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with no fields
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

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

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) Context() *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call) Run(run func()) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call) Return(_a0 context.Context) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call) RunAndReturn(run func() context.Context) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with no fields
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) Recv() (*streamingpb.AssignmentDiscoverRequest, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *streamingpb.AssignmentDiscoverRequest
	var r1 error
	if rf, ok := ret.Get(0).(func() (*streamingpb.AssignmentDiscoverRequest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *streamingpb.AssignmentDiscoverRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*streamingpb.AssignmentDiscoverRequest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) Recv() *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call) Run(run func()) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call) Return(_a0 *streamingpb.AssignmentDiscoverRequest, _a1 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call) RunAndReturn(run func() (*streamingpb.AssignmentDiscoverRequest, error)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) RecvMsg(m interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call) Run(run func(m interface{})) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call) Return(_a0 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) Send(_a0 *streamingpb.AssignmentDiscoverResponse) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*streamingpb.AssignmentDiscoverResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *streamingpb.AssignmentDiscoverResponse
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) Send(_a0 interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call) Run(run func(_a0 *streamingpb.AssignmentDiscoverResponse)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*streamingpb.AssignmentDiscoverResponse))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call) Return(_a0 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call) RunAndReturn(run func(*streamingpb.AssignmentDiscoverResponse) error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) SendHeader(_a0 interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call) Return(_a0 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) SendMsg(m interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call) Run(run func(m interface{})) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call) Return(_a0 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) SetHeader(_a0 interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call) Return(_a0 error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call) RunAndReturn(run func(metadata.MD) error) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockStreamingCoordAssignmentService_AssignmentDiscoverServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//   - _a0 metadata.MD
func (_e *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_Expecter) SetTrailer(_a0 interface{}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call {
	return &MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call) Return() *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call) RunAndReturn(run func(metadata.MD)) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer_SetTrailer_Call {
	_c.Run(run)
	return _c
}

// NewMockStreamingCoordAssignmentService_AssignmentDiscoverServer creates a new instance of MockStreamingCoordAssignmentService_AssignmentDiscoverServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamingCoordAssignmentService_AssignmentDiscoverServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamingCoordAssignmentService_AssignmentDiscoverServer {
	mock := &MockStreamingCoordAssignmentService_AssignmentDiscoverServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
