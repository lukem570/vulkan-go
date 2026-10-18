# vulkan-go

Unofficial Go bindings for Vulkan, auto-generated from `vk.xml`. Built on [goffi](https://github.com/go-webgpu/goffi) — no CGo required.

## Requirements

- `CGO_ENABLED=0` (goffi is incompatible with CGo)
- Vulkan runtime installed on your system

```sh
go env -w CGO_ENABLED=0
```

## Installation

```sh
CGO_ENABLED=0 go get github.com/lukem570/vulkan-go
```

## Quick Start

```go
import (
    vk "github.com/lukem570/vulkan-go/v1"
)

// Initialize Vulkan (loads the shared library and global function pointers)
if err := vk.Initialize(); err != nil {
    log.Fatal(err)
}

// Instance and device function pointers are loaded automatically on creation
instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
    ApplicationInfo: &vk.ApplicationInfo{
        ApplicationName: "my-app",
        ApiVersion:      vk.MakeAPIVersion(0, 1, 4, 0),
    },
}, nil)
if err != nil {
    log.Fatal(err)
}
defer instance.Destroy(nil)

// No LoadInstance/LoadDevice calls needed — dispatch tables are populated
// automatically when CreateInstance and CreateDevice return.
device, err := physicalDevice.CreateDevice(&vk.DeviceCreateInfo{...}, nil)
```

## Examples

See the [examples/](./examples) directory:

- [compute](./examples/compute/) — headless compute pipeline with buffer readback

There are no windowed examples; GLFW requires CGo and is incompatible with this
library. See [Windowing](./docs/api-changes.md#windowing).

## Differences from the C API

The bindings are idiomatic Go, not a literal translation:

- Commands taking a handle first become methods — `instance.Destroy(nil)`
- Output parameters become return values, `VkResult` becomes `error`
- `count` + `pFoo` pairs collapse into slices
- `pNext` becomes a typed `Next` field; `sType` is set for you
- PFN fields take Go closures
- Dispatch tables load themselves; there is no `LoadInstance`/`LoadDevice`

Full details, with C-vs-Go examples: **[docs/api-changes.md](./docs/api-changes.md)**

## Code Generation

Bindings are generated from `mod/Vulkan-Headers/registry/vk.xml` using the generator in [cmd/vkgen/](./cmd/vkgen/). Configuration (API version, extensions, platforms) is in [config.yml](./config.yml).

```sh
task generate
```

Or manually:

```sh
git submodule update --init mod/Vulkan-Headers
go run ./cmd/vkgen
```

Generated output: `v1/vulkan.go`, `v1/callbacks.go`, `v1/vulkan_linux.go`, `v1/vulkan_windows.go`, `v1/vulkan_darwin.go`

## Versioning

The import path carries the API generation: `v1` is a commitment to the binding
conventions in [docs/api-changes.md](./docs/api-changes.md), independent of
which Vulkan headers the code was generated from. If those conventions ever have
to change, a `v2` package will appear beside it rather than silently breaking
`v1` callers.

Release tags are derived from the Vulkan header version — `v1.4.365` means
*Vulkan headers 1.4.365*, not Go-library semver. Pin an exact tag if you need
reproducibility; do not read a tag bump as a compatibility signal in either
direction.
