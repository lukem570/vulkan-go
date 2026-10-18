# vulkan-go

Unofficial Go bindings for Vulkan, generated from `vk.xml`. Idiomatic rather
than literal: handle methods, slices, `error` returns, garbage-collected
handles. No CGo required.

```sh
go get github.com/lukem570/vulkan-go
```

You need a Vulkan runtime. The foreign-call backend is chosen by the `cgo` build
constraint, with nothing to configure:

| Build | Backend | Needs |
| --- | --- | --- |
| `CGO_ENABLED=0` (default) | [goffi](https://github.com/go-webgpu/goffi), pure Go | nothing |
| `CGO_ENABLED=1` | libffi through cgo | libffi and a C compiler |

Pure Go is the recommended default; the cgo backend exists so this library works
inside a program that needs cgo for something else — see
[docs/backends.md](./docs/backends.md).

## Quick start

```go
import (
    "log"

    vk "github.com/lukem570/vulkan-go/v1"
)

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

Two things are doing work there that the C API makes you write by hand:
instance- and device-level function pointers load themselves (no
`LoadInstance`/`LoadDevice`), and there is no teardown code — handles are
destroyed by the garbage collector, children before parents.

A full runnable program is in [examples/compute](./examples/compute/): a
headless compute pipeline with buffer readback. There are no windowed examples,
as GLFW needs CGo; see [Windowing](./docs/api-changes.md#windowing).

## Differences from the C API

- Commands taking a handle first become methods — `physDev.CreateDevice(ci, nil)`
- Output parameters become return values; `VkResult` becomes `error`
- `count` + `pFoo` pairs collapse into slices
- `pNext` becomes a typed `Next` field, and `sType` is set for you
- PFN fields take Go closures
- Dispatch tables load themselves
- **No `Destroy`/`Free` methods.** The GC destroys handles once unreachable,
  parents after children. You must still keep a handle reachable while
  submitted GPU work uses it.

Each of these, with C-vs-Go examples and the full lifetime rules:
**[docs/api-changes.md](./docs/api-changes.md)**.

## Versioning

Two independent axes:

- **The import path** carries the Go API generation. `v1` commits to the
  conventions above, whichever Vulkan headers the code was generated from. If
  those conventions ever have to change, a `v2` package appears beside it rather
  than breaking `v1` callers.
- **Release tags** carry the Vulkan header version. `v1.4.365` means *Vulkan
  headers 1.4.365*, not Go-library semver, so a tag bump is not a compatibility
  signal in either direction. Pin an exact tag if you need reproducibility.

## Working on the bindings

`v1/` is generated — never edit it by hand. Change the generator in
[internal/generator/](./internal/generator/) or the API surface in
[config.yml](./config.yml) (version, extensions, platforms), then:

```sh
task generate     # regenerate v1/ from mod/Vulkan-Headers/registry/vk.xml
task test         # driver-free test suite
task test:compute # run the compute example (needs glslc and a driver)
task test:both    # test suite against both FFI backends
```
