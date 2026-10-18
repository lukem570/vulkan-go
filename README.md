# vulkan-go

Unofficial Go bindings for Vulkan, auto-generated from `vk.xml`. Built on [goffi](https://github.com/go-webgpu/goffi) — no CGo required.

## Requirements

- `CGO_ENABLED=0` (goffi is incompatible with CGo)
- Vulkan runtime installed on your system

```sh
export CGO_ENABLED=0
```

Or permanently:

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

See the [examples/](./examples) directory for full working examples:

- [compute](./examples/compute/) — headless compute pipeline with buffer readback

> **Note:** There are no windowed examples. [go-gl/glfw](https://github.com/go-gl/glfw) requires CGo and is therefore incompatible with this library; a pure-Go windowing library is planned as a replacement.

## API Changes from Vulkan

### Initialization

`Initialize()` loads the Vulkan shared library (`libvulkan.so.1`, `vulkan-1.dll`, or `libvulkan.1.dylib`) and bootstraps global function pointers. No explicit `LoadInstance` or `LoadDevice` step is needed — dispatch tables are built automatically:

```go
vk.Initialize()

instance, _ := vk.CreateInstance(...)  // instance-level functions ready immediately
device, _ := physDevice.CreateDevice(...)  // device-level functions ready immediately
```

Each `Instance` carries its own table of instance-level function pointers; each `Device` carries its own table of device-level function pointers. Handles derived from a device (Queue, CommandBuffer, etc.) share the same device table.

### Methods on handles

Commands that take a handle as their first parameter become methods on that handle:

```c
// C
vkDestroyInstance(instance, pAllocator);
vkCreateDevice(physicalDevice, pCreateInfo, pAllocator, pDevice);
```

```go
// Go
instance.Destroy(pAllocator)
device, err := physDevice.CreateDevice(createInfo, nil)
```

### Return values

Functions that write to an output pointer parameter return that value directly (along with `error` for `VkResult`-returning functions):

```c
// C
VkResult vkCreateFence(VkDevice, const VkFenceCreateInfo*, const VkAllocationCallbacks*, VkFence* pFence);
```

```go
// Go
fence, err := device.CreateFence(&vk.FenceCreateInfo{}, nil)
```

### Arrays

Length+pointer pairs in structs and function signatures are collapsed into Go slices:

```c
// C
struct VkSubmitInfo {
    uint32_t        waitSemaphoreCount;
    VkSemaphore*    pWaitSemaphores;
    // ...
};
```

```go
// Go
vk.SubmitInfo{
    WaitSemaphores: []*vk.Semaphore{sem},
}
```

### Safe pNext chaining

Every struct that supports `pNext` has a `Next` field typed as the `Structure` interface. The `sType` field is set automatically:

```go
instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
    Next: &vk.ValidationFeaturesEXT{
        Next: &vk.DebugUtilsMessengerCreateInfoEXT{ /* ... */ },
        EnabledValidationFeatures: []vk.ValidationFeatureEnableEXT{
            vk.ValidationFeatureEnableBestPracticesEXT,
        },
    },
    // ...
}, nil)
```

### Function pointer callbacks

PFN fields on structs accept Go closures directly. The binding layer handles trampoline setup automatically via `ffi.NewCallback`:

```go
messenger, err := instance.CreateDebugUtilsMessengerEXT(&vk.DebugUtilsMessengerCreateInfoEXT{
    MessageSeverity: vk.DebugUtilsMessageSeverityFlagsEXT(
        vk.DebugUtilsMessageSeverityErrorBitEXT |
            vk.DebugUtilsMessageSeverityWarningBitEXT,
    ),
    MessageType: vk.DebugUtilsMessageTypeFlagsEXT(
        vk.DebugUtilsMessageTypeValidationBitEXT |
            vk.DebugUtilsMessageTypePerformanceBitEXT,
    ),
    PfnUserCallback: func(
        severity vk.DebugUtilsMessageSeverityFlagBitsEXT,
        types vk.DebugUtilsMessageTypeFlagsEXT,
        data *vk.DebugUtilsMessengerCallbackDataEXT,
        userData unsafe.Pointer,
    ) bool {
        fmt.Println("vulkan:", data.Message)
        return false
    },
}, nil)
```

## Windowing

**GLFW is not compatible with this library.** [go-gl/glfw](https://github.com/go-gl/glfw) requires CGo, which conflicts with goffi's requirement for `CGO_ENABLED=0`. A pure-Go windowing library integration is planned.

In the meantime, `SurfaceKHRFromHandle` can wrap an opaque surface handle obtained from any source that doesn't require CGo:

```go
surface := vk.SurfaceKHRFromHandle(handleValue)
```

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
conventions described above, independent of which Vulkan headers the code was
generated from. If those conventions ever have to change, a `v2` package will
appear beside it rather than silently breaking `v1` callers.

Release tags are derived from the Vulkan header version — `v1.4.365` means
*Vulkan headers 1.4.365*, not Go-library semver. Pin an exact tag if you need
reproducibility; do not read a tag bump as a compatibility signal in either
direction.

## Overhead

All Go structs are deep-copied to their C-layout counterparts on each call via goffi. This is the cost of the safe, idiomatic API surface. No CGo call overhead — goffi calls the Vulkan shared library directly via libffi.
