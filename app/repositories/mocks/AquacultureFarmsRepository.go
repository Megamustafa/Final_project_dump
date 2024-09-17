// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	models "aquaculture/models"

	mock "github.com/stretchr/testify/mock"
)

// AquacultureFarmsRepository is an autogenerated mock type for the AquacultureFarmsRepository type
type AquacultureFarmsRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: afReq
func (_m *AquacultureFarmsRepository) Create(afReq models.AquacultureFarmsRequest) (models.AquacultureFarms, error) {
	ret := _m.Called(afReq)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 models.AquacultureFarms
	var r1 error
	if rf, ok := ret.Get(0).(func(models.AquacultureFarmsRequest) (models.AquacultureFarms, error)); ok {
		return rf(afReq)
	}
	if rf, ok := ret.Get(0).(func(models.AquacultureFarmsRequest) models.AquacultureFarms); ok {
		r0 = rf(afReq)
	} else {
		r0 = ret.Get(0).(models.AquacultureFarms)
	}

	if rf, ok := ret.Get(1).(func(models.AquacultureFarmsRequest) error); ok {
		r1 = rf(afReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *AquacultureFarmsRepository) Delete(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *AquacultureFarmsRepository) GetAll() ([]models.AquacultureFarms, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.AquacultureFarms
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.AquacultureFarms, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.AquacultureFarms); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AquacultureFarms)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *AquacultureFarmsRepository) GetByID(id string) (models.AquacultureFarms, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 models.AquacultureFarms
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.AquacultureFarms, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.AquacultureFarms); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.AquacultureFarms)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: afReq, id
func (_m *AquacultureFarmsRepository) Update(afReq models.AquacultureFarmsRequest, id string) (models.AquacultureFarms, error) {
	ret := _m.Called(afReq, id)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 models.AquacultureFarms
	var r1 error
	if rf, ok := ret.Get(0).(func(models.AquacultureFarmsRequest, string) (models.AquacultureFarms, error)); ok {
		return rf(afReq, id)
	}
	if rf, ok := ret.Get(0).(func(models.AquacultureFarmsRequest, string) models.AquacultureFarms); ok {
		r0 = rf(afReq, id)
	} else {
		r0 = ret.Get(0).(models.AquacultureFarms)
	}

	if rf, ok := ret.Get(1).(func(models.AquacultureFarmsRequest, string) error); ok {
		r1 = rf(afReq, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAquacultureFarmsRepository creates a new instance of AquacultureFarmsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAquacultureFarmsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AquacultureFarmsRepository {
	mock := &AquacultureFarmsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
