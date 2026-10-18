# AGENTS.md

Notes for agents working in this repository.

## What this is

`github.com/lukem570/vulkan-go` — unofficial Go bindings for Vulkan, generated
from `vk.xml` and dispatched through [goffi](https://github.com/go-webgpu/goffi)
(libffi), so **no CGo**. See `README.md` for the public API tour; this file
covers how the repo is built and what not to break.

## The one rule

**`pkg/raw/*.go` is generated output. Never hand-edit it.**

`pkg/raw/vulkan.go` alone is ~33k lines. To change anything in the bindings,
change the generator in `internal/generator/` (or `config.yml`) and regenerate.
A patch to `pkg/raw` is silently reverted by the next `task generate` run —
including the scheduled CI run, which commits and tags on its own.

Hand-written code lives only in `internal/generator/`, `cmd/gen/`, `pkg/util/`,
and `test/`.

## Layout

| Path | Role |
| --- | --- |
| `cmd/gen/main.go` | Generator entry point; wires the pipeline and writes files |
| `internal/generator/` | The generator: parse → registry → filter → codegen |
| `config.yml` | API version, extension allow-list, platform list |
| `mod/Vulkan-Headers` | Submodule; source of `registry/vk.xml` |
| `pkg/raw/` | **Generated.** Bindings + per-platform files + callbacks |
| `pkg/util/` | Hand-written helpers (`VkMakeApiVersion`, `ToString`) |
| `test/compute/` | Runnable headless compute example (doubles as the test) |

Generator pipeline, in order (`cmd/gen/main.go`):

1. `ParseXML("mod/Vulkan-Headers/registry/vk.xml")` → raw XML tree (`xml.go`, `parse.go`)
2. `BuildRegistry` → full `Registry` of enums, structs, handles, commands (`spec.go`)
3. `Filter(registry, config)` → `ReducedRegistry`, trimmed to the configured
   version/extensions/platforms (`filter.go`)
4. `GeneratePackage` / `GenerateCallbacksFile` / `GeneratePlatformFiles`
   (`output.go`, `structs.go`, `commands.go`, `enums.go`, `types.go`, `handles.go`)
5. `WriteFile` → gofmt, then write

`WriteFile` runs `go/format` on the output. **If gofmt fails, the raw text is
dumped to `<path>.unformatted` and the error names it** — read that file to see
the broken emission rather than guessing. `*.unformatted` is gitignored.

Platform-specific symbols are split out by `PlatformBuildTag` in `spec.go`
(xlib/xcb/wayland → `linux`, win32 → `windows`, metal/macos → `darwin`), one
file per build tag.

## Commands

```sh
task setup          # init the Vulkan-Headers submodule (required once)
task generate       # rm -rf pkg/raw, then regenerate
task test           # generate, compile test.comp with glslc, run compute example
task test-surface   # see "known gaps" — the target it runs is absent
```

Environment, non-negotiable:

- `CGO_ENABLED=0` — goffi is incompatible with CGo. Anything pulling in a CGo
  dependency (e.g. go-gl/glfw) cannot be used here.
- `GOWORK=off` — set by every taskfile target; keep it when invoking `go`
  directly so a stray `go.work` does not shadow the module.

`task test` needs `glslc` on PATH and a working Vulkan runtime/driver. In a
headless or driver-less environment, `go build ./...` and `go vet ./...` are the
useful checks — expect a runtime failure from the example, not a code error.

## Conventions the generator enforces

These are load-bearing; a change in `internal/generator/` that breaks one of
them is an API break for every consumer.

- Commands whose first parameter is a handle become methods on that handle
  (`vkDestroyInstance` → `instance.Destroy`).
- Output-pointer parameters become return values; `VkResult` becomes `error`.
- `count` + `pFoo` pairs collapse into a single Go slice field.
- `pNext` becomes `Next Structure`; `sType` is set automatically, never by hand.
- PFN struct fields take Go closures; trampolines go through `ffi.NewCallback`.
- Dispatch tables are populated inside `CreateInstance` / `CreateDevice`, so
  there is no `LoadInstance` / `LoadDevice` step. Handles derived from a device
  share its table.

Adding an extension or platform is a `config.yml` edit, not a code change.

## CI and versioning

`.github/workflows/update.yml` runs every 6 hours: bumps the submodule, and if
`vk.xml` moved, regenerates, commits as `github-actions`, tags, and pushes to
whatever branch it is on.

Tags are derived from `VK_HEADER_VERSION_COMPLETE` in `vulkan_core.h`, i.e.
`v1.4.365` means *Vulkan header 1.4.365*, not library semver. Any v1 release
plan has to resolve that collision: the module's `v1.x` tag space is already
occupied by header versions, and Go's module rules read those tags as the
library's own API version.

## Known gaps

- `task test-surface` runs `./test/surface/test.go`, which is not in the repo —
  the GLFW-based examples were removed as CGo-incompatible. The README also
  links `test/triangle/` and `test/surface/`, both absent. A pure-Go windowing
  path is the intended replacement.
- The workflow installs Go 1.22 while `go.mod` declares `go 1.25`; it currently
  survives only on toolchain auto-download.
- No `go test` suite exists — `test/compute` is a `main` package run by hand.
