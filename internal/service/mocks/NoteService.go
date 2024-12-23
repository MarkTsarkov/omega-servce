// Code generated by mockery v2.48.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/marktsarkov/omega-service/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// NoteService is an autogenerated mock type for the NoteService type
type NoteService struct {
	mock.Mock
}

type NoteService_Expecter struct {
	mock *mock.Mock
}

func (_m *NoteService) EXPECT() *NoteService_Expecter {
	return &NoteService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, note
func (_m *NoteService) Create(ctx context.Context, note *entity.Note) (int64, error) {
	ret := _m.Called(ctx, note)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) (int64, error)); ok {
		return rf(ctx, note)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) int64); ok {
		r0 = rf(ctx, note)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Note) error); ok {
		r1 = rf(ctx, note)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NoteService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - note *entity.Note
func (_e *NoteService_Expecter) Create(ctx interface{}, note interface{}) *NoteService_Create_Call {
	return &NoteService_Create_Call{Call: _e.mock.On("Create", ctx, note)}
}

func (_c *NoteService_Create_Call) Run(run func(ctx context.Context, note *entity.Note)) *NoteService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Note))
	})
	return _c
}

func (_c *NoteService_Create_Call) Return(_a0 int64, _a1 error) *NoteService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_Create_Call) RunAndReturn(run func(context.Context, *entity.Note) (int64, error)) *NoteService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *NoteService) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *entity.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Note, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Note); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteService_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type NoteService_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *NoteService_Expecter) GetById(ctx interface{}, id interface{}) *NoteService_GetById_Call {
	return &NoteService_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *NoteService_GetById_Call) Run(run func(ctx context.Context, id int64)) *NoteService_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *NoteService_GetById_Call) Return(_a0 *entity.Note, _a1 error) *NoteService_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteService_GetById_Call) RunAndReturn(run func(context.Context, int64) (*entity.Note, error)) *NoteService_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// NewNoteService creates a new instance of NoteService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteService {
	mock := &NoteService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
