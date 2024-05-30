// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"

	ssz "github.com/ferranbt/fastssz"

	types "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
)

// ReadOnlyBeaconBlock is an autogenerated mock type for the ReadOnlyBeaconBlock type
type ReadOnlyBeaconBlock[BodyT interface{}] struct {
	mock.Mock
}

type ReadOnlyBeaconBlock_Expecter[BodyT interface{}] struct {
	mock *mock.Mock
}

func (_m *ReadOnlyBeaconBlock[BodyT]) EXPECT() *ReadOnlyBeaconBlock_Expecter[BodyT] {
	return &ReadOnlyBeaconBlock_Expecter[BodyT]{mock: &_m.Mock}
}

// GetBody provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetBody() BodyT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 BodyT
	if rf, ok := ret.Get(0).(func() BodyT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(BodyT)
	}

	return r0
}

// ReadOnlyBeaconBlock_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type ReadOnlyBeaconBlock_GetBody_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetBody() *ReadOnlyBeaconBlock_GetBody_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetBody_Call[BodyT]{Call: _e.mock.On("GetBody")}
}

func (_c *ReadOnlyBeaconBlock_GetBody_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetBody_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetBody_Call[BodyT]) Return(_a0 BodyT) *ReadOnlyBeaconBlock_GetBody_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetBody_Call[BodyT]) RunAndReturn(run func() BodyT) *ReadOnlyBeaconBlock_GetBody_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetHeader provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetHeader() *types.BeaconBlockHeader {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHeader")
	}

	var r0 *types.BeaconBlockHeader
	if rf, ok := ret.Get(0).(func() *types.BeaconBlockHeader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BeaconBlockHeader)
		}
	}

	return r0
}

// ReadOnlyBeaconBlock_GetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeader'
type ReadOnlyBeaconBlock_GetHeader_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetHeader is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetHeader() *ReadOnlyBeaconBlock_GetHeader_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetHeader_Call[BodyT]{Call: _e.mock.On("GetHeader")}
}

func (_c *ReadOnlyBeaconBlock_GetHeader_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetHeader_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetHeader_Call[BodyT]) Return(_a0 *types.BeaconBlockHeader) *ReadOnlyBeaconBlock_GetHeader_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetHeader_Call[BodyT]) RunAndReturn(run func() *types.BeaconBlockHeader) *ReadOnlyBeaconBlock_GetHeader_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetParentBlockRoot provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetParentBlockRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetParentBlockRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// ReadOnlyBeaconBlock_GetParentBlockRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParentBlockRoot'
type ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetParentBlockRoot is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetParentBlockRoot() *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT]{Call: _e.mock.On("GetParentBlockRoot")}
}

func (_c *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT]) Return(_a0 bytes.B32) *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT]) RunAndReturn(run func() bytes.B32) *ReadOnlyBeaconBlock_GetParentBlockRoot_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetProposerIndex provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetProposerIndex() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposerIndex")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// ReadOnlyBeaconBlock_GetProposerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposerIndex'
type ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetProposerIndex is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetProposerIndex() *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT]{Call: _e.mock.On("GetProposerIndex")}
}

func (_c *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT]) Return(_a0 math.U64) *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT]) RunAndReturn(run func() math.U64) *ReadOnlyBeaconBlock_GetProposerIndex_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetSlot() math.U64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 math.U64
	if rf, ok := ret.Get(0).(func() math.U64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.U64)
	}

	return r0
}

// ReadOnlyBeaconBlock_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type ReadOnlyBeaconBlock_GetSlot_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetSlot() *ReadOnlyBeaconBlock_GetSlot_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetSlot_Call[BodyT]{Call: _e.mock.On("GetSlot")}
}

func (_c *ReadOnlyBeaconBlock_GetSlot_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetSlot_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetSlot_Call[BodyT]) Return(_a0 math.U64) *ReadOnlyBeaconBlock_GetSlot_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetSlot_Call[BodyT]) RunAndReturn(run func() math.U64) *ReadOnlyBeaconBlock_GetSlot_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetStateRoot provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetStateRoot() bytes.B32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStateRoot")
	}

	var r0 bytes.B32
	if rf, ok := ret.Get(0).(func() bytes.B32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bytes.B32)
		}
	}

	return r0
}

// ReadOnlyBeaconBlock_GetStateRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStateRoot'
type ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetStateRoot is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetStateRoot() *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT]{Call: _e.mock.On("GetStateRoot")}
}

func (_c *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT]) Return(_a0 bytes.B32) *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT]) RunAndReturn(run func() bytes.B32) *ReadOnlyBeaconBlock_GetStateRoot_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// GetTree provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) GetTree() (*ssz.Node, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTree")
	}

	var r0 *ssz.Node
	var r1 error
	if rf, ok := ret.Get(0).(func() (*ssz.Node, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *ssz.Node); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ssz.Node)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyBeaconBlock_GetTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTree'
type ReadOnlyBeaconBlock_GetTree_Call[BodyT interface{}] struct {
	*mock.Call
}

// GetTree is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) GetTree() *ReadOnlyBeaconBlock_GetTree_Call[BodyT] {
	return &ReadOnlyBeaconBlock_GetTree_Call[BodyT]{Call: _e.mock.On("GetTree")}
}

func (_c *ReadOnlyBeaconBlock_GetTree_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_GetTree_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetTree_Call[BodyT]) Return(_a0 *ssz.Node, _a1 error) *ReadOnlyBeaconBlock_GetTree_Call[BodyT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyBeaconBlock_GetTree_Call[BodyT]) RunAndReturn(run func() (*ssz.Node, error)) *ReadOnlyBeaconBlock_GetTree_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// HashTreeRoot provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) HashTreeRoot() ([32]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRoot")
	}

	var r0 [32]byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([32]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() [32]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyBeaconBlock_HashTreeRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRoot'
type ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT interface{}] struct {
	*mock.Call
}

// HashTreeRoot is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) HashTreeRoot() *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT] {
	return &ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT]{Call: _e.mock.On("HashTreeRoot")}
}

func (_c *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT]) Return(_a0 [32]byte, _a1 error) *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT]) RunAndReturn(run func() ([32]byte, error)) *ReadOnlyBeaconBlock_HashTreeRoot_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// HashTreeRootWith provides a mock function with given fields: hh
func (_m *ReadOnlyBeaconBlock[BodyT]) HashTreeRootWith(hh ssz.HashWalker) error {
	ret := _m.Called(hh)

	if len(ret) == 0 {
		panic("no return value specified for HashTreeRootWith")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ssz.HashWalker) error); ok {
		r0 = rf(hh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadOnlyBeaconBlock_HashTreeRootWith_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashTreeRootWith'
type ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT interface{}] struct {
	*mock.Call
}

// HashTreeRootWith is a helper method to define mock.On call
//   - hh ssz.HashWalker
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) HashTreeRootWith(hh interface{}) *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT] {
	return &ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT]{Call: _e.mock.On("HashTreeRootWith", hh)}
}

func (_c *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT]) Run(run func(hh ssz.HashWalker)) *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ssz.HashWalker))
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT]) Return(_a0 error) *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT]) RunAndReturn(run func(ssz.HashWalker) error) *ReadOnlyBeaconBlock_HashTreeRootWith_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// IsNil provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) IsNil() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsNil")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReadOnlyBeaconBlock_IsNil_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsNil'
type ReadOnlyBeaconBlock_IsNil_Call[BodyT interface{}] struct {
	*mock.Call
}

// IsNil is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) IsNil() *ReadOnlyBeaconBlock_IsNil_Call[BodyT] {
	return &ReadOnlyBeaconBlock_IsNil_Call[BodyT]{Call: _e.mock.On("IsNil")}
}

func (_c *ReadOnlyBeaconBlock_IsNil_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_IsNil_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_IsNil_Call[BodyT]) Return(_a0 bool) *ReadOnlyBeaconBlock_IsNil_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_IsNil_Call[BodyT]) RunAndReturn(run func() bool) *ReadOnlyBeaconBlock_IsNil_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZ provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) MarshalSSZ() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZ")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyBeaconBlock_MarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZ'
type ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT interface{}] struct {
	*mock.Call
}

// MarshalSSZ is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) MarshalSSZ() *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT] {
	return &ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT]{Call: _e.mock.On("MarshalSSZ")}
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT]) Return(_a0 []byte, _a1 error) *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT]) RunAndReturn(run func() ([]byte, error)) *ReadOnlyBeaconBlock_MarshalSSZ_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// MarshalSSZTo provides a mock function with given fields: dst
func (_m *ReadOnlyBeaconBlock[BodyT]) MarshalSSZTo(dst []byte) ([]byte, error) {
	ret := _m.Called(dst)

	if len(ret) == 0 {
		panic("no return value specified for MarshalSSZTo")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(dst)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(dst)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOnlyBeaconBlock_MarshalSSZTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalSSZTo'
type ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT interface{}] struct {
	*mock.Call
}

// MarshalSSZTo is a helper method to define mock.On call
//   - dst []byte
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) MarshalSSZTo(dst interface{}) *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT] {
	return &ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT]{Call: _e.mock.On("MarshalSSZTo", dst)}
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT]) Run(run func(dst []byte)) *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT]) Return(_a0 []byte, _a1 error) *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT]) RunAndReturn(run func([]byte) ([]byte, error)) *ReadOnlyBeaconBlock_MarshalSSZTo_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// SizeSSZ provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) SizeSSZ() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SizeSSZ")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// ReadOnlyBeaconBlock_SizeSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SizeSSZ'
type ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT interface{}] struct {
	*mock.Call
}

// SizeSSZ is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) SizeSSZ() *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT] {
	return &ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT]{Call: _e.mock.On("SizeSSZ")}
}

func (_c *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT]) Return(_a0 int) *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT]) RunAndReturn(run func() int) *ReadOnlyBeaconBlock_SizeSSZ_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// UnmarshalSSZ provides a mock function with given fields: buf
func (_m *ReadOnlyBeaconBlock[BodyT]) UnmarshalSSZ(buf []byte) error {
	ret := _m.Called(buf)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalSSZ")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(buf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadOnlyBeaconBlock_UnmarshalSSZ_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalSSZ'
type ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT interface{}] struct {
	*mock.Call
}

// UnmarshalSSZ is a helper method to define mock.On call
//   - buf []byte
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) UnmarshalSSZ(buf interface{}) *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT] {
	return &ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT]{Call: _e.mock.On("UnmarshalSSZ", buf)}
}

func (_c *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT]) Run(run func(buf []byte)) *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT]) Return(_a0 error) *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT]) RunAndReturn(run func([]byte) error) *ReadOnlyBeaconBlock_UnmarshalSSZ_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// Version provides a mock function with given fields:
func (_m *ReadOnlyBeaconBlock[BodyT]) Version() uint32 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// ReadOnlyBeaconBlock_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type ReadOnlyBeaconBlock_Version_Call[BodyT interface{}] struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
func (_e *ReadOnlyBeaconBlock_Expecter[BodyT]) Version() *ReadOnlyBeaconBlock_Version_Call[BodyT] {
	return &ReadOnlyBeaconBlock_Version_Call[BodyT]{Call: _e.mock.On("Version")}
}

func (_c *ReadOnlyBeaconBlock_Version_Call[BodyT]) Run(run func()) *ReadOnlyBeaconBlock_Version_Call[BodyT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ReadOnlyBeaconBlock_Version_Call[BodyT]) Return(_a0 uint32) *ReadOnlyBeaconBlock_Version_Call[BodyT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ReadOnlyBeaconBlock_Version_Call[BodyT]) RunAndReturn(run func() uint32) *ReadOnlyBeaconBlock_Version_Call[BodyT] {
	_c.Call.Return(run)
	return _c
}

// NewReadOnlyBeaconBlock creates a new instance of ReadOnlyBeaconBlock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReadOnlyBeaconBlock[BodyT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *ReadOnlyBeaconBlock[BodyT] {
	mock := &ReadOnlyBeaconBlock[BodyT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}