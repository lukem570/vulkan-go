package ffi

import (
	"testing"
	"unsafe"
)

// codePtr reinterprets a function address returned by NewCallback as a
// pointer, the same way the generated bindings reinterpret handle values.
func codePtr(fn uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&fn))
}

// The tests below run against whichever backend is compiled in, so the same
// assertions cover goffi (CGO_ENABLED=0) and libffi/cgo (CGO_ENABLED=1).
//
// Each one builds a C function pointer with NewCallback and then invokes it
// through Prepare/Call, which exercises both halves of the seam: the Go → C
// call path and the C → Go callback path.

func TestBackendNamed(t *testing.T) {
	if Backend != "goffi" && Backend != "cgo" {
		t.Fatalf("unexpected backend %q", Backend)
	}
	t.Logf("backend: %s", Backend)
}

func TestCallbackUint32Result(t *testing.T) {
	fn := NewCallback(func(a uint32, b uint32) uint32 { return a*1000 + b })
	if fn == 0 {
		t.Fatal("NewCallback returned 0")
	}

	var cif CIF
	if err := Prepare(&cif, UInt32, []Type{UInt32, UInt32}); err != nil {
		t.Fatalf("Prepare: %v", err)
	}

	a, b := uint32(42), uint32(7)
	var ret uint32
	err := Call(&cif, codePtr(fn), unsafe.Pointer(&ret), []unsafe.Pointer{
		unsafe.Pointer(&a), unsafe.Pointer(&b),
	})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	if ret != 42007 {
		t.Fatalf("got %d, want 42007", ret)
	}
}

func TestCallbackPointerResult(t *testing.T) {
	// Mirrors the shape of VkAllocationCallbacks.pfnAllocation: pointer and
	// uintptr arguments, pointer result.
	var seen struct {
		user  unsafe.Pointer
		size  uintptr
		align uintptr
		scope uint32
	}
	target := make([]byte, 16)
	want := unsafe.Pointer(&target[0])

	fn := NewCallback(func(user unsafe.Pointer, size, align uintptr, scope uint32) unsafe.Pointer {
		seen.user, seen.size, seen.align, seen.scope = user, size, align, scope
		return want
	})

	var cif CIF
	if err := Prepare(&cif, Pointer, []Type{Pointer, Pointer, Pointer, UInt32}); err != nil {
		t.Fatalf("Prepare: %v", err)
	}

	userTag := make([]byte, 1)
	user := unsafe.Pointer(&userTag[0])
	size, align := uintptr(1024), uintptr(64)
	scope := uint32(3)
	var ret unsafe.Pointer
	err := Call(&cif, codePtr(fn), unsafe.Pointer(&ret), []unsafe.Pointer{
		unsafe.Pointer(&user), unsafe.Pointer(&size), unsafe.Pointer(&align), unsafe.Pointer(&scope),
	})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}

	if seen.user != user {
		t.Errorf("user data: got %p, want %p", seen.user, user)
	}
	if seen.size != size || seen.align != align {
		t.Errorf("size/align: got %d/%d, want %d/%d", seen.size, seen.align, size, align)
	}
	if seen.scope != scope {
		t.Errorf("scope: got %d, want %d", seen.scope, scope)
	}
	if ret != want {
		t.Errorf("result: got %p, want %p", ret, want)
	}
}

func TestCallbackVoidResult(t *testing.T) {
	var got uint64
	fn := NewCallback(func(_ unsafe.Pointer, v uint64) { got = v })

	var cif CIF
	if err := Prepare(&cif, Void, []Type{Pointer, UInt64}); err != nil {
		t.Fatalf("Prepare: %v", err)
	}

	var user unsafe.Pointer
	v := uint64(0xdeadbeefcafe)
	err := Call(&cif, codePtr(fn), nil, []unsafe.Pointer{
		unsafe.Pointer(&user), unsafe.Pointer(&v),
	})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	if got != v {
		t.Fatalf("got %#x, want %#x", got, v)
	}
}

func TestCallbackNarrowAndFloatArgs(t *testing.T) {
	if Backend == "goffi" {
		// goffi v0.5.0 delivers a float32 argument as 0 when it follows
		// integer arguments in a callback; libffi handles the same signature
		// correctly. No Vulkan callback mixes floats with integers, so this
		// only matters for other users of the seam.
		t.Skip("goffi mis-marshals float arguments mixed with integers")
	}
	fn := NewCallback(func(a uint8, b int16, c float32, d float64) float64 {
		return float64(a) + float64(b) + float64(c) + d
	})

	var cif CIF
	if err := Prepare(&cif, Double, []Type{UInt8, SInt16, Float, Double}); err != nil {
		t.Fatalf("Prepare: %v", err)
	}

	a, b, c, d := uint8(200), int16(-300), float32(0.5), 0.25
	var ret float64
	err := Call(&cif, codePtr(fn), unsafe.Pointer(&ret), []unsafe.Pointer{
		unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&d),
	})
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	if want := -99.25; ret != want {
		t.Fatalf("got %v, want %v", ret, want)
	}
}

func TestUnpreparedCallFails(t *testing.T) {
	var cif CIF
	fn := NewCallback(func() {})
	if err := Call(&cif, codePtr(fn), nil, nil); err == nil {
		t.Fatal("expected an error for an unprepared call interface")
	}

	// A wrong argument count is rejected rather than read off the end.
	if err := Prepare(&cif, Void, []Type{UInt32}); err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if err := Call(&cif, codePtr(fn), nil, nil); err == nil {
		t.Fatal("expected an error for a mismatched argument count")
	}
}
