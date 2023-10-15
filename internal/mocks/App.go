// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	ads "homework10/internal/ads"

	mock "github.com/stretchr/testify/mock"

	time "time"

	users "homework10/internal/users"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// ChangeAdStatus provides a mock function with given fields: adId, userId, published
func (_m *App) ChangeAdStatus(adId int64, userId int64, published bool) (ads.Ad, error) {
	ret := _m.Called(adId, userId, published)

	var r0 ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, bool) (ads.Ad, error)); ok {
		return rf(adId, userId, published)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, bool) ads.Ad); ok {
		r0 = rf(adId, userId, published)
	} else {
		r0 = ret.Get(0).(ads.Ad)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, bool) error); ok {
		r1 = rf(adId, userId, published)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAd provides a mock function with given fields: title, text, userId
func (_m *App) CreateAd(title string, text string, userId int64) (ads.Ad, error) {
	ret := _m.Called(title, text, userId)

	var r0 ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) (ads.Ad, error)); ok {
		return rf(title, text, userId)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) ads.Ad); ok {
		r0 = rf(title, text, userId)
	} else {
		r0 = ret.Get(0).(ads.Ad)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(title, text, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: name, email
func (_m *App) CreateUser(name string, email string) (users.User, error) {
	ret := _m.Called(name, email)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (users.User, error)); ok {
		return rf(name, email)
	}
	if rf, ok := ret.Get(0).(func(string, string) users.User); ok {
		r0 = rf(name, email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAd provides a mock function with given fields: adId, userId
func (_m *App) DeleteAd(adId int64, userId int64) error {
	ret := _m.Called(adId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(adId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: userId
func (_m *App) DeleteUser(userId int64) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAd provides a mock function with given fields: adId
func (_m *App) GetAd(adId int64) (ads.Ad, error) {
	ret := _m.Called(adId)

	var r0 ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (ads.Ad, error)); ok {
		return rf(adId)
	}
	if rf, ok := ret.Get(0).(func(int64) ads.Ad); ok {
		r0 = rf(adId)
	} else {
		r0 = ret.Get(0).(ads.Ad)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(adId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userId
func (_m *App) GetUser(userId int64) (users.User, error) {
	ret := _m.Called(userId)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (users.User, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int64) users.User); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAds provides a mock function with given fields: pubFilter, userFilter, timeFilter
func (_m *App) ListAds(pubFilter bool, userFilter int64, timeFilter time.Time) ([]ads.Ad, error) {
	ret := _m.Called(pubFilter, userFilter, timeFilter)

	var r0 []ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(bool, int64, time.Time) ([]ads.Ad, error)); ok {
		return rf(pubFilter, userFilter, timeFilter)
	}
	if rf, ok := ret.Get(0).(func(bool, int64, time.Time) []ads.Ad); ok {
		r0 = rf(pubFilter, userFilter, timeFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(bool, int64, time.Time) error); ok {
		r1 = rf(pubFilter, userFilter, timeFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchAds provides a mock function with given fields: pattern
func (_m *App) SearchAds(pattern string) ([]ads.Ad, error) {
	ret := _m.Called(pattern)

	var r0 []ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]ads.Ad, error)); ok {
		return rf(pattern)
	}
	if rf, ok := ret.Get(0).(func(string) []ads.Ad); ok {
		r0 = rf(pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAd provides a mock function with given fields: adId, userId, title, text
func (_m *App) UpdateAd(adId int64, userId int64, title string, text string) (ads.Ad, error) {
	ret := _m.Called(adId, userId, title, text)

	var r0 ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, string, string) (ads.Ad, error)); ok {
		return rf(adId, userId, title, text)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, string, string) ads.Ad); ok {
		r0 = rf(adId, userId, title, text)
	} else {
		r0 = ret.Get(0).(ads.Ad)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, string, string) error); ok {
		r1 = rf(adId, userId, title, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: userId, name, email
func (_m *App) UpdateUser(userId int64, name string, email string) (users.User, error) {
	ret := _m.Called(userId, name, email)

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, string) (users.User, error)); ok {
		return rf(userId, name, email)
	}
	if rf, ok := ret.Get(0).(func(int64, string, string) users.User); ok {
		r0 = rf(userId, name, email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(userId, name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewApp creates a new instance of App. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApp(t mockConstructorTestingTNewApp) *App {
	mock := &App{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
