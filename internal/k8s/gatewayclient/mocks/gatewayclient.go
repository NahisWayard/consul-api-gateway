// Code generated by MockGen. DO NOT EDIT.
// Source: ./gatewayclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/hashicorp/consul-api-gateway/pkg/apis/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	v1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateDeployment mocks base method.
func (m *MockClient) CreateOrUpdateDeployment(ctx context.Context, deployment *v1.Deployment, mutators ...func() error) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, deployment}
	for _, a := range mutators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateDeployment", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateDeployment indicates an expected call of CreateOrUpdateDeployment.
func (mr *MockClientMockRecorder) CreateOrUpdateDeployment(ctx, deployment interface{}, mutators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, deployment}, mutators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateDeployment", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateDeployment), varargs...)
}

// CreateOrUpdateService mocks base method.
func (m *MockClient) CreateOrUpdateService(ctx context.Context, service *v10.Service, mutators ...func() error) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, service}
	for _, a := range mutators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateService", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateService indicates an expected call of CreateOrUpdateService.
func (mr *MockClientMockRecorder) CreateOrUpdateService(ctx, service interface{}, mutators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, service}, mutators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateService", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateService), varargs...)
}

// DeleteService mocks base method.
func (m *MockClient) DeleteService(ctx context.Context, service *v10.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockClientMockRecorder) DeleteService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockClient)(nil).DeleteService), ctx, service)
}

// DeploymentForGateway mocks base method.
func (m *MockClient) DeploymentForGateway(ctx context.Context, gw *v1beta1.Gateway) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentForGateway", ctx, gw)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentForGateway indicates an expected call of DeploymentForGateway.
func (mr *MockClientMockRecorder) DeploymentForGateway(ctx, gw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentForGateway", reflect.TypeOf((*MockClient)(nil).DeploymentForGateway), ctx, gw)
}

// EnsureFinalizer mocks base method.
func (m *MockClient) EnsureFinalizer(ctx context.Context, object client.Object, finalizer string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFinalizer", ctx, object, finalizer)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureFinalizer indicates an expected call of EnsureFinalizer.
func (mr *MockClientMockRecorder) EnsureFinalizer(ctx, object, finalizer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFinalizer", reflect.TypeOf((*MockClient)(nil).EnsureFinalizer), ctx, object, finalizer)
}

// EnsureServiceAccount mocks base method.
func (m *MockClient) EnsureServiceAccount(ctx context.Context, owner *v1beta1.Gateway, serviceAccount *v10.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureServiceAccount", ctx, owner, serviceAccount)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureServiceAccount indicates an expected call of EnsureServiceAccount.
func (mr *MockClientMockRecorder) EnsureServiceAccount(ctx, owner, serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureServiceAccount", reflect.TypeOf((*MockClient)(nil).EnsureServiceAccount), ctx, owner, serviceAccount)
}

// GatewayClassConfigInUse mocks base method.
func (m *MockClient) GatewayClassConfigInUse(ctx context.Context, gcc *v1alpha1.GatewayClassConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassConfigInUse", ctx, gcc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassConfigInUse indicates an expected call of GatewayClassConfigInUse.
func (mr *MockClientMockRecorder) GatewayClassConfigInUse(ctx, gcc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassConfigInUse", reflect.TypeOf((*MockClient)(nil).GatewayClassConfigInUse), ctx, gcc)
}

// GatewayClassInUse mocks base method.
func (m *MockClient) GatewayClassInUse(ctx context.Context, gc *v1beta1.GatewayClass) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassInUse", ctx, gc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassInUse indicates an expected call of GatewayClassInUse.
func (mr *MockClientMockRecorder) GatewayClassInUse(ctx, gc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassInUse", reflect.TypeOf((*MockClient)(nil).GatewayClassInUse), ctx, gc)
}

// GatewayClassesUsingConfig mocks base method.
func (m *MockClient) GatewayClassesUsingConfig(ctx context.Context, gcc *v1alpha1.GatewayClassConfig) (*v1beta1.GatewayClassList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassesUsingConfig", ctx, gcc)
	ret0, _ := ret[0].(*v1beta1.GatewayClassList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassesUsingConfig indicates an expected call of GatewayClassesUsingConfig.
func (mr *MockClientMockRecorder) GatewayClassesUsingConfig(ctx, gcc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassesUsingConfig", reflect.TypeOf((*MockClient)(nil).GatewayClassesUsingConfig), ctx, gcc)
}

// GetConfigForGatewayClassName mocks base method.
func (m *MockClient) GetConfigForGatewayClassName(ctx context.Context, name string) (v1alpha1.GatewayClassConfig, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigForGatewayClassName", ctx, name)
	ret0, _ := ret[0].(v1alpha1.GatewayClassConfig)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConfigForGatewayClassName indicates an expected call of GetConfigForGatewayClassName.
func (mr *MockClientMockRecorder) GetConfigForGatewayClassName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigForGatewayClassName", reflect.TypeOf((*MockClient)(nil).GetConfigForGatewayClassName), ctx, name)
}

// GetDeployment mocks base method.
func (m *MockClient) GetDeployment(ctx context.Context, key types.NamespacedName) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, key)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockClientMockRecorder) GetDeployment(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockClient)(nil).GetDeployment), ctx, key)
}

// GetGateway mocks base method.
func (m *MockClient) GetGateway(ctx context.Context, key types.NamespacedName) (*v1beta1.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGateway", ctx, key)
	ret0, _ := ret[0].(*v1beta1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGateway indicates an expected call of GetGateway.
func (mr *MockClientMockRecorder) GetGateway(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGateway", reflect.TypeOf((*MockClient)(nil).GetGateway), ctx, key)
}

// GetGatewayClass mocks base method.
func (m *MockClient) GetGatewayClass(ctx context.Context, key types.NamespacedName) (*v1beta1.GatewayClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayClass", ctx, key)
	ret0, _ := ret[0].(*v1beta1.GatewayClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayClass indicates an expected call of GetGatewayClass.
func (mr *MockClientMockRecorder) GetGatewayClass(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayClass", reflect.TypeOf((*MockClient)(nil).GetGatewayClass), ctx, key)
}

// GetGatewayClassConfig mocks base method.
func (m *MockClient) GetGatewayClassConfig(ctx context.Context, key types.NamespacedName) (*v1alpha1.GatewayClassConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayClassConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.GatewayClassConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayClassConfig indicates an expected call of GetGatewayClassConfig.
func (mr *MockClientMockRecorder) GetGatewayClassConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayClassConfig", reflect.TypeOf((*MockClient)(nil).GetGatewayClassConfig), ctx, key)
}

// GetGatewaysInNamespace mocks base method.
func (m *MockClient) GetGatewaysInNamespace(ctx context.Context, ns string) ([]v1beta1.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewaysInNamespace", ctx, ns)
	ret0, _ := ret[0].([]v1beta1.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewaysInNamespace indicates an expected call of GetGatewaysInNamespace.
func (mr *MockClientMockRecorder) GetGatewaysInNamespace(ctx, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewaysInNamespace", reflect.TypeOf((*MockClient)(nil).GetGatewaysInNamespace), ctx, ns)
}

// GetHTTPRoute mocks base method.
func (m *MockClient) GetHTTPRoute(ctx context.Context, key types.NamespacedName) (*v1alpha2.HTTPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRoute", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.HTTPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPRoute indicates an expected call of GetHTTPRoute.
func (mr *MockClientMockRecorder) GetHTTPRoute(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRoute", reflect.TypeOf((*MockClient)(nil).GetHTTPRoute), ctx, key)
}

// GetHTTPRoutesInNamespace mocks base method.
func (m *MockClient) GetHTTPRoutesInNamespace(ctx context.Context, ns string) ([]v1alpha2.HTTPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRoutesInNamespace", ctx, ns)
	ret0, _ := ret[0].([]v1alpha2.HTTPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPRoutesInNamespace indicates an expected call of GetHTTPRoutesInNamespace.
func (mr *MockClientMockRecorder) GetHTTPRoutesInNamespace(ctx, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRoutesInNamespace", reflect.TypeOf((*MockClient)(nil).GetHTTPRoutesInNamespace), ctx, ns)
}

// GetMeshService mocks base method.
func (m *MockClient) GetMeshService(ctx context.Context, key types.NamespacedName) (*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshService", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshService indicates an expected call of GetMeshService.
func (mr *MockClientMockRecorder) GetMeshService(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshService", reflect.TypeOf((*MockClient)(nil).GetMeshService), ctx, key)
}

// GetNamespace mocks base method.
func (m *MockClient) GetNamespace(ctx context.Context, key types.NamespacedName) (*v10.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", ctx, key)
	ret0, _ := ret[0].(*v10.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockClientMockRecorder) GetNamespace(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockClient)(nil).GetNamespace), ctx, key)
}

// GetReferenceGrantsInNamespace mocks base method.
func (m *MockClient) GetReferenceGrantsInNamespace(ctx context.Context, namespace string) ([]v1alpha2.ReferenceGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReferenceGrantsInNamespace", ctx, namespace)
	ret0, _ := ret[0].([]v1alpha2.ReferenceGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReferenceGrantsInNamespace indicates an expected call of GetReferenceGrantsInNamespace.
func (mr *MockClientMockRecorder) GetReferenceGrantsInNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReferenceGrantsInNamespace", reflect.TypeOf((*MockClient)(nil).GetReferenceGrantsInNamespace), ctx, namespace)
}

// GetSecret mocks base method.
func (m *MockClient) GetSecret(ctx context.Context, key types.NamespacedName) (*v10.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, key)
	ret0, _ := ret[0].(*v10.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockClientMockRecorder) GetSecret(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockClient)(nil).GetSecret), ctx, key)
}

// GetService mocks base method.
func (m *MockClient) GetService(ctx context.Context, key types.NamespacedName) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, key)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockClientMockRecorder) GetService(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockClient)(nil).GetService), ctx, key)
}

// GetTCPRoute mocks base method.
func (m *MockClient) GetTCPRoute(ctx context.Context, key types.NamespacedName) (*v1alpha2.TCPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTCPRoute", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.TCPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTCPRoute indicates an expected call of GetTCPRoute.
func (mr *MockClientMockRecorder) GetTCPRoute(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTCPRoute", reflect.TypeOf((*MockClient)(nil).GetTCPRoute), ctx, key)
}

// GetTCPRoutesInNamespace mocks base method.
func (m *MockClient) GetTCPRoutesInNamespace(ctx context.Context, ns string) ([]v1alpha2.TCPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTCPRoutesInNamespace", ctx, ns)
	ret0, _ := ret[0].([]v1alpha2.TCPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTCPRoutesInNamespace indicates an expected call of GetTCPRoutesInNamespace.
func (mr *MockClientMockRecorder) GetTCPRoutesInNamespace(ctx, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTCPRoutesInNamespace", reflect.TypeOf((*MockClient)(nil).GetTCPRoutesInNamespace), ctx, ns)
}

// HasManagedDeployment mocks base method.
func (m *MockClient) HasManagedDeployment(ctx context.Context, gw *v1beta1.Gateway) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasManagedDeployment", ctx, gw)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasManagedDeployment indicates an expected call of HasManagedDeployment.
func (mr *MockClientMockRecorder) HasManagedDeployment(ctx, gw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasManagedDeployment", reflect.TypeOf((*MockClient)(nil).HasManagedDeployment), ctx, gw)
}

// IsManagedRoute mocks base method.
func (m *MockClient) IsManagedRoute(ctx context.Context, namespace string, parents []v1alpha2.ParentReference) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManagedRoute", ctx, namespace, parents)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManagedRoute indicates an expected call of IsManagedRoute.
func (mr *MockClientMockRecorder) IsManagedRoute(ctx, namespace, parents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManagedRoute", reflect.TypeOf((*MockClient)(nil).IsManagedRoute), ctx, namespace, parents)
}

// PodsWithLabels mocks base method.
func (m *MockClient) PodsWithLabels(ctx context.Context, labels map[string]string) ([]v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodsWithLabels", ctx, labels)
	ret0, _ := ret[0].([]v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PodsWithLabels indicates an expected call of PodsWithLabels.
func (mr *MockClientMockRecorder) PodsWithLabels(ctx, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodsWithLabels", reflect.TypeOf((*MockClient)(nil).PodsWithLabels), ctx, labels)
}

// RemoveFinalizer mocks base method.
func (m *MockClient) RemoveFinalizer(ctx context.Context, object client.Object, finalizer string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFinalizer", ctx, object, finalizer)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFinalizer indicates an expected call of RemoveFinalizer.
func (mr *MockClientMockRecorder) RemoveFinalizer(ctx, object, finalizer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizer", reflect.TypeOf((*MockClient)(nil).RemoveFinalizer), ctx, object, finalizer)
}

// SetControllerOwnership mocks base method.
func (m *MockClient) SetControllerOwnership(owner, object client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerOwnership", owner, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerOwnership indicates an expected call of SetControllerOwnership.
func (mr *MockClientMockRecorder) SetControllerOwnership(owner, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerOwnership", reflect.TypeOf((*MockClient)(nil).SetControllerOwnership), owner, object)
}

// Update mocks base method.
func (m *MockClient) Update(ctx context.Context, obj client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockClientMockRecorder) Update(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), ctx, obj)
}

// UpdateStatus mocks base method.
func (m *MockClient) UpdateStatus(ctx context.Context, obj client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockClientMockRecorder) UpdateStatus(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClient)(nil).UpdateStatus), ctx, obj)
}
