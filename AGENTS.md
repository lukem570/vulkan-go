# AGENTS.md

Notes for agents working in this repository.

## What this is

`github.com/lukem570/vulkan-go` — unofficial Go bindings for Vulkan, generated
from `vk.xml` and dispatched through [goffi](https://github.com/go-webgpu/goffi)
(libffi), so **no CGo**. `README.md` is the overview and
`docs/api-changes.md` is the full public API tour; this file covers how the repo
is built and what not to break.

## The one rule

**`v1/*.go` is generated output. Never hand-edit it.**

`v1/vulkan.go` alone is ~33k lines. To change anything in the bindings, change
the generator in `internal/generator/` (or `config.yml`) and regenerate. A patch
to `v1/` is silently reverted by the next `task generate` run — including the
scheduled CI run, which commits and tags on its own.

`task generate` does `rm -rf ./v1`, so hand-written files cannot live there. If
you need a new hand-written helper in `package vk`, emit it from the generator
(see `runtimeHelpers` in `internal/generator/spec.go`, which is where
`MakeAPIVersion` and `GoString` come from); otherwise put it in a new
non-generated package.

Hand-written code lives only in `internal/generator/`, `cmd/vkgen/`, and
`examples/`.

## Layout

| Path | Role |
| --- | --- |
| `v1/` | **Generated.** `package vk` — bindings, callbacks, per-platform files |
| `cmd/vkgen/main.go` | Generator entry point; wires the pipeline and writes files |
| `internal/generator/` | The generator: parse → registry → filter → codegen |
| `config.yml` | API version, extension allow-list, platform list |
| `mod/Vulkan-Headers` | Submodule; source of `registry/vk.xml` |
| `examples/compute/` | Runnable headless compute example (doubles as the test) |
| `docs/` | Prose docs: `api-changes.md` is the C-vs-Go API reference |

Generator pipeline, in order (`cmd/vkgen/main.go`):

1. `ParseXML("mod/Vulkan-Headers/registry/vk.xml")` → raw XML tree (`xml.go`, `parse.go`)
2. `BuildRegistry` → full `Registry` of enums, structs, handles, commands (`spec.go`)
3. `Filter(registry, config)` → `ReducedRegistry`, trimmed to the configured
   version/extensions/platforms (`filter.go`)
4. `GeneratePackage` / `GenerateCallbacksFile` / `GeneratePlatformFiles`
   (`output.go`, `structs.go`, `commands.go`, `enums.go`, `types.go`,
   `handles.go`, `lifetime.go`)
5. `WriteFile` → gofmt, then write

`WriteFile` runs `go/format` on the output. **If gofmt fails, the raw text is
dumped to `<path>.unformatted` and the error names it** — read that file to see
the broken emission rather than guessing. `*.unformatted` is gitignored.

Platform-specific symbols are split out by `PlatformBuildTag` in `spec.go`
(xlib/xcb/wayland → `linux`, win32 → `windows`, metal/macos → `darwin`), one
file per build tag.

## Commands

```sh
task setup      # init the Vulkan-Headers submodule (required once)
task generate   # rm -rf v1, then regenerate
task test       # generate, compile test.comp with glslc, run the compute example
```

Environment, non-negotiable:

- `CGO_ENABLED=0` — goffi is incompatible with CGo. Anything pulling in a CGo
  dependency (e.g. go-gl/glfw) cannot be used here.
- `GOWORK=off` — set by every taskfile target; keep it when invoking `go`
  directly so a stray `go.work` does not shadow the module.

`task test` needs `glslc` on PATH and a working Vulkan runtime/driver. In a
headless or driver-less environment, `go build ./...` and `go vet ./...` are the
useful checks — expect a runtime failure from the example, not a code error.
Note that `go build ./...` fails until `examples/compute/test.spv` exists, since
it is `go:embed`-ed and gitignored; run glslc first.

## Conventions the generator enforces

These are load-bearing; a change in `internal/generator/` that breaks one of
them is an API break for every consumer of `v1`. They are documented for users
in `docs/api-changes.md` — keep that file in sync with any change here.

- Commands whose first parameter is a handle become methods on that handle
  (`vkDestroyInstance` → `instance.Destroy`).
- Output-pointer parameters become return values; `VkResult` becomes `error`.
- `count` + `pFoo` pairs collapse into a single Go slice field.
- `pNext` becomes `Next Structure`; `sType` is set automatically, never by hand.
- PFN struct fields take Go closures; trampolines go through `ffi.NewCallback`.
- Dispatch tables are populated inside `CreateInstance` / `CreateDevice`, so
  there is no `LoadInstance` / `LoadDevice` step. Handles derived from a device
  share its table.
- `vkDestroy*`/`vkFree*` wrappers are generated **unexported**. Creation sites
  register a `runtime.AddCleanup` instead, so the GC destroys handles
  (`internal/generator/lifetime.go`). The invariants there are load-bearing:
  a cleanup argument must hold the parent (so parents outlive children) and
  must never hold the handle being destroyed (or it is pinned forever and the
  cleanup never runs). Cleanups take `_cleanupMu` because Go runs them
  concurrently while Vulkan requires external synchronization.

Adding an extension or platform is a `config.yml` edit, not a code change.

## Versioning

Two independent axes; do not conflate them.

- **Import path** (`vulkan-go/v1`) is the Go API generation. It promises the
  conventions above. Breaking one means adding a `v2` package beside `v1`, not
  mutating `v1`. Both forms resolve fine as plain subdirectories — the `/vN`
  module-path prohibition applies to `module` lines in go.mod, not to package
  directories. Caveat: if a real `vulkan-go/v2` *module* is ever published, it
  takes precedence over a `v2/` directory, so pick one mechanism and stay there.
- **Release tags** are the Vulkan header version. `v1.4.365` means *headers
  1.4.365*, minted by CI, and carries no Go-API compatibility meaning. Library
  semver cannot be expressed in these tags: a `v1.0.0` would sort below the
  existing `v1.4.x` tags and never resolve as `latest`.

## CI

`.github/workflows/update.yml` runs every 6 hours: bumps the submodule, and if
`vk.xml` moved, regenerates, commits as `github-actions`, tags, and pushes to
whatever branch it is on.

`.github/dependabot.yml` covers Go modules and action versions weekly. The
submodule is deliberately left to the workflow above. A goffi bump may need
`task generate` re-run on the Dependabot branch before it is mergeable.

`.gitattributes` marks `v1/*.go` as `linguist-generated` so GitHub collapses
those diffs and omits them from language stats.

## Known gaps

- No windowed example. GLFW needs CGo and so cannot be used; a pure-Go
  windowing path is the intended replacement. `SurfaceKHRFromHandle` exists for
  wrapping a surface obtained elsewhere.
- No CI runs on pull requests — the only workflow is the scheduled
  regeneration, so a Dependabot PR arrives with nothing actually built.
- The update workflow installs Go 1.22 while `go.mod` declares `go 1.25`; it
  currently survives only on toolchain auto-download.
- No `go test` suite exists — `examples/compute` is a `main` package run by hand.
- Count fields collapsed against a fixed-size array are lost:
  `PhysicalDeviceMemoryProperties` exposes `MemoryTypes [32]MemoryType` but not
  `memoryTypeCount`, so callers cannot tell which entries are valid.
- GC-driven destruction cannot observe GPU timelines. A handle dropped while
  submitted work still uses it is destroyed early — documented in
  `docs/api-changes.md`, not solvable inside the bindings.
