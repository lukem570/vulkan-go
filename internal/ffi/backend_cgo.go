//go:build cgo

package ffi

/*
#cgo linux   LDFLAGS: -lffi -ldl
#cgo darwin  LDFLAGS: -lffi
#cgo freebsd LDFLAGS: -lffi
#cgo windows LDFLAGS: -lffi

#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include <ffi.h>

#ifdef _WIN32
#include <windows.h>
static void *vkffi_dlopen(const char *name) { return (void *)LoadLibraryA(name); }
static void *vkffi_dlsym(void *lib, const char *name) { return (void *)GetProcAddress((HMODULE)lib, name); }
static const char *vkffi_dlerror(void) { return "LoadLibrary/GetProcAddress failed"; }
#else
#include <dlfcn.h>
static void *vkffi_dlopen(const char *name) { return dlopen(name, RTLD_NOW | RTLD_LOCAL); }
static void *vkffi_dlsym(void *lib, const char *name) { return dlsym(lib, name); }
static const char *vkffi_dlerror(void) { const char *e = dlerror(); return e ? e : "unknown error"; }
#endif

// vkffi_cif owns the prepared signature. ffi_prep_cif keeps the argument type
// array alive inside the cif, so that array must be C memory, not Go memory.
typedef struct {
	ffi_cif    cif;
	ffi_type **atypes;
} vkffi_cif;

static vkffi_cif *vkffi_prep(ffi_type *ret, ffi_type **atypes, unsigned nargs) {
	vkffi_cif *c = (vkffi_cif *)calloc(1, sizeof(vkffi_cif));
	if (c == NULL) {
		return NULL;
	}
	c->atypes = atypes;
	if (ffi_prep_cif(&c->cif, FFI_DEFAULT_ABI, nargs, ret, atypes) != FFI_OK) {
		free(c);
		return NULL;
	}
	return c;
}

static void vkffi_call(vkffi_cif *c, void *fn, void *ret, void **avalue) {
	ffi_call(&c->cif, (void (*)(void))fn, ret, avalue);
}

// Closures: libffi calls vkffi_closure_entry on the foreign thread, which hands
// the raw argument array to Go.
extern void vkffiClosureDispatch(uintptr_t user, void *ret, void **args);

static void vkffi_closure_entry(ffi_cif *cif, void *ret, void **args, void *user) {
	(void)cif;
	vkffiClosureDispatch((uintptr_t)user, ret, args);
}

static void *vkffi_closure_new(ffi_type *ret, ffi_type **atypes, unsigned nargs, uintptr_t user) {
	void *code = NULL;
	ffi_closure *cl = ffi_closure_alloc(sizeof(ffi_closure), &code);
	if (cl == NULL) {
		return NULL;
	}
	ffi_cif *cif = (ffi_cif *)calloc(1, sizeof(ffi_cif));
	if (cif == NULL) {
		ffi_closure_free(cl);
		return NULL;
	}
	if (ffi_prep_cif(cif, FFI_DEFAULT_ABI, nargs, ret, atypes) != FFI_OK) {
		free(cif);
		ffi_closure_free(cl);
		return NULL;
	}
	if (ffi_prep_closure_loc(cl, cif, vkffi_closure_entry, (void *)user, code) != FFI_OK) {
		free(cif);
		ffi_closure_free(cl);
		return NULL;
	}
	return code;
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"unsafe"
)

// Backend names the linked foreign-call implementation.
const Backend = "cgo"

// CIF is a prepared call signature. The zero value is unprepared.
//
// Argument widths are kept on the Go side: Call copies each argument's bytes
// into C scratch memory, so that the void* array handed to libffi contains only
// C pointers. Passing a Go-allocated array of Go pointers to C would violate
// the cgo pointer rules.
type CIF struct {
	c        unsafe.Pointer // *C.vkffi_cif
	argSizes []uintptr
	retSize  uintptr
}

func ffiType(t Type) *C.ffi_type {
	switch t {
	case Void:
		return &C.ffi_type_void
	case UInt8:
		return &C.ffi_type_uint8
	case SInt8:
		return &C.ffi_type_sint8
	case UInt16:
		return &C.ffi_type_uint16
	case SInt16:
		return &C.ffi_type_sint16
	case UInt32:
		return &C.ffi_type_uint32
	case SInt32:
		return &C.ffi_type_sint32
	case UInt64:
		return &C.ffi_type_uint64
	case SInt64:
		return &C.ffi_type_sint64
	case Float:
		return &C.ffi_type_float
	case Double:
		return &C.ffi_type_double
	default:
		return &C.ffi_type_pointer
	}
}

// cTypeArray allocates a C-owned ffi_type* array. ffi_prep_cif retains it, so
// it is never freed: one array per command, allocated once at Initialize.
func cTypeArray(args []Type) **C.ffi_type {
	if len(args) == 0 {
		return nil
	}
	n := C.size_t(len(args)) * C.size_t(unsafe.Sizeof(uintptr(0)))
	arr := (**C.ffi_type)(C.malloc(n))
	slots := unsafe.Slice(arr, len(args))
	for i, a := range args {
		slots[i] = ffiType(a)
	}
	return arr
}

// LoadLibrary opens a shared library by name or path.
func LoadLibrary(name string) (unsafe.Pointer, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	lib := C.vkffi_dlopen(cname)
	if lib == nil {
		return nil, fmt.Errorf("ffi: cannot load %s: %s", name, C.GoString(C.vkffi_dlerror()))
	}
	return lib, nil
}

// Symbol resolves a symbol in a library opened by LoadLibrary.
func Symbol(lib unsafe.Pointer, name string) (unsafe.Pointer, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	sym := C.vkffi_dlsym(lib, cname)
	if sym == nil {
		return nil, fmt.Errorf("ffi: symbol %s not found: %s", name, C.GoString(C.vkffi_dlerror()))
	}
	return sym, nil
}

// Prepare computes the call signature for a return type and argument list.
func Prepare(cif *CIF, ret Type, args []Type) error {
	if cif == nil {
		return errNotPrepared
	}

	c := C.vkffi_prep(ffiType(ret), cTypeArray(args), C.unsigned(len(args)))
	if c == nil {
		return errPrepFailed
	}

	sizes := make([]uintptr, len(args))
	for i, a := range args {
		sizes[i] = a.size()
	}
	cif.c = unsafe.Pointer(c)
	cif.argSizes = sizes
	cif.retSize = ret.size()
	return nil
}

// maxScratchArgs bounds the pooled scratch block. Vulkan commands pass only
// scalars and pointers, and none takes more than a handful of arguments.
const maxScratchArgs = 16

// scratch is one C block holding the void* argument array, a slot per argument
// value, and a return slot. The return slot is a full ffi_arg: libffi widens
// integral returns narrower than that.
type scratch struct {
	block unsafe.Pointer
}

const (
	wordSize    = unsafe.Sizeof(uintptr(0))
	scratchSize = C.size_t(maxScratchArgs*wordSize /* avalue */ + maxScratchArgs*8 /* slots */ + 16 /* ret */)
)

var scratchPool = sync.Pool{
	New: func() any { return &scratch{block: C.malloc(scratchSize)} },
}

// Call invokes fn through a prepared cif. Each element of args points to the
// storage holding that argument's value; ret points to storage for the return
// value, or is nil for a void call.
func Call(cif *CIF, fn, ret unsafe.Pointer, args []unsafe.Pointer) error {
	if cif == nil || cif.c == nil {
		return errNotPrepared
	}
	if len(args) != len(cif.argSizes) {
		return errArgCount
	}
	if len(args) > maxScratchArgs {
		return fmt.Errorf("ffi: %d arguments exceeds the %d supported", len(args), maxScratchArgs)
	}

	s := scratchPool.Get().(*scratch)
	defer scratchPool.Put(s)

	avalue := unsafe.Slice((*unsafe.Pointer)(s.block), maxScratchArgs)
	slots := unsafe.Pointer(uintptr(s.block) + maxScratchArgs*wordSize)
	retSlot := unsafe.Pointer(uintptr(slots) + maxScratchArgs*8)

	// Copy each argument's bytes from Go storage into the C slot. Done in Go
	// so no Go pointer is ever handed to C as part of the argument vector.
	for i, size := range cif.argSizes {
		slot := unsafe.Pointer(uintptr(slots) + uintptr(i)*8)
		src := args[i]
		if src == nil {
			*(*uint64)(slot) = 0
		} else {
			switch size {
			case 1:
				*(*uint8)(slot) = *(*uint8)(src)
			case 2:
				*(*uint16)(slot) = *(*uint16)(src)
			case 4:
				*(*uint32)(slot) = *(*uint32)(src)
			default:
				*(*uint64)(slot) = *(*uint64)(src)
			}
		}
		avalue[i] = slot
	}

	*(*uint64)(retSlot) = 0
	*(*uint64)(unsafe.Pointer(uintptr(retSlot) + 8)) = 0

	C.vkffi_call((*C.vkffi_cif)(cif.c), fn, retSlot, (*unsafe.Pointer)(s.block))

	if ret == nil || cif.retSize == 0 {
		return nil
	}
	switch cif.retSize {
	case 1:
		*(*uint8)(ret) = *(*uint8)(retSlot)
	case 2:
		*(*uint16)(ret) = *(*uint16)(retSlot)
	case 4:
		*(*uint32)(ret) = *(*uint32)(retSlot)
	default:
		*(*uint64)(ret) = *(*uint64)(retSlot)
	}
	return nil
}

// ---------------------------------------------------------------------------
// Callbacks
// ---------------------------------------------------------------------------

// callback is a Go function reachable from a libffi closure.
type callback struct {
	fn     reflect.Value
	params []Type
	ret    Type
}

var (
	callbackMu   sync.RWMutex
	callbacks    = map[uintptr]*callback{}
	callbackNext atomic.Uintptr
)

// NewCallback wraps a Go function as a C function pointer, using a libffi
// closure. The Go function's parameters and result must be C-representable
// scalars or unsafe.Pointer. The closure is never released: the generated
// trampolines are package-level and live for the life of the process.
func NewCallback(fn any) uintptr {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic("ffi: NewCallback requires a function")
	}
	t := v.Type()
	if t.IsVariadic() {
		panic("ffi: NewCallback does not support variadic functions")
	}
	if t.NumOut() > 1 {
		panic("ffi: NewCallback requires at most one result")
	}

	cb := &callback{fn: v, ret: Void}
	for i := 0; i < t.NumIn(); i++ {
		cb.params = append(cb.params, typeForKind(t.In(i)))
	}
	if t.NumOut() == 1 {
		cb.ret = typeForKind(t.Out(0))
	}

	id := callbackNext.Add(1)
	callbackMu.Lock()
	callbacks[id] = cb
	callbackMu.Unlock()

	code := C.vkffi_closure_new(ffiType(cb.ret), cTypeArray(cb.params), C.unsigned(len(cb.params)), C.uintptr_t(id))
	if code == nil {
		panic("ffi: ffi_prep_closure_loc failed")
	}
	return uintptr(code)
}

// typeForKind maps a Go type in a callback signature to its C counterpart.
func typeForKind(t reflect.Type) Type {
	switch t.Kind() {
	case reflect.UnsafePointer, reflect.Uintptr, reflect.Ptr:
		return Pointer
	case reflect.Uint8:
		return UInt8
	case reflect.Int8:
		return SInt8
	case reflect.Uint16:
		return UInt16
	case reflect.Int16:
		return SInt16
	case reflect.Uint32:
		return UInt32
	case reflect.Int32:
		return SInt32
	case reflect.Uint64, reflect.Uint:
		return UInt64
	case reflect.Int64, reflect.Int:
		return SInt64
	case reflect.Float32:
		return Float
	case reflect.Float64:
		return Double
	case reflect.Bool:
		return UInt32 // VkBool32
	default:
		panic("ffi: unsupported callback type " + t.String())
	}
}

//export vkffiClosureDispatch
func vkffiClosureDispatch(user C.uintptr_t, ret unsafe.Pointer, args *unsafe.Pointer) {
	callbackMu.RLock()
	cb := callbacks[uintptr(user)]
	callbackMu.RUnlock()
	if cb == nil {
		return
	}

	in := make([]reflect.Value, len(cb.params))
	argv := unsafe.Slice(args, len(cb.params))
	for i, p := range cb.params {
		in[i] = readArg(cb.fn.Type().In(i), p, argv[i])
	}

	out := cb.fn.Call(in)
	if len(out) == 0 || ret == nil {
		return
	}
	writeResult(ret, cb.ret, out[0])
}

// readArg reconstructs the i-th Go argument from the slot libffi filled in.
func readArg(goType reflect.Type, c Type, slot unsafe.Pointer) reflect.Value {
	v := reflect.New(goType).Elem()
	switch c {
	case UInt8:
		setUint(v, uint64(*(*uint8)(slot)))
	case SInt8:
		setInt(v, int64(*(*int8)(slot)))
	case UInt16:
		setUint(v, uint64(*(*uint16)(slot)))
	case SInt16:
		setInt(v, int64(*(*int16)(slot)))
	case UInt32:
		setUint(v, uint64(*(*uint32)(slot)))
	case SInt32:
		setInt(v, int64(*(*int32)(slot)))
	case UInt64:
		setUint(v, *(*uint64)(slot))
	case SInt64:
		setInt(v, *(*int64)(slot))
	case Float:
		v.SetFloat(float64(*(*float32)(slot)))
	case Double:
		v.SetFloat(*(*float64)(slot))
	default: // Pointer
		p := *(*unsafe.Pointer)(slot)
		if goType.Kind() == reflect.Uintptr {
			v.SetUint(uint64(uintptr(p)))
		} else {
			v = reflect.NewAt(goType, unsafe.Pointer(&p)).Elem()
		}
	}
	return v
}

func setUint(v reflect.Value, u uint64) {
	if v.Kind() == reflect.Bool {
		v.SetBool(u != 0)
		return
	}
	if v.CanUint() {
		v.SetUint(u)
		return
	}
	v.SetInt(int64(u))
}

func setInt(v reflect.Value, i int64) {
	if v.CanInt() {
		v.SetInt(i)
		return
	}
	v.SetUint(uint64(i))
}

// writeResult stores a callback's result where libffi expects it. Integral
// results narrower than ffi_arg are widened, as the closure API requires.
func writeResult(ret unsafe.Pointer, c Type, v reflect.Value) {
	switch c {
	case Float:
		*(*float32)(ret) = float32(v.Float())
	case Double:
		*(*float64)(ret) = v.Float()
	case Pointer:
		if v.Kind() == reflect.Uintptr {
			*(*uintptr)(ret) = uintptr(v.Uint())
		} else {
			*(*unsafe.Pointer)(ret) = v.UnsafePointer()
		}
	default:
		var u uint64
		switch {
		case v.Kind() == reflect.Bool:
			if v.Bool() {
				u = 1
			}
		case v.CanUint():
			u = v.Uint()
		case v.CanInt():
			u = uint64(v.Int())
		}
		*(*uint64)(ret) = u
	}
}
