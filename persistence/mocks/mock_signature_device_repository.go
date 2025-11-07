package mocks

import (
	"reflect"

	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/domain"
	"github.com/pfdsilva1/fiskaly/signing-service-challenge-go/persistence"
	"go.uber.org/mock/gomock"
)

// MockSignatureDeviceRepository is a gomock-compatible mock for persistence.SignatureDeviceRepository.
type MockSignatureDeviceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSignatureDeviceRepositoryMockRecorder
}

// MockSignatureDeviceRepositoryMockRecorder records expectations for MockSignatureDeviceRepository.
type MockSignatureDeviceRepositoryMockRecorder struct {
	mock *MockSignatureDeviceRepository
}

// NewMockSignatureDeviceRepository creates a new mock instance.
func NewMockSignatureDeviceRepository(ctrl *gomock.Controller) *MockSignatureDeviceRepository {
	mock := &MockSignatureDeviceRepository{ctrl: ctrl}
	mock.recorder = &MockSignatureDeviceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignatureDeviceRepository) EXPECT() *MockSignatureDeviceRepositoryMockRecorder {
	return m.recorder
}

var _ persistence.SignatureDeviceRepository = (*MockSignatureDeviceRepository)(nil)

// CreateSignatureDevice mocks base method.
func (m *MockSignatureDeviceRepository) CreateSignatureDevice(signatureDevice *domain.SignatureDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSignatureDevice", signatureDevice)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSignatureDevice indicates an expected call of CreateSignatureDevice.
func (mr *MockSignatureDeviceRepositoryMockRecorder) CreateSignatureDevice(signatureDevice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignatureDevice", reflect.TypeOf((*MockSignatureDeviceRepository)(nil).CreateSignatureDevice), signatureDevice)
}

// GetSignatureDevice mocks base method.
func (m *MockSignatureDeviceRepository) GetSignatureDevice(id string) (*domain.SignatureDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignatureDevice", id)
	ret0, _ := ret[0].(*domain.SignatureDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignatureDevice indicates an expected call of GetSignatureDevice.
func (mr *MockSignatureDeviceRepositoryMockRecorder) GetSignatureDevice(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignatureDevice", reflect.TypeOf((*MockSignatureDeviceRepository)(nil).GetSignatureDevice), id)
}

// ListSignatureDevices mocks base method.
func (m *MockSignatureDeviceRepository) ListSignatureDevices() ([]*domain.SignatureDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSignatureDevices")
	ret0, _ := ret[0].([]*domain.SignatureDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSignatureDevices indicates an expected call of ListSignatureDevices.
func (mr *MockSignatureDeviceRepositoryMockRecorder) ListSignatureDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSignatureDevices", reflect.TypeOf((*MockSignatureDeviceRepository)(nil).ListSignatureDevices))
}

// SaveSignature mocks base method.
func (m *MockSignatureDeviceRepository) SaveSignature(deviceID string, signature domain.SignatureRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSignature", deviceID, signature)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSignature indicates an expected call of SaveSignature.
func (mr *MockSignatureDeviceRepositoryMockRecorder) SaveSignature(deviceID, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSignature", reflect.TypeOf((*MockSignatureDeviceRepository)(nil).SaveSignature), deviceID, signature)
}

// ListSignatures mocks base method.
func (m *MockSignatureDeviceRepository) ListSignatures(deviceID string) ([]domain.SignatureRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSignatures", deviceID)
	ret0, _ := ret[0].([]domain.SignatureRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSignatures indicates an expected call of ListSignatures.
func (mr *MockSignatureDeviceRepositoryMockRecorder) ListSignatures(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSignatures", reflect.TypeOf((*MockSignatureDeviceRepository)(nil).ListSignatures), deviceID)
}
