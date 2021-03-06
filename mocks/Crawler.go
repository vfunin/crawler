// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	crawler "github.com/vfunin/crawler/internal/crawler"
)

// Crawler is an autogenerated mock type for the Crawler type
type Crawler struct {
	mock.Mock
}

// Crawl provides a mock function with given fields: ctx, cancel, url, withPanic, depth, errCh
func (_m *Crawler) Crawl(ctx context.Context, cancel context.CancelFunc, url string, withPanic bool, depth uint64, errCh chan<- error) {
	_m.Called(ctx, cancel, url, withPanic, depth, errCh)
}

// DecCnt provides a mock function with given fields:
func (_m *Crawler) DecCnt() {
	_m.Called()
}

// GetCnt provides a mock function with given fields:
func (_m *Crawler) GetCnt() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// IncCnt provides a mock function with given fields:
func (_m *Crawler) IncCnt() {
	_m.Called()
}

// IncMaxDepth provides a mock function with given fields: step
func (_m *Crawler) IncMaxDepth(step uint64) {
	_m.Called(step)
}

// MaxDepth provides a mock function with given fields:
func (_m *Crawler) MaxDepth() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ResultCh provides a mock function with given fields:
func (_m *Crawler) ResultCh() chan crawler.Result {
	ret := _m.Called()

	var r0 chan crawler.Result
	if rf, ok := ret.Get(0).(func() chan crawler.Result); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan crawler.Result)
		}
	}

	return r0
}
