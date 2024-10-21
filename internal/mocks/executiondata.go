// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-emulator/internal (interfaces: ExecutionDataAPIClient)
//
// Generated by this command:
//
//	mockgen -destination=internal/mocks/executiondata.go -package=mocks github.com/onflow/flow-emulator/internal ExecutionDataAPIClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	executiondata "github.com/onflow/flow/protobuf/go/flow/executiondata"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockExecutionDataAPIClient is a mock of ExecutionDataAPIClient interface.
type MockExecutionDataAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionDataAPIClientMockRecorder
	isgomock struct{}
}

// MockExecutionDataAPIClientMockRecorder is the mock recorder for MockExecutionDataAPIClient.
type MockExecutionDataAPIClientMockRecorder struct {
	mock *MockExecutionDataAPIClient
}

// NewMockExecutionDataAPIClient creates a new mock instance.
func NewMockExecutionDataAPIClient(ctrl *gomock.Controller) *MockExecutionDataAPIClient {
	mock := &MockExecutionDataAPIClient{ctrl: ctrl}
	mock.recorder = &MockExecutionDataAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutionDataAPIClient) EXPECT() *MockExecutionDataAPIClientMockRecorder {
	return m.recorder
}

// GetExecutionDataByBlockID mocks base method.
func (m *MockExecutionDataAPIClient) GetExecutionDataByBlockID(ctx context.Context, in *executiondata.GetExecutionDataByBlockIDRequest, opts ...grpc.CallOption) (*executiondata.GetExecutionDataByBlockIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExecutionDataByBlockID", varargs...)
	ret0, _ := ret[0].(*executiondata.GetExecutionDataByBlockIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecutionDataByBlockID indicates an expected call of GetExecutionDataByBlockID.
func (mr *MockExecutionDataAPIClientMockRecorder) GetExecutionDataByBlockID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionDataByBlockID", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).GetExecutionDataByBlockID), varargs...)
}

// GetRegisterValues mocks base method.
func (m *MockExecutionDataAPIClient) GetRegisterValues(ctx context.Context, in *executiondata.GetRegisterValuesRequest, opts ...grpc.CallOption) (*executiondata.GetRegisterValuesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegisterValues", varargs...)
	ret0, _ := ret[0].(*executiondata.GetRegisterValuesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegisterValues indicates an expected call of GetRegisterValues.
func (mr *MockExecutionDataAPIClientMockRecorder) GetRegisterValues(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegisterValues", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).GetRegisterValues), varargs...)
}

// SubscribeAccountStatusesFromLatestBlock mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeAccountStatusesFromLatestBlock(ctx context.Context, in *executiondata.SubscribeAccountStatusesFromLatestBlockRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromLatestBlockClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeAccountStatusesFromLatestBlock", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromLatestBlockClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeAccountStatusesFromLatestBlock indicates an expected call of SubscribeAccountStatusesFromLatestBlock.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeAccountStatusesFromLatestBlock(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAccountStatusesFromLatestBlock", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeAccountStatusesFromLatestBlock), varargs...)
}

// SubscribeAccountStatusesFromStartBlockID mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeAccountStatusesFromStartBlockID(ctx context.Context, in *executiondata.SubscribeAccountStatusesFromStartBlockIDRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromStartBlockIDClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeAccountStatusesFromStartBlockID", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromStartBlockIDClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeAccountStatusesFromStartBlockID indicates an expected call of SubscribeAccountStatusesFromStartBlockID.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeAccountStatusesFromStartBlockID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAccountStatusesFromStartBlockID", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeAccountStatusesFromStartBlockID), varargs...)
}

// SubscribeAccountStatusesFromStartHeight mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeAccountStatusesFromStartHeight(ctx context.Context, in *executiondata.SubscribeAccountStatusesFromStartHeightRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromStartHeightClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeAccountStatusesFromStartHeight", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeAccountStatusesFromStartHeightClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeAccountStatusesFromStartHeight indicates an expected call of SubscribeAccountStatusesFromStartHeight.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeAccountStatusesFromStartHeight(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeAccountStatusesFromStartHeight", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeAccountStatusesFromStartHeight), varargs...)
}

// SubscribeEvents mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeEvents(ctx context.Context, in *executiondata.SubscribeEventsRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeEvents", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeEvents indicates an expected call of SubscribeEvents.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeEvents(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEvents", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeEvents), varargs...)
}

// SubscribeEventsFromLatest mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeEventsFromLatest(ctx context.Context, in *executiondata.SubscribeEventsFromLatestRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeEventsFromLatestClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeEventsFromLatest", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeEventsFromLatestClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeEventsFromLatest indicates an expected call of SubscribeEventsFromLatest.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeEventsFromLatest(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEventsFromLatest", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeEventsFromLatest), varargs...)
}

// SubscribeEventsFromStartBlockID mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeEventsFromStartBlockID(ctx context.Context, in *executiondata.SubscribeEventsFromStartBlockIDRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeEventsFromStartBlockIDClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeEventsFromStartBlockID", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeEventsFromStartBlockIDClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeEventsFromStartBlockID indicates an expected call of SubscribeEventsFromStartBlockID.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeEventsFromStartBlockID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEventsFromStartBlockID", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeEventsFromStartBlockID), varargs...)
}

// SubscribeEventsFromStartHeight mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeEventsFromStartHeight(ctx context.Context, in *executiondata.SubscribeEventsFromStartHeightRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeEventsFromStartHeightClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeEventsFromStartHeight", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeEventsFromStartHeightClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeEventsFromStartHeight indicates an expected call of SubscribeEventsFromStartHeight.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeEventsFromStartHeight(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEventsFromStartHeight", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeEventsFromStartHeight), varargs...)
}

// SubscribeExecutionData mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeExecutionData(ctx context.Context, in *executiondata.SubscribeExecutionDataRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeExecutionDataClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeExecutionData", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeExecutionDataClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeExecutionData indicates an expected call of SubscribeExecutionData.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeExecutionData(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeExecutionData", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeExecutionData), varargs...)
}

// SubscribeExecutionDataFromLatest mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeExecutionDataFromLatest(ctx context.Context, in *executiondata.SubscribeExecutionDataFromLatestRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeExecutionDataFromLatestClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeExecutionDataFromLatest", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeExecutionDataFromLatestClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeExecutionDataFromLatest indicates an expected call of SubscribeExecutionDataFromLatest.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeExecutionDataFromLatest(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeExecutionDataFromLatest", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeExecutionDataFromLatest), varargs...)
}

// SubscribeExecutionDataFromStartBlockHeight mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeExecutionDataFromStartBlockHeight(ctx context.Context, in *executiondata.SubscribeExecutionDataFromStartBlockHeightRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeExecutionDataFromStartBlockHeightClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeExecutionDataFromStartBlockHeight", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeExecutionDataFromStartBlockHeightClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeExecutionDataFromStartBlockHeight indicates an expected call of SubscribeExecutionDataFromStartBlockHeight.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeExecutionDataFromStartBlockHeight(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeExecutionDataFromStartBlockHeight", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeExecutionDataFromStartBlockHeight), varargs...)
}

// SubscribeExecutionDataFromStartBlockID mocks base method.
func (m *MockExecutionDataAPIClient) SubscribeExecutionDataFromStartBlockID(ctx context.Context, in *executiondata.SubscribeExecutionDataFromStartBlockIDRequest, opts ...grpc.CallOption) (executiondata.ExecutionDataAPI_SubscribeExecutionDataFromStartBlockIDClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubscribeExecutionDataFromStartBlockID", varargs...)
	ret0, _ := ret[0].(executiondata.ExecutionDataAPI_SubscribeExecutionDataFromStartBlockIDClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeExecutionDataFromStartBlockID indicates an expected call of SubscribeExecutionDataFromStartBlockID.
func (mr *MockExecutionDataAPIClientMockRecorder) SubscribeExecutionDataFromStartBlockID(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeExecutionDataFromStartBlockID", reflect.TypeOf((*MockExecutionDataAPIClient)(nil).SubscribeExecutionDataFromStartBlockID), varargs...)
}
