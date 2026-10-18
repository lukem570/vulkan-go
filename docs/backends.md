# FFI backends

Every Vulkan call goes through one of two interchangeable foreign-call
backends, chosen when you compile:

| Build | Backend | Requires |
| --- | --- | --- |
| `CGO_ENABLED=0` (default) | [goffi](https://github.com/go-webgpu/goffi), pure Go | nothing |
| `CGO_ENABLED=1` | libffi through cgo | libffi and a C compiler |

```sh
CGO_ENABLED=0 go build ./...                 # goffi
CGO_ENABLED=1 go build ./...                 # cgo + libffi
CGO_ENABLED=1 go build -tags vkcgo ./...     # cgo, and fail the build if cgo is off
```

There is no backend switch to set. goffi cannot be compiled with cgo enabled —
it dispatches through `cgo_import_dynamic`, which requires `CGO_ENABLED=0` — and
a cgo file cannot be compiled with cgo disabled, so the `cgo` build constraint
partitions the two cleanly. The `vkcgo` tag only turns a silent fallback to
goffi into a build error, which is useful in CI where `CGO_ENABLED` may not be
what you assume.

`ffi.Backend` reports which one is linked in, as `"goffi"` or `"cgo"`.

## Which to use

Prefer the default. The goffi backend needs no C toolchain, cross-compiles
without one, and is what the library is tuned for.

Reach for `CGO_ENABLED=1` when something else in your program requires cgo —
another binding, a C library, a profiler — because then goffi cannot be built at
all and the whole binary would otherwise be unbuildable.

## What the seam looks like

The generated bindings never name an FFI library. They declare one call
interface per Vulkan command and go through five primitives in
`internal/ffi`:

```go
LoadLibrary(name string) (unsafe.Pointer, error)
Symbol(lib unsafe.Pointer, name string) (unsafe.Pointer, error)
Prepare(cif *CIF, ret Type, args []Type) error
Call(cif *CIF, fn, ret unsafe.Pointer, args []unsafe.Pointer) error
NewCallback(fn any) uintptr
```

`Type` is a plain code (`Pointer`, `UInt32`, `SInt32`, …) rather than a backend
type descriptor, which is what keeps the generated code backend-agnostic. Adding
a third backend means implementing those five functions and nothing else.

## Differences you may notice

- **libffi is a build dependency** of the cgo backend. On Debian/Ubuntu that is
  `libffi-dev`; on Arch, `libffi`; on macOS it ships with the SDK.
- **goffi remains a module dependency either way.** `go mod tidy` resolves
  imports across build configurations, so goffi stays in `go.sum` even for a
  cgo-only build. It is not linked into the binary.
- **Mixed float and integer callback arguments** are marshalled incorrectly by
  goffi v0.5.0 (a `float32` following integer arguments arrives as 0); libffi
  handles them. No Vulkan callback has that shape, so the bindings are
  unaffected — see the skipped case in `internal/ffi/ffi_test.go`.
- **Per-call cost differs.** The cgo backend copies each argument into C scratch
  memory before `ffi_call`, on top of cgo's own call overhead. The goffi backend
  calls the driver directly.

## Why arguments are copied into C memory

In the cgo backend, `Call` copies each argument's bytes into a pooled C block
and hands libffi a `void**` that points only into C memory. Passing a
Go-allocated array whose elements are Go pointers to C would violate the cgo
pointer rules and is rejected by the runtime's pointer checks.

Relatedly, any storage the driver writes through — a returned handle, an
enumerate count, a queried struct — is heap-allocated and pinned by the
generated wrapper rather than left on the goroutine stack. A Go callback firing
during a Vulkan call (a debug messenger under a validation layer, for instance)
runs Go code on the goroutine that is inside the call; if that grows the
goroutine's stack, Go relocates it, and any stack address the driver still holds
goes stale. The symptom was `EnumeratePhysicalDevices` returning zero devices
with no error, but only while a messenger callback was firing. Pinned heap
storage removes the hazard on both backends.
