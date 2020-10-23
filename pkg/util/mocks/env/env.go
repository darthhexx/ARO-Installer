// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/env (interfaces: Interface)

// Package mock_env is a generated GoMock package.
package mock_env

import (
	context "context"
	rsa "crypto/rsa"
	x509 "crypto/x509"
	net "net"
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"

	clientauthorizer "github.com/Azure/ARO-RP/pkg/util/clientauthorizer"
	deployment "github.com/Azure/ARO-RP/pkg/util/deployment"
	refreshable "github.com/Azure/ARO-RP/pkg/util/refreshable"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ACRName mocks base method
func (m *MockInterface) ACRName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRName indicates an expected call of ACRName
func (mr *MockInterfaceMockRecorder) ACRName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRName", reflect.TypeOf((*MockInterface)(nil).ACRName))
}

// ACRResourceID mocks base method
func (m *MockInterface) ACRResourceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRResourceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRResourceID indicates an expected call of ACRResourceID
func (mr *MockInterfaceMockRecorder) ACRResourceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRResourceID", reflect.TypeOf((*MockInterface)(nil).ACRResourceID))
}

// AROOperatorImage mocks base method
func (m *MockInterface) AROOperatorImage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AROOperatorImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// AROOperatorImage indicates an expected call of AROOperatorImage
func (mr *MockInterfaceMockRecorder) AROOperatorImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AROOperatorImage", reflect.TypeOf((*MockInterface)(nil).AROOperatorImage))
}

// AdminClientAuthorizer mocks base method
func (m *MockInterface) AdminClientAuthorizer() clientauthorizer.ClientAuthorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminClientAuthorizer")
	ret0, _ := ret[0].(clientauthorizer.ClientAuthorizer)
	return ret0
}

// AdminClientAuthorizer indicates an expected call of AdminClientAuthorizer
func (mr *MockInterfaceMockRecorder) AdminClientAuthorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminClientAuthorizer", reflect.TypeOf((*MockInterface)(nil).AdminClientAuthorizer))
}

// ArmClientAuthorizer mocks base method
func (m *MockInterface) ArmClientAuthorizer() clientauthorizer.ClientAuthorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArmClientAuthorizer")
	ret0, _ := ret[0].(clientauthorizer.ClientAuthorizer)
	return ret0
}

// ArmClientAuthorizer indicates an expected call of ArmClientAuthorizer
func (mr *MockInterfaceMockRecorder) ArmClientAuthorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArmClientAuthorizer", reflect.TypeOf((*MockInterface)(nil).ArmClientAuthorizer))
}

// ClustersGenevaLoggingConfigVersion mocks base method
func (m *MockInterface) ClustersGenevaLoggingConfigVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersGenevaLoggingConfigVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClustersGenevaLoggingConfigVersion indicates an expected call of ClustersGenevaLoggingConfigVersion
func (mr *MockInterfaceMockRecorder) ClustersGenevaLoggingConfigVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersGenevaLoggingConfigVersion", reflect.TypeOf((*MockInterface)(nil).ClustersGenevaLoggingConfigVersion))
}

// ClustersGenevaLoggingEnvironment mocks base method
func (m *MockInterface) ClustersGenevaLoggingEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersGenevaLoggingEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClustersGenevaLoggingEnvironment indicates an expected call of ClustersGenevaLoggingEnvironment
func (mr *MockInterfaceMockRecorder) ClustersGenevaLoggingEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersGenevaLoggingEnvironment", reflect.TypeOf((*MockInterface)(nil).ClustersGenevaLoggingEnvironment))
}

// ClustersGenevaLoggingSecret mocks base method
func (m *MockInterface) ClustersGenevaLoggingSecret() (*rsa.PrivateKey, *x509.Certificate) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersGenevaLoggingSecret")
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].(*x509.Certificate)
	return ret0, ret1
}

// ClustersGenevaLoggingSecret indicates an expected call of ClustersGenevaLoggingSecret
func (mr *MockInterfaceMockRecorder) ClustersGenevaLoggingSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersGenevaLoggingSecret", reflect.TypeOf((*MockInterface)(nil).ClustersGenevaLoggingSecret))
}

// ClustersKeyvaultURI mocks base method
func (m *MockInterface) ClustersKeyvaultURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersKeyvaultURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClustersKeyvaultURI indicates an expected call of ClustersKeyvaultURI
func (mr *MockInterfaceMockRecorder) ClustersKeyvaultURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersKeyvaultURI", reflect.TypeOf((*MockInterface)(nil).ClustersKeyvaultURI))
}

// CreateARMResourceGroupRoleAssignment mocks base method
func (m *MockInterface) CreateARMResourceGroupRoleAssignment(arg0 context.Context, arg1 refreshable.Authorizer, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateARMResourceGroupRoleAssignment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateARMResourceGroupRoleAssignment indicates an expected call of CreateARMResourceGroupRoleAssignment
func (mr *MockInterfaceMockRecorder) CreateARMResourceGroupRoleAssignment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateARMResourceGroupRoleAssignment", reflect.TypeOf((*MockInterface)(nil).CreateARMResourceGroupRoleAssignment), arg0, arg1, arg2)
}

// DeploymentMode mocks base method
func (m *MockInterface) DeploymentMode() deployment.Mode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentMode")
	ret0, _ := ret[0].(deployment.Mode)
	return ret0
}

// DeploymentMode indicates an expected call of DeploymentMode
func (mr *MockInterfaceMockRecorder) DeploymentMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentMode", reflect.TypeOf((*MockInterface)(nil).DeploymentMode))
}

// DialContext mocks base method
func (m *MockInterface) DialContext(arg0 context.Context, arg1, arg2 string) (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialContext indicates an expected call of DialContext
func (mr *MockInterfaceMockRecorder) DialContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialContext", reflect.TypeOf((*MockInterface)(nil).DialContext), arg0, arg1, arg2)
}

// Domain mocks base method
func (m *MockInterface) Domain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Domain")
	ret0, _ := ret[0].(string)
	return ret0
}

// Domain indicates an expected call of Domain
func (mr *MockInterfaceMockRecorder) Domain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Domain", reflect.TypeOf((*MockInterface)(nil).Domain))
}

// FPAuthorizer mocks base method
func (m *MockInterface) FPAuthorizer(arg0, arg1 string) (refreshable.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FPAuthorizer", arg0, arg1)
	ret0, _ := ret[0].(refreshable.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FPAuthorizer indicates an expected call of FPAuthorizer
func (mr *MockInterfaceMockRecorder) FPAuthorizer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FPAuthorizer", reflect.TypeOf((*MockInterface)(nil).FPAuthorizer), arg0, arg1)
}

// GetBase64Secret mocks base method
func (m *MockInterface) GetBase64Secret(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBase64Secret", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBase64Secret indicates an expected call of GetBase64Secret
func (mr *MockInterfaceMockRecorder) GetBase64Secret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBase64Secret", reflect.TypeOf((*MockInterface)(nil).GetBase64Secret), arg0, arg1)
}

// GetCertificateSecret mocks base method
func (m *MockInterface) GetCertificateSecret(arg0 context.Context, arg1 string) (*rsa.PrivateKey, []*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSecret", arg0, arg1)
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].([]*x509.Certificate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCertificateSecret indicates an expected call of GetCertificateSecret
func (mr *MockInterfaceMockRecorder) GetCertificateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSecret", reflect.TypeOf((*MockInterface)(nil).GetCertificateSecret), arg0, arg1)
}

// InitializeAuthorizers mocks base method
func (m *MockInterface) InitializeAuthorizers() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeAuthorizers")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeAuthorizers indicates an expected call of InitializeAuthorizers
func (mr *MockInterfaceMockRecorder) InitializeAuthorizers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeAuthorizers", reflect.TypeOf((*MockInterface)(nil).InitializeAuthorizers))
}

// Listen mocks base method
func (m *MockInterface) Listen() (net.Listener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listen")
	ret0, _ := ret[0].(net.Listener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listen indicates an expected call of Listen
func (mr *MockInterfaceMockRecorder) Listen() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockInterface)(nil).Listen))
}

// Location mocks base method
func (m *MockInterface) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location
func (mr *MockInterfaceMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockInterface)(nil).Location))
}

// NewRPAuthorizer mocks base method
func (m *MockInterface) NewRPAuthorizer(arg0 string) (autorest.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRPAuthorizer", arg0)
	ret0, _ := ret[0].(autorest.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRPAuthorizer indicates an expected call of NewRPAuthorizer
func (mr *MockInterfaceMockRecorder) NewRPAuthorizer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRPAuthorizer", reflect.TypeOf((*MockInterface)(nil).NewRPAuthorizer), arg0)
}

// ResourceGroup mocks base method
func (m *MockInterface) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup
func (mr *MockInterfaceMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockInterface)(nil).ResourceGroup))
}

// SubscriptionID mocks base method
func (m *MockInterface) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID
func (mr *MockInterfaceMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockInterface)(nil).SubscriptionID))
}

// TenantID mocks base method
func (m *MockInterface) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID
func (mr *MockInterfaceMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockInterface)(nil).TenantID))
}

// Zones mocks base method
func (m *MockInterface) Zones(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Zones", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Zones indicates an expected call of Zones
func (mr *MockInterfaceMockRecorder) Zones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Zones", reflect.TypeOf((*MockInterface)(nil).Zones), arg0)
}