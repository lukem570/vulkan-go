# vulkan-go

Unofficial Go bindings for Vulkan, auto-generated from `vk.xml`.

## Installation

```sh
go get github.com/lukem570/vulkan-go/pkg/raw
```

## Quick Start

```go
import vk "github.com/lukem570/vulkan-go/pkg/raw"

// Initialize Vulkan (must be called first)
if err := vk.Initialize(); err != nil {
    log.Fatal(err)
}

instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
    ApplicationInfo: &vk.ApplicationInfo{
        ApplicationName: "my-app",
        ApiVersion:      vk.MakeApiVersion(0, 1, 4, 0),
    },
}, nil)
if err != nil {
    log.Fatal(err)
}
defer instance.Destroy(nil)

vk.LoadInstance(instance)
```

## Examples

See the [test/](./test) directory for full working examples:

- [triangle](./test/triangle/) — windowed graphics pipeline with GLFW, swapchain, render pass, and shader modules
- [compute](./test/compute/) — headless compute pipeline with buffer readback
- [surface](./test/surface/) — surface creation and physical device selection with GLFW

## API Changes from Vulkan

### Initialization

Vulkan-go uses [Volk](https://github.com/zeux/volk) to load the Vulkan dynamic library at runtime. Call `Initialize()` before any other Vulkan function, then call `LoadInstance` and `LoadDevice` after creating them to load extension function pointers:

```go
vk.Initialize()

instance, _ := vk.CreateInstance(...)
vk.LoadInstance(instance)

device, _ := physDevice.CreateDevice(...)
vk.LoadDevice(device)
```

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

Length+pointer pairs in structs are collapsed into a single Go slice:

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

Every struct that supports `pNext` has a `Next` field typed as the `Structure` interface, which ensures only compatible structs can be chained. The `sType` field is set automatically:

```go
type Structure interface {
    GetType() StructureType
}
```

```go
// Type-safe chaining — sType is set automatically
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

PFN fields on structs accept Go closures directly. The binding layer handles CGo trampoline setup automatically:

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

### GLFW interop

Surface creation with [go-gl/glfw](https://github.com/go-gl/glfw) requires passing the instance handle and converting the returned surface pointer:

```go
surfPtr, err := window.CreateWindowSurface((*byte)(instance.Handle()), nil)
surface := vk.SurfaceKHRFromGLFW(surfPtr)
defer instance.DestroySurfaceKHR(surface, nil)
```

## Code Generation

Bindings are generated from `mod/Vulkan-Headers/registry/vk.xml` using the generator in [cmd/gen/](./cmd/gen/). Configuration (API version, extensions, platforms) is in [config.yml](./config.yml).

```sh
go run ./cmd/gen
```

Generated output: `pkg/raw/vulkan.go`, `pkg/raw/vulkan_linux.go`, `pkg/raw/vulkan_windows.go`, `pkg/raw/volk_wrappers.{c,h}`

## Overhead

All Go structs are deep-copied to their C counterparts on each call. This is the cost of the safe, idiomatic API surface.
