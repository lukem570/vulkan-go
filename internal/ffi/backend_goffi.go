//go:build !cgo && !vkcgo

package ffi

import (
	"reflect"
	"runtime"
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
	if runtime.GOOS == "windows" {
		fn = widenForSyscall(fn)
	}
	return goffi.NewCallback(fn)
}

// widenForSyscall adapts a callback for syscall.NewCallback, which goffi uses
// on Windows and which accepts only pointer-sized arguments and exactly one
// pointer-sized result. A Vulkan callback has neither: the debug messenger
// takes two uint32 flags, and the allocator's pfnFree returns nothing, so goffi
// panics on them at package initialization.
//
// The returned function has an all-uintptr signature. It narrows each argument
// back to the type the original expects and widens the result, which is sound
// on Win64 because integer arguments there occupy a full register regardless of
// their declared width.
func widenForSyscall(fn any) any {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return fn
	}
	t := v.Type()
	if t.IsVariadic() {
		return fn
	}

	word := reflect.TypeOf(uintptr(0))
	in := make([]reflect.Type, t.NumIn())
	for i := range in {
		in[i] = word
	}
	widened := reflect.FuncOf(in, []reflect.Type{word}, false)

	return reflect.MakeFunc(widened, func(args []reflect.Value) []reflect.Value {
		call := make([]reflect.Value, len(args))
		for i, a := range args {
			call[i] = narrowWord(uintptr(a.Uint()), t.In(i))
		}
		out := v.Call(call)
		if len(out) == 0 {
			return []reflect.Value{reflect.ValueOf(uintptr(0))}
		}
		return []reflect.Value{reflect.ValueOf(widenWord(out[0]))}
	}).Interface()
}

// narrowWord reinterprets one machine word as the parameter type want.
func narrowWord(word uintptr, want reflect.Type) reflect.Value {
	v := reflect.New(want).Elem()
	switch want.Kind() {
	case reflect.UnsafePointer, reflect.Ptr:
		// Reinterpret the bits; the word is a C address, not a Go pointer.
		return reflect.NewAt(want, unsafe.Pointer(&word)).Elem()
	case reflect.Bool:
		v.SetBool(word != 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(truncate(uint64(word), want.Bits())))
	default:
		v.SetUint(truncate(uint64(word), want.Bits()))
	}
	return v
}

// widenWord packs a callback result into one machine word.
func widenWord(v reflect.Value) uintptr {
	switch v.Kind() {
	case reflect.UnsafePointer, reflect.Ptr:
		return v.Pointer()
	case reflect.Bool:
		if v.Bool() {
			return 1
		}
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uintptr(v.Int())
	default:
		return uintptr(v.Uint())
	}
}

// truncate keeps the low bits bits of u, so narrowing never overflows the
// destination type. Win64 leaves the upper half of a register holding a 32-bit
// argument undefined.
func truncate(u uint64, bits int) uint64 {
	if bits >= 64 {
		return u
	}
	return u & (1<<uint(bits) - 1)
}
