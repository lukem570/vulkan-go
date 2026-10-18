// Package ffi is the foreign-call seam used by the generated bindings.
//
// The generated code never names an FFI library directly. It declares one CIF
// per Vulkan command, prepares it with Prepare, and invokes it with Call, so
// the whole binding can be retargeted by swapping the backend that implements
// those five primitives.
//
// Two backends exist, selected at compile time by the cgo build constraint:
//
//	CGO_ENABLED=0  →  goffi (pure Go, no C toolchain)
//	CGO_ENABLED=1  →  libffi through cgo
//
// Nothing needs configuring: goffi cannot be built with cgo enabled (it relies
// on cgo_import_dynamic, which requires CGO_ENABLED=0), and a cgo file cannot
// be built with cgo disabled, so the two backends are mutually exclusive by
// construction. Build with -tags vkcgo to assert the cgo backend and fail the
// build if cgo is unavailable, rather than silently getting goffi.
//
// Backend reports which one is linked in.
package ffi

import "errors"

// Errors shared by both backends.
var (
	errNotPrepared = errors.New("ffi: call interface not prepared")
	errPrepFailed  = errors.New("ffi: preparing the call interface failed")
	errArgCount    = errors.New("ffi: wrong number of arguments for prepared call")
)

// Type identifies a C primitive in a call signature. It is a plain code rather
// than a backend type descriptor so the generated bindings stay independent of
// whichever backend is compiled in.
type Type uint8

const (
	Void Type = iota
	UInt8
	SInt8
	UInt16
	SInt16
	UInt32
	SInt32
	UInt64
	SInt64
	Float
	Double
	Pointer
)

// size returns the width of t in bytes.
func (t Type) size() uintptr {
	switch t {
	case Void:
		return 0
	case UInt8, SInt8:
		return 1
	case UInt16, SInt16:
		return 2
	case UInt32, SInt32, Float:
		return 4
	default: // UInt64, SInt64, Double, Pointer
		return 8
	}
}

func (t Type) String() string {
	switch t {
	case Void:
		return "void"
	case UInt8:
		return "uint8"
	case SInt8:
		return "sint8"
	case UInt16:
		return "uint16"
	case SInt16:
		return "sint16"
	case UInt32:
		return "uint32"
	case SInt32:
		return "sint32"
	case UInt64:
		return "uint64"
	case SInt64:
		return "sint64"
	case Float:
		return "float"
	case Double:
		return "double"
	case Pointer:
		return "pointer"
	}
	return "unknown"
}
