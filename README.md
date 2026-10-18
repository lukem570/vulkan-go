# vulkan-go

Unofficial Go bindings for Vulkan, automatically generated from `vk.xml`.
For changes made to the api check (api-changes.md)[docs/api-changes.md].
Vulkan-go also has a no cgo mode.

## Installing

```sh
go get github.com/lukem570/vulkan-go
```

## Quick start

```go
// Loads the Vulkan shared library and the global function pointers.
if err := vk.Initialize(); err != nil {
	log.Fatal(err)
}

instance, err := vk.CreateInstance(&vk.InstanceCreateInfo{
	ApplicationInfo: &vk.ApplicationInfo{
    ApplicationName: "my-app",
    ApiVersion:      vk.MakeAPIVersion(0, 1, 4, 0),
	},
}, nil)
if err != nil {
	log.Fatal(err)
}

physicalDevices, err := instance.EnumeratePhysicalDevices()
if err != nil {
	log.Fatal(err)
}
physDev := physicalDevices[0]
props := physDev.GetProperties()
log.Println(vk.GoString(props.DeviceName[:]))

device, err := physDev.CreateDevice(&vk.DeviceCreateInfo{
    QueueCreateInfos: []vk.DeviceQueueCreateInfo{{
        QueueFamilyIndex: 0,
        QueuePriorities:  []float32{1.0},
    }},
}, nil)
if err != nil {
	log.Fatal(err)
}
queue := device.GetQueue(0, 0)
```

## Common taskfile commands

```sh
task generate     # regenerate v1/ from mod/Vulkan-Headers/registry/vk.xml
task test         # driver-free test suite
task test:compute # run the compute example (needs glslc and a driver)
task test:both    # test suite against both FFI backends
```
