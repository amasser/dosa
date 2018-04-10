// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/uber-go/dosa (interfaces: Client)

package mocks

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	dosa "github.com/uber-go/dosa"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) CreateIfNotExists(_param0 context.Context, _param1 dosa.DomainObject) error {
	ret := _m.ctrl.Call(_m, "CreateIfNotExists", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) CreateIfNotExists(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateIfNotExists", arg0, arg1)
}

func (_m *MockClient) GetRegistrar() dosa.Registrar {
	ret := _m.ctrl.Call(_m, "GetRegistrar")
	ret0, _ := ret[0].(dosa.Registrar)
	return ret0
}

func (_mr *_MockClientRecorder) GetRegistrar() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRegistrar")
}

func (_m *MockClient) Initialize(_param0 context.Context) error {
	ret := _m.ctrl.Call(_m, "Initialize", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Initialize(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Initialize", arg0)
}

func (_m *MockClient) MultiRead(_param0 context.Context, _param1 []string, _param2 ...dosa.DomainObject) (dosa.MultiResult, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "MultiRead", _s...)
	ret0, _ := ret[0].(dosa.MultiResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) MultiRead(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiRead", _s...)
}

func (_m *MockClient) MultiUpsert(_param0 context.Context, _param1 []string, _param2 ...dosa.DomainObject) (dosa.MultiResult, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "MultiUpsert", _s...)
	ret0, _ := ret[0].(dosa.MultiResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) MultiUpsert(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MultiUpsert", _s...)
}

func (_m *MockClient) Range(_param0 context.Context, _param1 *dosa.RangeOp) ([]dosa.DomainObject, string, error) {
	ret := _m.ctrl.Call(_m, "Range", _param0, _param1)
	ret0, _ := ret[0].([]dosa.DomainObject)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockClientRecorder) Range(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Range", arg0, arg1)
}

func (_m *MockClient) Read(_param0 context.Context, _param1 []string, _param2 dosa.DomainObject) error {
	ret := _m.ctrl.Call(_m, "Read", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1, arg2)
}

func (_m *MockClient) Remove(_param0 context.Context, _param1 dosa.DomainObject) error {
	ret := _m.ctrl.Call(_m, "Remove", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0, arg1)
}

func (_m *MockClient) RemoveRange(_param0 context.Context, _param1 *dosa.RemoveRangeOp) error {
	ret := _m.ctrl.Call(_m, "RemoveRange", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) RemoveRange(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveRange", arg0, arg1)
}

func (_m *MockClient) ScanEverything(_param0 context.Context, _param1 *dosa.ScanOp) ([]dosa.DomainObject, string, error) {
	ret := _m.ctrl.Call(_m, "ScanEverything", _param0, _param1)
	ret0, _ := ret[0].([]dosa.DomainObject)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockClientRecorder) ScanEverything(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ScanEverything", arg0, arg1)
}

func (_m *MockClient) Upsert(_param0 context.Context, _param1 []string, _param2 dosa.DomainObject) error {
	ret := _m.ctrl.Call(_m, "Upsert", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Upsert(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Upsert", arg0, arg1, arg2)
}

func (_m *MockClient) WalkRange(_param0 context.Context, _param1 *dosa.RangeOp, _param2 func(dosa.DomainObject) error) error {
	ret := _m.ctrl.Call(_m, "WalkRange", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) WalkRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WalkRange", arg0, arg1, arg2)
}
