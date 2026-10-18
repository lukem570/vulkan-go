//go:build !cgo && !vkcgo

package ffi

import (
	"unsafe"

	goffi "github.com/go-webgpu/goffi/ffi"
	"github.com/go-webgpu/goffi/types"
)

// Backend names the linked foreign-call implementation.
const Backend = "goffi"

// CIF is a prepared call signature.
type CIF = types.CallInterface

var descriptors = [...]*types.TypeDescriptor{
	Void:    types.VoidTypeDescriptor,
	UInt8:   types.UInt8TypeDescriptor,
	SInt8:   types.SInt8TypeDescriptor,
	UInt16:  types.UInt16TypeDescriptor,
	SInt16:  types.SInt16TypeDescriptor,
	UInt32:  types.UInt32TypeDescriptor,
	SInt32:  types.SInt32TypeDescriptor,
	UInt64:  types.UInt64TypeDescriptor,
	SInt64:  types.SInt64TypeDescriptor,
	Float:   types.FloatTypeDescriptor,
	Double:  types.DoubleTypeDescriptor,
	Pointer: types.PointerTypeDescriptor,
}

// LoadLibrary opens a shared library by name or path.
func LoadLibrary(name string) (unsafe.Pointer, error) {
	return goffi.LoadLibrary(name)
}

// Symbol resolves a symbol in a library opened by LoadLibrary.
func Symbol(lib unsafe.Pointer, name string) (unsafe.Pointer, error) {
	return goffi.GetSymbol(lib, name)
}

// Prepare computes the call signature for a return type and argument list.
func Prepare(cif *CIF, ret Type, args []Type) error {
	argDescs := make([]*types.TypeDescriptor, len(args))
	for i, a := range args {
		argDescs[i] = descriptors[a]
	}
	if err := goffi.PrepareCallInterface(cif, types.DefaultCall, descriptors[ret], argDescs); err != nil {
		return err
	}
	if cif.ReturnType == nil {
		return errPrepFailed
	}
	return nil
}

// Call invokes fn through a prepared cif. Each element of args points to the
// storage holding that argument's value; ret points to storage for the return
// value, or is nil for a void call.
func Call(cif *CIF, fn, ret unsafe.Pointer, args []unsafe.Pointer) error {
	// goffi panics on a zero-valued interface; both backends must report this
	// the same way.
	if cif == nil || cif.ReturnType == nil {
		return errNotPrepared
	}
	if len(args) != cif.ArgCount {
		return errArgCount
	}
	return goffi.CallFunction(cif, fn, ret, args)
}

// NewCallback wraps a Go function as a C function pointer. The Go function's
// parameters and result must be C-representable scalars or unsafe.Pointer.
func NewCallback(fn any) uintptr {
	return goffi.NewCallback(fn)
}
