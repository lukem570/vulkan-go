//go:build !cgo && vkcgo

// The vkcgo build tag asserts the cgo backend, but cgo is disabled, so the
// backend cannot be compiled.
//
// Either enable cgo:
//
//	CGO_ENABLED=1 go build -tags vkcgo ./...
//
// or drop the tag to use the pure-Go goffi backend:
//
//	CGO_ENABLED=0 go build ./...

package ffi

// Compile-time assertion: the identifier names the problem, and the package
// comment above explains it.
var _ = VKCGO_TAG_REQUIRES_CGO_ENABLED_1
