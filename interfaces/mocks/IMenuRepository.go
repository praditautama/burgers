// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/opam22/burgers/models"

// IMenuRepository is an autogenerated mock type for the IMenuRepository type
type IMenuRepository struct {
	mock.Mock
}

// FetchAllMenu provides a mock function with given fields:
func (_m *IMenuRepository) FetchAllMenu() ([]models.Menu, error) {
	ret := _m.Called()

	var r0 []models.Menu
	if rf, ok := ret.Get(0).(func() []models.Menu); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Menu)
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

// FetchMenu provides a mock function with given fields: idMenu
func (_m *IMenuRepository) FetchMenu(idMenu int) (models.Menu, error) {
	ret := _m.Called(idMenu)

	var r0 models.Menu
	if rf, ok := ret.Get(0).(func(int) models.Menu); ok {
		r0 = rf(idMenu)
	} else {
		r0 = ret.Get(0).(models.Menu)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idMenu)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchMenuIngredients provides a mock function with given fields: idMenu
func (_m *IMenuRepository) FetchMenuIngredients(idMenu int) ([]models.MenuIngredient, error) {
	ret := _m.Called(idMenu)

	var r0 []models.MenuIngredient
	if rf, ok := ret.Get(0).(func(int) []models.MenuIngredient); ok {
		r0 = rf(idMenu)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.MenuIngredient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idMenu)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
