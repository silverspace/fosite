// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite/handler/oauth2 (interfaces: ClientCredentialsGrantStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of ClientCredentialsGrantStorage interface
type MockClientCredentialsGrantStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockClientCredentialsGrantStorageRecorder
}

// Recorder for MockClientCredentialsGrantStorage (not exported)
type _MockClientCredentialsGrantStorageRecorder struct {
	mock *MockClientCredentialsGrantStorage
}

func NewMockClientCredentialsGrantStorage(ctrl *gomock.Controller) *MockClientCredentialsGrantStorage {
	mock := &MockClientCredentialsGrantStorage{ctrl: ctrl}
	mock.recorder = &_MockClientCredentialsGrantStorageRecorder{mock}
	return mock
}

func (_m *MockClientCredentialsGrantStorage) EXPECT() *_MockClientCredentialsGrantStorageRecorder {
	return _m.recorder
}

func (_m *MockClientCredentialsGrantStorage) CreateAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientCredentialsGrantStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockClientCredentialsGrantStorage) DeleteAccessTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientCredentialsGrantStorageRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0, arg1)
}

func (_m *MockClientCredentialsGrantStorage) GetAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientCredentialsGrantStorageRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1, arg2)
}
