# API Changes from Vulkan

How the generated Go surface differs from the C API in `vk.xml`. These
conventions are what the `v1` import path commits to — see
[Versioning](../README.md#versioning).

## Initialization

`Initialize()` loads the Vulkan shared library (`libvulkan.so.1`, `vulkan-1.dll`, or `libvulkan.1.dylib`) and bootstraps global function pointers. No explicit `LoadInstance` or `LoadDevice` step is needed — dispatch tables are built automatically:

```go
vk.Initialize()

instance, _ := vk.CreateInstance(...)  // instance-level functions ready immediately
device, _ := physDevice.CreateDevice(...)  // device-level functions ready immediately
```

Each `Instance` carries its own table of instance-level function pointers; each `Device` carries its own table of device-level function pointers. Handles derived from a device (Queue, CommandBuffer, etc.) share the same device table.

## Methods on handles

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

## Return values

Functions that write to an output pointer parameter return that value directly (along with `error` for `VkResult`-returning functions):

```c
// C
VkResult vkCreateFence(VkDevice, const VkFenceCreateInfo*, const VkAllocationCallbacks*, VkFence* pFence);
```

```go
// Go
fence, err := device.CreateFence(&vk.FenceCreateInfo{}, nil)
```

## Arrays

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

## Safe pNext chaining

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

## Function pointer callbacks

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

## Helpers

Two hand-written helpers are emitted into `package vk` alongside the bindings:

```go
vk.MakeAPIVersion(0, 1, 4, 0)      // VK_MAKE_API_VERSION
vk.GoString(props.DeviceName[:])   // null-padded char array → string
```

## Windowing

**GLFW is not compatible with this library.** [go-gl/glfw](https://github.com/go-gl/glfw) requires CGo, which conflicts with goffi's requirement for `CGO_ENABLED=0`. A pure-Go windowing library integration is planned.

In the meantime, `SurfaceKHRFromHandle` can wrap an opaque surface handle obtained from any source that doesn't require CGo:

```go
surface := vk.SurfaceKHRFromHandle(handleValue)
```

## Overhead

All Go structs are deep-copied to their C-layout counterparts on each call via goffi. This is the cost of the safe, idiomatic API surface. No CGo call overhead — goffi calls the Vulkan shared library directly via libffi.
