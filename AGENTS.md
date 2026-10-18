# AGENTS.md

Notes for agents working in this repository.

## What this is

`github.com/lukem570/vulkan-go` — unofficial Go bindings for Vulkan, generated
from `vk.xml` and dispatched through [goffi](https://github.com/go-webgpu/goffi)
(libffi), with or without CGo. `README.md` is the overview and
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

Hand-written code lives only in `internal/generator/`, `internal/abi/`,
`cmd/vkgen/`, and `examples/`.

The generator also emits `v1/abi_data_test.go` (plus one per platform build
tag). Those are generated like everything else in `v1/` — do not hand-edit
them; they are the input data for the ABI conformance test described under
[Tests](#tests).

## Layout

| Path | Role |
| --- | --- |
| `v1/` | **Generated.** `package vk` — bindings, callbacks, per-platform files, ABI test data |
| `cmd/vkgen/main.go` | Generator entry point; wires the pipeline and writes files |
| `internal/generator/` | The generator: parse → registry → filter → codegen |
| `internal/abi/` | Hand-written. Compares the generated layout against the C headers |
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
5. `GenerateABIDataFiles` (`abidata.go`) — must run after `GeneratePackage`,
   which decides via `SetReserved` which enumerants the package declares
6. `WriteFile` → gofmt, then write

`WriteFile` runs `go/format` on the output. **If gofmt fails, the raw text is
dumped to `<path>.unformatted` and the error names it** — read that file to see
the broken emission rather than guessing. `*.unformatted` is gitignored.

Platform-specific symbols are split out by `PlatformBuildTag` in `spec.go`
(xlib/xcb/wayland → `linux`, win32 → `windows`, metal/macos → `darwin`), one
file per build tag.

## Commands

```sh
task setup          # init the Vulkan-Headers submodule (required once)
task generate       # rm -rf v1, then regenerate
task test           # generate, then go test ./internal/... ./v1/... (no GPU needed)
task test:compute   # generate, compile test.comp with glslc, run the compute example
task golden:update  # rewrite the generator golden files after an intended change
```

Environment:

- `CGO_ENABLED` must keep working both ways. goffi supports `0` (its own
  fakecgo) and `1` (the real cgo runtime) as of v0.6.0, and the bindings are
  expected to build and run under either, so run `task test:both` after
  touching the generator.
- `GOWORK=off` — set by every taskfile target; keep it when invoking `go`
  directly so a stray `go.work` does not shadow the module.

`task test:compute` needs `glslc` on PATH and a working Vulkan runtime/driver.
In a headless or driver-less environment, run `task test` instead — it needs
neither, only a C compiler. Note that `go build ./...` fails until
`examples/compute/test.spv` exists, since it is `go:embed`-ed and gitignored;
that is why `task test` scopes itself to `./internal/... ./v1/...`.

## Tests

Two suites, both driver-free; `task test` runs both.

**ABI conformance** (`internal/abi`, data generated into `v1/abi_data*_test.go`).
Nothing in the normal build ever checks a `cVkFoo` struct against the C struct
it stands in for — there is no CGo to do it. A wrong field offset or a missing
pad byte produces no compile error and no validation message, just corruption at
the FFI boundary. So the test emits a C program against
`mod/Vulkan-Headers`, reporting `sizeof`/`_Alignof`/`offsetof` for every
generated struct and the value of every generated enumerant, and compares the
answers with Go's `reflect` view. Needs `cc`, `gcc` or `clang`; skips if none is
present, as does any platform group whose system headers are missing.

Subtests are split per group (`core`, `xlib`, `xcb`, …) and per dimension
(`structs`, `enums`, `enums_declared_but_absent_from_headers`) so one defect
leaves the other dimensions usable as a gate.

**Generator goldens** (`internal/generator/golden_test.go`, fixtures in
`testdata/`). Each fixture is a hand-written registry fragment small enough to
read, pinning the conventions listed above — handle receivers, output params,
slice collapsing, union padding, bitfield packing, platform splitting. After an
intended change run `task golden:update` and read the diff; an unexplained line
in it is an API break for consumers of `v1`.

Two invariants keep the suites honest, and both are load-bearing:
`Structured.cLayoutFields` is the single source of truth for C-layout field
order (shared by the struct emitter and the ABI data), and
`Enum.generatedElements` is the single source of truth for which enumerants get
constants (shared by the const-block emitter and the ABI data). Bypass either
and the test starts checking something other than what ships.

## Conventions the generator enforces

These are load-bearing; a change in `internal/generator/` that breaks one of
them is an API break for every consumer of `v1`. They are documented for users
in `docs/api-changes.md` — keep that file in sync with any change here.

- Callbacks go through the generated `_newCallback`, never `ffi.NewCallback`
  directly. On Windows goffi delegates to `syscall.NewCallback`, which takes
  only pointer-sized arguments and demands exactly one result, and the Vulkan
  PFNs have neither; `_newCallback` adapts them. Since the trampolines are
  package-level variables, skipping it panics on import, not at call time.
- Storage the driver writes through — out params, enumerate counts, byte-data
  sizes — is `new`-allocated and pinned, never a goroutine stack local. A Go
  callback firing mid-call (a debug messenger under a validation layer) runs Go
  code on the goroutine that is inside the call; if its stack grows, Go
  relocates it and any stack address the driver still holds goes stale. The
  symptom was `EnumeratePhysicalDevices` returning zero devices, no error,
  under cgo only.
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

`.github/workflows/ci.yml` runs on pull requests and on pushes to `main` and
`road-to-v1`:

- **test**, on ubuntu, macos and windows. Regenerates, vets, and runs
  `go test ./internal/... ./v1/...`. The matrix is not redundant — each runner
  is the only one that can verify its own platform structs against the real
  system headers, so the win32 and metal/macos ABI groups are checked nowhere
  else. Ubuntu additionally installs `libx11-dev libxcb1-dev libwayland-dev`,
  without which those three groups skip instead of verifying.
- **test** also fails if `git status --porcelain v1/` is non-empty after
  regenerating, which catches a generator change committed without its
  regenerated output. Ubuntu only, since the runners disagree on line endings.
- **cross-compile** builds `./v1` for linux, windows and darwin on both amd64
  and arm64. The per-platform files are build-tagged, so a regression in one is
  invisible to a build on any other.

`.github/workflows/update.yml` runs every 6 hours: bumps the submodule, and if
`vk.xml` moved, regenerates, commits as `github-actions`, tags, and pushes to
whatever branch it is on.

That update is gated: after regenerating it runs `task test`, then vets and
cross-compiles, and only then commits and tags. A header change that shifts a
struct offset, changes an enumerant value, breaks a generator convention, or
emits uncompilable code now stops the workflow instead of being tagged as a
release. **A failing suite therefore halts automated header updates** — that is
the intended trade, but it means a red `task test` needs fixing promptly rather
than at leisure.

Three defects found by `TestABIConformance` had to be fixed before that gate
could hold, and each is now pinned by a golden fixture as well:

- Enumerants aliased in a top-level `<enums>` block generated as
  `0 // TODO: unknown value`, because `XMLEnum` had no `alias,attr` field while
  `XMLRequireEnum` did. Six constants were wrong, among them
  `StencilFrontAndBack` and `PipelineStage2TransferBit`.
- `VkBaseInStructure`/`VkBaseOutStructure` generated empty, 0 bytes against
  C's 16: `parseStruct` dropped `sType`/`pNext` unconditionally, but
  `detectSType` only recognises them when a `values=` attribute is present,
  which these two lack.
- 117 enumerants were declared that no Vulkan header defines.
  `applyEnumExtensions` walked every feature and extension in vk.xml, ignoring
  the feature's `api=`, the extension's `supported=`, and the `<require>`
  block's `api=`. Those three now intersect into
  `EnumElement.APIConstraints` and are pruned against `config.yml`'s `api:` in
  `Filter`. **`api:` is load-bearing now rather than decorative** — vk.xml is a
  combined registry carrying the Vulkan SC feature and every
  reserved-but-disabled extension alongside the real ones.

Both workflows take the toolchain from `go-version-file: go.mod` rather than a
pinned version.

`.github/dependabot.yml` covers Go modules and action versions weekly. The
submodule is deliberately left to the workflow above. A goffi bump may need
`task generate` re-run on the Dependabot branch before it is mergeable.

`.gitattributes` marks `v1/*.go` as `linguist-generated` so GitHub collapses
those diffs and omits them from language stats.

## Known gaps

- `AllocationCallbacks` does not work. The generated trampolines key the
  callback registry on the *last* callback parameter, which is `pUserData` for
  the debug messenger but the allocation scope for the allocator PFNs, so the
  lookup misses and the trampoline returns NULL. Passing a custom allocator
  fails with `ErrorOutOfHostMemory` on both backends. Fix: key the registry on
  whichever parameter actually carries `pUserData`.

- No windowed example. GLFW needs CGo and so cannot be used; a pure-Go
  windowing path is the intended replacement. `SurfaceKHRFromHandle` exists for
  wrapping a surface obtained elsewhere.
- Nothing runs the compute example in CI. A software Vulkan device (lavapipe)
  on the ubuntu runner would make `task test:compute` reachable there.
- The ABI probe only reports the host architecture. A 32-bit target, where
  non-dispatchable handles are `uint64` rather than pointer-sized, is unchecked.
- The `api` filter is applied to enumerants only. Structs, commands and types
  carry `api=`/`supported=` attributes too, and nothing prunes against them; it
  has not bitten because the generator reaches those through the enabled
  feature and extension lists rather than by walking the whole registry, which
  is precisely what `applyEnumExtensions` used to do.
- Count fields collapsed against a fixed-size array are lost:
  `PhysicalDeviceMemoryProperties` exposes `MemoryTypes [32]MemoryType` but not
  `memoryTypeCount`, so callers cannot tell which entries are valid.
- GC-driven destruction cannot observe GPU timelines. A handle dropped while
  submitted work still uses it is destroyed early — documented in
  `docs/api-changes.md`, not solvable inside the bindings.
