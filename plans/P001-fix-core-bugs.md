# P001-fix-core-bugs

## Context
Initial investigation of the `doptctl` codebase revealed critical bugs that prevent basic operation:
1. `ClientContext.URL()` incorrectly uses `Name` instead of `Host`. (Ref: [Issue #1](https://github.com/felipedreis/doptctl/issues/1))
2. `commands.Run` panics when global options are provided due to uninitialized map. (Ref: [Issue #2](https://github.com/felipedreis/doptctl/issues/2))
3. `commands.getCommand` returns `nil` for unknown commands, causing dereference panics. (Ref: [Issue #2](https://github.com/felipedreis/doptctl/issues/2))

## Tasks
- [x] Fix `commands.Run` uninitialized map panic.
- [x] Fix `commands.getCommand` nil dereference panic.
- [x] Create `commands/command_test.go` to verify panic fixes.
- [x] Fix `ClientContext.URL()` to use `Host`.
- [x] Fix path concatenation in `context.go` using `filepath.Join`.
- [x] Create `commands/context/context_test.go` to verify URL and path fixes.

## Design Decisions
- Maintain existing architecture but improve robustness.
- Use `filepath.Join` for cross-platform compatibility.
- Ensure `getCommand` fails gracefully by returning a help/error state instead of `nil`.

## Acceptance Criteria
- `go test ./...` passes.
- `doptctl --any=val context list` does not panic.
- `ClientContext.URL()` returns `host:port`.
- Unknown commands display help instead of panicking.
