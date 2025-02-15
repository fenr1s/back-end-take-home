// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/fenr1s/back-end-take-home/domain/models"

// Airporter is an autogenerated mock type for the Airporter type
type Airporter struct {
	mock.Mock
}

// CheckExistance provides a mock function with given fields: airports, iata3
func (_m *Airporter) CheckExistance(airports []*models.Airport, iata3 string) (bool, error) {
	ret := _m.Called(airports, iata3)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]*models.Airport, string) bool); ok {
		r0 = rf(airports, iata3)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*models.Airport, string) error); ok {
		r1 = rf(airports, iata3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAirports provides a mock function with given fields:
func (_m *Airporter) GetAirports() ([]*models.Airport, error) {
	ret := _m.Called()

	var r0 []*models.Airport
	if rf, ok := ret.Get(0).(func() []*models.Airport); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Airport)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
