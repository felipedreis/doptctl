# P004 - Refactor Command Pattern

## Objective
Address the structural anti-pattern identified in PR #13 (repetitive `if` statements in `getCommand` for argument validation). The goal is to move validation logic into the individual commands so that they are aware of their expected parameters and display their own help when arguments are missing. This also involves modifying the command interface so that subcommand execution logic is defined at the constructor level.

## Key Files & Context
- `commands/command.go`: Central dispatcher and `Command` interface.
- `commands/benchmark/cmd.go`, `commands/context/cmd.go`, `commands/describe/cmd.go`, `commands/list/cmd.go`, `commands/simulation/cmd.go`: Subcommand packages.

## Implementation Steps
1. **Update `Command` Interface**
   Modify `Command` in `commands/command.go` to include a `Help()` method:
   ```go
   type Command interface {
       Execute(conn *grpc.ClientConn, opts map[string]string)
       Help()
   }
   ```

2. **Simplify Central Dispatcher**
   Refactor `getCommand` in `commands/command.go` to pass the remaining `input` arguments directly to the subcommand constructors, removing all bounds checks from the central dispatcher:
   ```go
   func getCommand(input []string) Command {
       if len(input) == 0 {
           Help()
           return nil
       }
       cmdName := input[0]
       args := input[1:]
       switch cmdName {
       case "benchmark": return benchmark.NewCommand(args)
       case "context": return context.NewCommand(args)
       case "describe": return describe.NewCommand(args)
       case "list": return list.NewCommand(args)
       case "simulation": return simulation.NewCommand(args)
       case "help": return help(args)
       default: Help(); return nil
       }
   }
   ```

3. **Refactor Subcommand Constructors and Types**
   For each sub-package (`benchmark`, `context`, `describe`, `list`, `simulation`), implement the Command pattern properly:
   - Change `NewXxxCommand` to `NewCommand(args []string) commands.Command`.
   - Update the existing command structs to implement `Help()`.
   - If the constructor receives an empty `args` slice (or missing required positional arguments), it should return an instance of a dedicated "Help Command" or configure the main command to execute the `Help()` function instead of proceeding with gRPC calls.

4. **Adjust `commands/command_test.go`**
   Update tests to reflect the new constructor signatures and ensure the "no panic" behavior is maintained.

## Verification & Testing
- Run `doptctl context`, `doptctl list`, `doptctl describe`, etc. without arguments and verify they display the specific help menu.
- Ensure the test suite passes (`go test ./...`).
