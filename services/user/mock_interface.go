package user

import (
	models "ThreeLayeredArchitecture/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

type MockuserStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockuserStoreInterfaceMockRecorder
}

type MockuserStoreInterfaceMockRecorder struct {
	mock *MockuserStoreInterface
}

func NewMockuserStoreInterface(ctrl *gomock.Controller) *MockuserStoreInterface {
	mock := &MockuserStoreInterface{ctrl: ctrl}
	mock.recorder = &MockuserStoreInterfaceMockRecorder{mock}

	return mock
}

func (m *MockuserStoreInterface) EXPECT() *MockuserStoreInterfaceMockRecorder {
	return m.recorder
}

func (m *MockuserStoreInterface) CreateUser(user models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockuserStoreInterfaceMockRecorder) CreateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockuserStoreInterface)(nil).CreateUser), user)
}

func (m *MockuserStoreInterface) GetAllUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockuserStoreInterfaceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockuserStoreInterface)(nil).GetAllUsers))
}

func (m *MockuserStoreInterface) GetUserByID(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *MockuserStoreInterfaceMockRecorder) GetUserByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockuserStoreInterface)(nil).GetUserByID), id)
}
