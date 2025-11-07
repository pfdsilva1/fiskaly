package mocks

import (
    "reflect"

    "go.uber.org/mock/gomock"
    "github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
    "github.com/pfdsilva1/fiskaly/signing-service-challenge-go/service"
)

// MockSignatureDeviceService is a gomock-compatible mock for service.SignatureDeviceService.
type MockSignatureDeviceService struct {
    ctrl     *gomock.Controller
    recorder *MockSignatureDeviceServiceMockRecorder
}

// MockSignatureDeviceServiceMockRecorder records expectations for MockSignatureDeviceService.
type MockSignatureDeviceServiceMockRecorder struct {
    mock *MockSignatureDeviceService
}

// NewMockSignatureDeviceService creates a new mock instance.
func NewMockSignatureDeviceService(ctrl *gomock.Controller) *MockSignatureDeviceService {
    mock := &MockSignatureDeviceService{ctrl: ctrl}
    mock.recorder = &MockSignatureDeviceServiceMockRecorder{mock}
    return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignatureDeviceService) EXPECT() *MockSignatureDeviceServiceMockRecorder {
    return m.recorder
}

var _ service.SignatureDeviceService = (*MockSignatureDeviceService)(nil)

// NewSignatureDevice mocks base method.
func (m *MockSignatureDeviceService) NewSignatureDevice(algorithm string, label string) (string, error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "NewSignatureDevice", algorithm, label)
    ret0, _ := ret[0].(string)
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

// NewSignatureDevice indicates an expected call of NewSignatureDevice.
func (mr *MockSignatureDeviceServiceMockRecorder) NewSignatureDevice(algorithm, label interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSignatureDevice", reflect.TypeOf((*MockSignatureDeviceService)(nil).NewSignatureDevice), algorithm, label)
}

// SignTransaction mocks base method.
func (m *MockSignatureDeviceService) SignTransaction(deviceID string, data []byte) (domain.SignatureRecord, error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "SignTransaction", deviceID, data)
    ret0, _ := ret[0].(domain.SignatureRecord)
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

// SignTransaction indicates an expected call of SignTransaction.
func (mr *MockSignatureDeviceServiceMockRecorder) SignTransaction(deviceID, data interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockSignatureDeviceService)(nil).SignTransaction), deviceID, data)
}

// ListSignatureDevices mocks base method.
func (m *MockSignatureDeviceService) ListSignatureDevices() ([]*domain.SignatureDevice, error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListSignatureDevices")
    ret0, _ := ret[0].([]*domain.SignatureDevice)
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

// ListSignatureDevices indicates an expected call of ListSignatureDevices.
func (mr *MockSignatureDeviceServiceMockRecorder) ListSignatureDevices() *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSignatureDevices", reflect.TypeOf((*MockSignatureDeviceService)(nil).ListSignatureDevices))
}

// ListSignatures mocks base method.
func (m *MockSignatureDeviceService) ListSignatures(deviceID string) ([]domain.SignatureRecord, error) {
    m.ctrl.T.Helper()
    ret := m.ctrl.Call(m, "ListSignatures", deviceID)
    ret0, _ := ret[0].([]domain.SignatureRecord)
    ret1, _ := ret[1].(error)
    return ret0, ret1
}

// ListSignatures indicates an expected call of ListSignatures.
func (mr *MockSignatureDeviceServiceMockRecorder) ListSignatures(deviceID interface{}) *gomock.Call {
    mr.mock.ctrl.T.Helper()
    return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSignatures", reflect.TypeOf((*MockSignatureDeviceService)(nil).ListSignatures), deviceID)
}


