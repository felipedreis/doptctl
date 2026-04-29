# P005-describe-and-formatting.md

## Context
This plan addresses two issues:
1.  **Issue #6:** Implement `describe` command for Agents and Regions.
2.  **Issue #5:** Enhance Command Output Formatting (using tables for `list` and readable format for `describe`).

The current implementation of `describe` contains empty functions, and `list` uses `fmt.Println` which is not user-friendly for a CLI.

## Tasks
- [ ] Research and add a table formatting library (e.g., `github.com/olekukonko/tablewriter`).
- [ ] Create a utility for shared output formatting.
- [ ] Implement `describeAgent` in `commands/describe/cmd.go`.
- [ ] Implement `describeRegion` in `commands/describe/cmd.go`.
- [ ] Update `listAgents` in `commands/list/cmd.go` to use table formatting.
- [ ] Update `listRegions` in `commands/list/cmd.go` to use table formatting.
- [ ] Add unit/functional tests for the new functionality.

## Technical Approach

### 1. Output Formatting
I will use `github.com/olekukonko/tablewriter` for table output.
For `describe` commands, I will use a key-value pair formatting to display detailed information.

### 2. Describe Command Implementation
- **Agent:** Fetch details using `client.DescribeAgent`. Display fields: Agent ID, Lifetime, Start Time, Current Time, Complete Executions, Required Solutions, Heuristic, Memory Tax, and Best Solution.
- **Region:** Fetch details using `client.DescribeRegion`. Display fields: Region ID, Started Time, Current Time, Started Status, Number of Solutions, and Best Solution.

### 3. Testing Strategy
- Create functional tests using a mock gRPC server if possible, or at least unit tests for the command logic by abstracting the client.
- Since the project prefers avoiding mocks, I will explore if I can use a real connection to a local test server if available, but for now, I will focus on unit testing the command parsing and execution flow.

## Classes to be modified
- `go.mod` (add `tablewriter`)
- `commands/describe/cmd.go`
- `commands/list/cmd.go`
- `commands/output/formatter.go` (new file for shared formatting logic)

## Acceptance Criteria
- `doptctl list agents` outputs a clean, readable table.
- `doptctl list regions` outputs a clean, readable table.
- `doptctl describe agent <id>` fetches and displays agent details in a readable format.
- `doptctl describe region <id>` fetches and displays region details in a readable format.
- Graceful error handling for non-existent IDs.
