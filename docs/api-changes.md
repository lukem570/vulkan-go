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
vkCreateDevice(physicalDevice, pCreateInfo, pAllocator, pDevice);
vkGetDeviceQueue(device, queueFamilyIndex, queueIndex, pQueue);
```

```go
// Go
device, err := physDevice.CreateDevice(createInfo, nil)
queue := device.GetDeviceQueue(queueFamilyIndex, queueIndex)
```

The destroy half of the API is gone entirely — see
[Object lifetimes](#object-lifetimes).

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

## Object lifetimes

**There are no `Destroy` or `Free` methods.** Handles are destroyed by the
garbage collector: every create and allocate call registers a
`runtime.AddCleanup`, and the matching `vkDestroy*`/`vkFree*` runs once the
handle becomes unreachable.

```go
buf, err := device.CreateBuffer(&vk.BufferCreateInfo{...}, nil)
// no defer; no buf.Destroy(nil). Dropping the last reference to buf
// eventually calls vkDestroyBuffer.
```

### What the bindings guarantee

- **Parents outlive children.** A child's cleanup argument holds its parent, and
  every handle stores its parent in an unexported field, so the chain
  `Buffer → Device → PhysicalDevice → Instance` tears down bottom-up. A
  `vkDestroyDevice` cannot run while any of its objects are still alive.
- **Create-time dependencies are retained.** The create-info value is held for
  the lifetime of the object it created, which keeps every handle it references
  alive too — a `RenderPass` outlives a `Framebuffer` built from it, an `Image`
  outlives its `ImageView`. `BindBufferMemory` and `BindImageMemory` likewise
  retain the memory on the bound object.
- **Destruction is serialized.** Cleanups run concurrently in Go, so they take a
  package-wide lock before calling into Vulkan; `vkDestroyDevice` requires the
  device to be externally synchronized, and two cleanups racing on one device is
  a driver data race.
- **Pool-scoped frees work.** `CommandBuffer` and `DescriptorSet` cleanups keep
  their pool alive and free through it.

### What you are still responsible for

**Reachability in Go says nothing about the GPU timeline.** The collector cannot
see that submitted work still reads a resource, so a handle whose last Go
reference disappears mid-flight can be destroyed while the device is using it —
undefined behaviour, not a leak:

```go
cmdBuf.BindVertexBuffers(0, []*vk.Buffer{buf}, offsets)
queue.Submit(submits, fence)
// buf is unreachable here; the GPU is not finished with it.
```

Keep the handle reachable until the work that uses it has completed — hold it in
the struct that owns the frame, or keep it live until you have waited on the
fence. Note that recording a handle into a command buffer does **not** retain
it; only create-info dependencies and bound memory are retained.

Two more consequences of GC-driven destruction:

- Cleanups are not guaranteed to run before the process exits. That is normally
  harmless, since the driver reclaims everything at exit, but it means teardown
  is not a reliable place for external side effects.
- Full teardown takes several GC cycles, because each level of the ownership
  chain only becomes unreachable after the level below it has been cleaned up.

Handles obtained rather than created — `GetDeviceQueue`, the images from
`GetSwapchainImagesKHR`, a surface from `SurfaceKHRFromHandle` — are not owned
by the bindings and get no cleanup, which is correct: swapchain images must
never be passed to `vkDestroyImage`.

## Helpers

Two hand-written helpers are emitted into `package vk` alongside the bindings:

```go
vk.MakeAPIVersion(0, 1, 4, 0)      // VK_MAKE_API_VERSION
vk.GoString(props.DeviceName[:])   // null-padded char array → string
```

## Windowing

**There are no windowed examples.** [go-gl/glfw](https://github.com/go-gl/glfw) requires CGo, which goffi now tolerates (`CGO_ENABLED=1` works as of goffi v0.6.0), so GLFW is no longer ruled out — but the windowed examples have not been revived, and a pure-Go windowing library is still the plan.

In the meantime, `SurfaceKHRFromHandle` can wrap an opaque surface handle obtained from any source that doesn't require CGo:

```go
surface := vk.SurfaceKHRFromHandle(handleValue)
```

## Overhead

All Go structs are deep-copied to their C-layout counterparts on each call via goffi. This is the cost of the safe, idiomatic API surface. No CGo call overhead — goffi calls the Vulkan shared library directly via libffi.
