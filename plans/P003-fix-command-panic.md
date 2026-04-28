# P003 - Fix Command Panic

## Objective
Fix the runtime panic in `doptctl` caused by executing commands without the required subcommands or arguments (Issue #12). Specifically, when a command is missing arguments, the CLI should print the help documentation instead of crashing.

## Key Files & Context
- `commands/command.go`: Central command dispatcher that currently lacks slice bounds checking.

## Implementation Steps
1. Modify `getCommand(input []string) Command` in `commands/command.go`:
   - Add a check for `len(input) == 0` at the beginning of the function. If true, call the global `Help()` and return.
   - For each case (`benchmark`, `context`, `describe`, `list`, `simulation`), verify the length of `input` before accessing indices like `input[1]` or `input[2]`.
   - If the required arguments are missing, redirect to the `help` function with the specific `commandName` to print the command's help message.

### Example logic for length checks
- `benchmark`, `context`, `list`, `simulation`: Requires `len(input) >= 2`.
- `describe`: Requires `len(input) >= 3` because it uses `input[1]` (entity type) and `input[2]` (entity ID).

2. Modify `commands/command_test.go` (Optional/Bonus):
   - Add test cases to verify that `doptctl context`, `doptctl list`, etc. do not panic when run with missing arguments.

## Verification & Testing
- Run `doptctl context` and verify it prints the context help.
- Run `doptctl list` and verify it prints the list help.
- Run `doptctl describe` and verify it prints the describe help.
- Run `go test ./...` to ensure no panics occur.
