// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	parser "github.com/vfunin/crawler/internal/parser"
)

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, url
func (_m *Fetcher) Fetch(ctx context.Context, url string) (parser.Page, error) {
	ret := _m.Called(ctx, url)

	var r0 parser.Page
	if rf, ok := ret.Get(0).(func(context.Context, string) parser.Page); ok {
		r0 = rf(ctx, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(parser.Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}