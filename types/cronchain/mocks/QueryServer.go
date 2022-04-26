// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	cronchain "github.com/volumefi/conductor/types/cronchain"

	testing "testing"
)

// QueryServer is an autogenerated mock type for the QueryServer type
type QueryServer struct {
	mock.Mock
}

// QueuedMessagesForSigning provides a mock function with given fields: _a0, _a1
func (_m *QueryServer) QueuedMessagesForSigning(_a0 context.Context, _a1 *cronchain.QueryQueuedMessagesForSigningRequest) (*cronchain.QueryQueuedMessagesForSigningResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cronchain.QueryQueuedMessagesForSigningResponse
	if rf, ok := ret.Get(0).(func(context.Context, *cronchain.QueryQueuedMessagesForSigningRequest) *cronchain.QueryQueuedMessagesForSigningResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cronchain.QueryQueuedMessagesForSigningResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cronchain.QueryQueuedMessagesForSigningRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryServer creates a new instance of QueryServer. It also registers a cleanup function to assert the mocks expectations.
func NewQueryServer(t testing.TB) *QueryServer {
	mock := &QueryServer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
