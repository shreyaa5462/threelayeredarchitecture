package user

import (
	models "ThreeLayeredArchitecture/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
	isgomock struct{}
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserServiceInterface) CreateUser(ctx *gofr.Context, user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceInterfaceMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceInterface)(nil).CreateUser), ctx, user)
}

// GetAllUsers mocks base method.
func (m *MockUserServiceInterface) GetAllUsers(ctx *gofr.Context) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserServiceInterfaceMockRecorder) GetAllUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserServiceInterface)(nil).GetAllUsers), ctx)
}

// GetUserByID mocks base method.
func (m *MockUserServiceInterface) GetUserByID(ctx *gofr.Context, id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceInterfaceMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserServiceInterface)(nil).GetUserByID), ctx, id)
}
