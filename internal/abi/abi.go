// Package abi verifies that the generated C-layout types in package vk match
// the real Vulkan C ABI.
//
// The bindings dispatch through libffi rather than CGo, so nothing in the
// normal build ever compares a cVkFoo struct against the C struct it stands in
// for. A wrong field offset or a missing pad byte produces no compile error and
// no validation message — just memory corruption at the FFI boundary. This
// package closes that gap by asking a C compiler: it emits a small program that
// reports sizeof, _Alignof and offsetof for every generated struct (plus the
// value of every generated enumerant) against the real Vulkan headers, then
// compares the answers with Go's own reflect-reported layout.
//
// The specs are produced by the generator, which is also what emits the
// cVkFoo structs, so the set under test always matches the current bindings.
package abi

import "reflect"

// StructSpec describes one generated C-layout type to check against C.
type StructSpec struct {
	// CName is the struct's name in the Vulkan headers, e.g. "VkDeviceCreateInfo".
	CName string

	// GoType is the generated C-layout type, e.g. reflect.TypeOf(cVkDeviceCreateInfo{}).
	GoType reflect.Type

	// CFields names the C struct member corresponding to each field of GoType,
	// in declaration order. An empty entry marks a synthetic field with no C
	// counterpart — explicit padding, or the uint32 word that C bitfield members
	// are folded into — which has no offsetof() to compare against.
	CFields []string

	// IsUnion reports whether the C type is a union. Unions are represented in
	// Go as [N]byte, so their size is checked but their alignment is not: Go
	// infers alignment 1 for a byte array, and the generator compensates with
	// explicit padding at each use site.
	IsUnion bool
}

// EnumSpec describes one generated constant to check against C.
type EnumSpec struct {
	// CName is the enumerant's name in the Vulkan headers, e.g. "VK_SUCCESS".
	CName string

	// GoValue is the generated Go constant, widened by [V].
	GoValue uint64
}

// Group is a set of specs sharing a single C translation unit.
//
// Macro, when non-empty, is a VK_USE_PLATFORM_* macro defined ahead of
// <vulkan/vulkan.h> so that the platform's own system headers are pulled in.
// A group with a Macro is skipped rather than failed when it does not compile,
// since that normally means the platform's development headers are not
// installed rather than that the bindings are wrong.
type Group struct {
	Name    string
	Macro   string
	Structs []StructSpec
	Enums   []EnumSpec
}

// V widens a generated Vulkan constant to uint64 for comparison.
//
// The bit pattern of negative values (VkResult error codes) and of bit 63
// (VkFlags64 bitmasks) is preserved. The conversion deliberately runs on a
// non-constant operand: converting a typed constant such as Result(-1) straight
// to uint64 is a compile-time overflow error, whereas the same conversion on a
// variable sign-extends as intended.
func V[T ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint](v T) uint64 {
	return uint64(v)
}
