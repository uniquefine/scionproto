// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/pkg/ca/renewal (interfaces: CACertProvider,PolicyGen,DB)

// Package mock_renewal is a generated GoMock package.
package mock_renewal

import (
	context "context"
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	cppki "github.com/scionproto/scion/go/lib/scrypto/cppki"
	trust "github.com/scionproto/scion/go/pkg/trust"
	reflect "reflect"
)

// MockCACertProvider is a mock of CACertProvider interface
type MockCACertProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCACertProviderMockRecorder
}

// MockCACertProviderMockRecorder is the mock recorder for MockCACertProvider
type MockCACertProviderMockRecorder struct {
	mock *MockCACertProvider
}

// NewMockCACertProvider creates a new mock instance
func NewMockCACertProvider(ctrl *gomock.Controller) *MockCACertProvider {
	mock := &MockCACertProvider{ctrl: ctrl}
	mock.recorder = &MockCACertProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCACertProvider) EXPECT() *MockCACertProviderMockRecorder {
	return m.recorder
}

// CACerts mocks base method
func (m *MockCACertProvider) CACerts(arg0 context.Context) ([]*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CACerts", arg0)
	ret0, _ := ret[0].([]*x509.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CACerts indicates an expected call of CACerts
func (mr *MockCACertProviderMockRecorder) CACerts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CACerts", reflect.TypeOf((*MockCACertProvider)(nil).CACerts), arg0)
}

// MockPolicyGen is a mock of PolicyGen interface
type MockPolicyGen struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyGenMockRecorder
}

// MockPolicyGenMockRecorder is the mock recorder for MockPolicyGen
type MockPolicyGenMockRecorder struct {
	mock *MockPolicyGen
}

// NewMockPolicyGen creates a new mock instance
func NewMockPolicyGen(ctrl *gomock.Controller) *MockPolicyGen {
	mock := &MockPolicyGen{ctrl: ctrl}
	mock.recorder = &MockPolicyGenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyGen) EXPECT() *MockPolicyGenMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockPolicyGen) Generate(arg0 context.Context) (cppki.CAPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(cppki.CAPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate
func (mr *MockPolicyGenMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockPolicyGen)(nil).Generate), arg0)
}

// MockDB is a mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// ClientChains mocks base method
func (m *MockDB) ClientChains(arg0 context.Context, arg1 trust.ChainQuery) ([][]*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientChains", arg0, arg1)
	ret0, _ := ret[0].([][]*x509.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientChains indicates an expected call of ClientChains
func (mr *MockDBMockRecorder) ClientChains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientChains", reflect.TypeOf((*MockDB)(nil).ClientChains), arg0, arg1)
}

// Close mocks base method
func (m *MockDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// InsertClientChain mocks base method
func (m *MockDB) InsertClientChain(arg0 context.Context, arg1 []*x509.Certificate) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertClientChain", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertClientChain indicates an expected call of InsertClientChain
func (mr *MockDBMockRecorder) InsertClientChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertClientChain", reflect.TypeOf((*MockDB)(nil).InsertClientChain), arg0, arg1)
}

// SetMaxIdleConns mocks base method
func (m *MockDB) SetMaxIdleConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxIdleConns", arg0)
}

// SetMaxIdleConns indicates an expected call of SetMaxIdleConns
func (mr *MockDBMockRecorder) SetMaxIdleConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxIdleConns", reflect.TypeOf((*MockDB)(nil).SetMaxIdleConns), arg0)
}

// SetMaxOpenConns mocks base method
func (m *MockDB) SetMaxOpenConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxOpenConns", arg0)
}

// SetMaxOpenConns indicates an expected call of SetMaxOpenConns
func (mr *MockDBMockRecorder) SetMaxOpenConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxOpenConns", reflect.TypeOf((*MockDB)(nil).SetMaxOpenConns), arg0)
}