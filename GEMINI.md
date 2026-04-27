# GEMINI.md - Project Context

## Project Overview
`doptctl` is a command-line interface (CLI) tool written in Go, designed to interact with the **D-Optimas** simulation and optimization engine. It is a companion tool for the [d-optimas](https://github.com/felipedreis/d-optimas) project, communicating with the server using gRPC. The tool allows users to manage connection contexts, list and describe simulation entities (agents, regions), and control simulations.

## GitNexus & D-Optimas Integration
This project is indexed by **GitNexus**, providing advanced code intelligence.
- **Dependency Analysis:** Use GitNexus MCP tools to analyze how `doptctl` depends on and interacts with the `d-optimas` core.
- **Impact Analysis:** Before modifying any core logic or gRPC interfaces, run `gitnexus_impact` to assess the blast radius across both `doptctl` and `d-optimas`.
- **Navigation:** Use `gitnexus_query` and `gitnexus_context` to understand the end-to-end execution flows from the CLI to the simulation engine.

### Key Technologies
- **Language:** Go (1.20+)
- **Communication:** gRPC, Protocol Buffers (Protobuf)
- **State Management:** Local JSON files in `~/.doptctl/contexts/` for managing server connection details.
- **Intelligence:** GitNexus MCP for cross-project dependency and impact analysis.

## Architecture
The project follows a modular command-based architecture:
- `main.go`: The entry point that parses global arguments and delegates to the `commands` package.
- `commands/`: Central command dispatcher.
    - `command.go`: Defines the `Command` interface and the `Run` loop.
    - Sub-packages (e.g., `list`, `describe`, `simulation`): Implement specific command logic.
- `doptimas/api/`: Contains the gRPC client and message definitions generated from `.proto` files.

## Building and Running
### Prerequisites
- Go 1.20 or higher.
- Access to a D-Optimas gRPC server.

### Commands
- **Build:** `go build -o doptctl main.go`
- **Run:** `./doptctl [command] [subcommand] [args]`
- **Test:** `go test ./...` (TODO: Verify if tests exist)

### Common Usage
```bash
# Configure a new connection context
doptctl context configure

# Set the active context
doptctl context set <context-name>

# List available agents
doptctl list agents

# Describe a specific agent
doptctl describe agent <agent-id>
```

---
## Workflow

Before starting any implementation, create a plan document in the `plans/` directory. Use the naming convention `P###-<name>.md` (e.g., `P001-fix-ga-bugs.md`). The plan should describe:
- What the feature/fix does
- Tasks to complete
- Dependencies and prerequisites
- Design decisions, architectural considerations, and trade-offs
- Use diagrams (UML, flowcharts) where appropriate
- Update the plan during implementation to reflect completed steps and any changes.
- Avoid significant plan changes without proper consideration.
- The plans should be extensive and tracked by the version control. 

After plan, we enter in development mode. We have to create a new branch 
to commit our changes. Bellow are the rules: 

- When we pass for the implementation, we should have a test-drivevn development
approach
- Propose the test classes and the test cases first, I have to approve them first,  then we progress 
    to the implementation. 
- After the test is approved and implemented we progress to the implementation of the component
- Tests should be meaningful, and test edge cases and happy paths. 
- We should avoid mocking stuff
- Complex scenarios should be tested in a functional test 

When the development is finished, commit our changes to the branch
created earlier, and then push it to the remote repo. Open a PR to review.
The owner will review the changes. 

### GitHub Integration (`gh` CLI)
We use the GitHub CLI for managing the project's lifecycle. 

#### Creating Issues
When creating issues, use a detailed body. Using a heredoc with quotes (`<<'EOF'`) prevents the shell from interpreting special characters like backticks.

```bash
gh issue create --title "[Category] Title" --body - <<'EOF'
### Context
[Describe the problem or feature background]

### Technical Approach
[Describe the high-level strategy and detailed steps]

### Classes to be modified
* [Reference the classes that will be touched]

### Acceptance Criteria
* [List verifiable criteria that can be checked by AI or automated tests]
EOF
```

#### Creating Pull Requests
When creating pull requests, use a structured body. Using a heredoc with quotes (`<<'EOF'`) is the most robust way to handle multi-line content.

```bash
gh pr create --title "[Category] Title" --body - <<'EOF'
### Context
[Describe the problem or feature background and link to issues]

### Technical Approach
[Describe the implementation details and design decisions]

### Acceptance Criteria
- [ ] [Criterion 1]
- [ ] [Criterion 2]

Fixes #[Issue Number]
EOF
```

Common operations:
- **List issues:** `gh issue list`
- **View issue & comments:** `gh issue view <number> --comments`
- **Close issue:** `gh issue close <number>`
- **Comment on issue:** `gh issue comment <number> --body '...'`
- **List PRs:** `gh pr list`
- **View PR:** `gh pr view <number>`
- **Edit PR:** `gh pr edit <number> --body '...'`

### Code Analysis & Discovery (GitNexus)
Always use GitNexus tools to perform structured analysis of the codebase before proposing changes or creating issues.
- **Querying:** Use `mcp_gitnexus_query` to find concepts or execution flows.
- **Context:** Use `mcp_gitnexus_context` for a 360-degree view of a symbol (callers, callees, properties).
- **Impact:** Use `mcp_gitnexus_impact` to analyze the blast radius of a change.
- **Structural Discovery:** Use `mcp_gitnexus_cypher` for complex relationship queries.
- **API Impact:** Use `mcp_gitnexus_api_impact` before modifying any API route handler.

---

## Development Conventions
- **Command Implementation:** Each command should implement the `Command` interface:
  ```go
  type Command interface {
      Execute(conn *grpc.ClientConn, opts map[string]string)
  }
  ```
- **Context Handling:** Most commands require an active context (server URL) which is automatically loaded by `commands.Run` before execution, establishing the gRPC connection.
- **Error Handling:** Uses `log.Fatalf` for critical failures in the CLI flow.
- **CLI Parsing:** Currently uses a custom manual parser in `commands/command.go` instead of standard libraries like Cobra.

## Key Files
- `main.go`: Application entry point.
- `commands/command.go`: Core CLI execution logic and gRPC connection management.
- `commands/context/context.go`: Logic for persisting and loading server connection metadata.
- `doptimas/api/`: Generated gRPC code (do not edit manually).
