//go:build !cgo && !vkcgo

package ffi

import (
	"reflect"
	"testing"
	"unsafe"
)

// widenForSyscall only takes effect on Windows, but its conversions are
// platform-independent, so they are exercised directly here. The Windows
// integration itself is covered by CI.

func TestWidenForSyscallMessengerShape(t *testing.T) {
	var got struct {
		sev, types uint32
		data, user unsafe.Pointer
	}
	fn := func(sev, types uint32, data, user unsafe.Pointer) uint32 {
		got.sev, got.types, got.data, got.user = sev, types, data, user
		return 0
	}

	widened := reflect.ValueOf(widenForSyscall(fn))
	if n := widened.Type().NumIn(); n != 4 {
		t.Fatalf("widened signature takes %d arguments, want 4", n)
	}
	for i := 0; i < widened.Type().NumIn(); i++ {
		if k := widened.Type().In(i).Kind(); k != reflect.Uintptr {
			t.Errorf("argument %d is %s, want uintptr", i, k)
		}
	}
	if n := widened.Type().NumOut(); n != 1 {
		t.Fatalf("widened signature returns %d values, want exactly 1", n)
	}

	// Win64 passes a 32-bit argument in a full register and leaves the upper
	// half undefined, so the adapter must ignore the high bits.
	tag := make([]byte, 1)
	user := unsafe.Pointer(&tag[0])
	out := widened.Call([]reflect.Value{
		reflect.ValueOf(uintptr(0xdeadbeef_00000010)),
		reflect.ValueOf(uintptr(0xffffffff_00000002)),
		reflect.ValueOf(uintptr(0)),
		reflect.ValueOf(uintptr(0)),
	})

	if got.sev != 0x10 {
		t.Errorf("severity: got %#x, want 0x10", got.sev)
	}
	if got.types != 0x2 {
		t.Errorf("type: got %#x, want 0x2", got.types)
	}
	if out[0].Kind() != reflect.Uintptr || out[0].Uint() != 0 {
		t.Errorf("result: got %v, want uintptr 0", out[0])
	}
	_ = user
}

func TestWidenForSyscallPointerRoundTrip(t *testing.T) {
	target := make([]byte, 8)
	addr := unsafe.Pointer(&target[0])

	var seen unsafe.Pointer
	fn := func(p unsafe.Pointer) unsafe.Pointer {
		seen = p
		return p
	}

	widened := reflect.ValueOf(widenForSyscall(fn))
	out := widened.Call([]reflect.Value{reflect.ValueOf(uintptr(addr))})

	if seen != addr {
		t.Errorf("argument: got %p, want %p", seen, addr)
	}
	if uintptr(out[0].Uint()) != uintptr(addr) {
		t.Errorf("result: got %#x, want %#x", out[0].Uint(), uintptr(addr))
	}
}

func TestWidenForSyscallVoidResult(t *testing.T) {
	// VkAllocationCallbacks.pfnFree returns nothing, which syscall.NewCallback
	// rejects outright; the adapter has to invent a result.
	called := false
	fn := func(_ unsafe.Pointer, _ unsafe.Pointer) { called = true }

	widened := reflect.ValueOf(widenForSyscall(fn))
	if n := widened.Type().NumOut(); n != 1 {
		t.Fatalf("widened signature returns %d values, want exactly 1", n)
	}
	out := widened.Call([]reflect.Value{reflect.ValueOf(uintptr(0)), reflect.ValueOf(uintptr(0))})

	if !called {
		t.Error("original function was not called")
	}
	if out[0].Uint() != 0 {
		t.Errorf("result: got %d, want 0", out[0].Uint())
	}
}

func TestWidenForSyscallNarrowKinds(t *testing.T) {
	var seen struct {
		b    bool
		i32  int32
		u8   uint8
		u64  uint64
		size uintptr
	}
	fn := func(b bool, i32 int32, u8 uint8, u64 uint64, size uintptr) bool {
		seen.b, seen.i32, seen.u8, seen.u64, seen.size = b, i32, u8, u64, size
		return b
	}

	widened := reflect.ValueOf(widenForSyscall(fn))
	out := widened.Call([]reflect.Value{
		reflect.ValueOf(uintptr(1)),
		reflect.ValueOf(uintptr(0xffffffff)), // -1 once narrowed to int32
		reflect.ValueOf(uintptr(0xff01)),     // 0x01 once narrowed to uint8
		reflect.ValueOf(uintptr(0x1234)),
		reflect.ValueOf(uintptr(64)),
	})

	if !seen.b {
		t.Error("bool: got false, want true")
	}
	if seen.i32 != -1 {
		t.Errorf("int32: got %d, want -1", seen.i32)
	}
	if seen.u8 != 0x01 {
		t.Errorf("uint8: got %#x, want 0x01", seen.u8)
	}
	if seen.u64 != 0x1234 {
		t.Errorf("uint64: got %#x, want 0x1234", seen.u64)
	}
	if seen.size != 64 {
		t.Errorf("uintptr: got %d, want 64", seen.size)
	}
	if out[0].Uint() != 1 {
		t.Errorf("bool result: got %d, want 1", out[0].Uint())
	}
}
