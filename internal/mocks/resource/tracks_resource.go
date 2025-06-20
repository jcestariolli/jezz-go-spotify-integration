// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	model "jezz-go-spotify-integration/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// TracksResource is an autogenerated mock type for the TracksResource type
type TracksResource struct {
	mock.Mock
}

// GetTrack provides a mock function with given fields: accessToken, market, trackID
func (_m *TracksResource) GetTrack(accessToken model.AccessToken, market *model.AvailableMarket, trackID model.ID) (model.Track, error) {
	ret := _m.Called(accessToken, market, trackID)

	if len(ret) == 0 {
		panic("no return value specified for GetTrack")
	}

	var r0 model.Track
	var r1 error
	if rf, ok := ret.Get(0).(func(model.AccessToken, *model.AvailableMarket, model.ID) (model.Track, error)); ok {
		return rf(accessToken, market, trackID)
	}
	if rf, ok := ret.Get(0).(func(model.AccessToken, *model.AvailableMarket, model.ID) model.Track); ok {
		r0 = rf(accessToken, market, trackID)
	} else {
		r0 = ret.Get(0).(model.Track)
	}

	if rf, ok := ret.Get(1).(func(model.AccessToken, *model.AvailableMarket, model.ID) error); ok {
		r1 = rf(accessToken, market, trackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTracks provides a mock function with given fields: accessToken, market, tracksIDs
func (_m *TracksResource) GetTracks(accessToken model.AccessToken, market *model.AvailableMarket, tracksIDs model.TracksIDs) ([]model.Track, error) {
	ret := _m.Called(accessToken, market, tracksIDs)

	if len(ret) == 0 {
		panic("no return value specified for GetTracks")
	}

	var r0 []model.Track
	var r1 error
	if rf, ok := ret.Get(0).(func(model.AccessToken, *model.AvailableMarket, model.TracksIDs) ([]model.Track, error)); ok {
		return rf(accessToken, market, tracksIDs)
	}
	if rf, ok := ret.Get(0).(func(model.AccessToken, *model.AvailableMarket, model.TracksIDs) []model.Track); ok {
		r0 = rf(accessToken, market, tracksIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Track)
		}
	}

	if rf, ok := ret.Get(1).(func(model.AccessToken, *model.AvailableMarket, model.TracksIDs) error); ok {
		r1 = rf(accessToken, market, tracksIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTracksResource creates a new instance of TracksResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTracksResource(t interface {
	mock.TestingT
	Cleanup(func())
}) *TracksResource {
	mock := &TracksResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
