# Vulkan go

Vulkan-go is the unofficial bindings 
for `vulkan.h`.

All functions are generated using 
`vk.xml`, so the bindings should update
as new Vulkan updates are released

This adds overhead as all go structs are
deep copied to the c structs.

Examples are in the [test](./test) directory.

## Changes made

Below is a list of changes made to the
default Vulkan API. 

### Volk library loading

Vulkan-go uses Volk to load the Vulkan
dynamic library at runtime. To initialize
this loading use the following function.

```go
err := vulkan.Initialize()
```

This must be called before using any other
Vulkan functions

### Array detection

In the classic Vulkan API arrays are defined 
as such:

```c
struct VkDescriptorSetAllocateInfo {
	//...
    uint32_t descriptorSetCount;
	VkDescriptorSetLayout pSetLayouts;     
};
```

Vulkan-go converts this into a go array

```go
type DescriptorSetAllocateInfo struct {
	// ...
	SetLayouts []DescriptorSetLayout
}
```

### Object detection

In the classic Vulkan API there are 
many commands that begin with a handle.
For instance,

```c
void vkDestroyInstance(
    VkInstance instance, 
    const VkAllocationCallbacks* pAllocator);
```

Vulkan-go converts to:

```go
func (h *Instance) Destroy(pAllocator *AllocationCallbacks)
```

### Safe linked lists

In classic Vulkan the user is expected
to attach a `sType` and `pNext` to several
structs. If the `sType` doesn't match the
structure then undefined behavior occurs

To fix this every structure that has `sType` 
implements the interface:

```go
type Structure interface {
    GetType() StructureType
}
```

All references to `pNext` are converted
into `Next` and are typed as `Structure`.

### Return type detection

Most functions in Vulkan return a `VkResult`
or `void`. And the function modifies a pointer
passed in as a parameter.
