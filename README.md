# doptctl

`doptctl` is a command-line interface (CLI) tool for managing and inspecting **D-Optimas** simulations. It provides commands for configuring connection contexts, listing entities, and controlling simulations via gRPC.

This project is closely connected to the [d-optimas](https://github.com/felipedreis/d-optimas) project.

## Code Intelligence with GitNexus

This project is indexed by **GitNexus**. You can use GitNexus MCP tools to:
- Analyze dependencies and understand the relationship with the `d-optimas` core.
- Perform impact analysis before making changes.
- Navigate the codebase using semantic queries and execution flows.

If you are using an AI assistant with MCP support, you can leverage tools like `gitnexus_impact`, `gitnexus_query`, and `gitnexus_context` for deep codebase analysis.

## Usage

### Context Management
```shell
doptctl context configure
doptctl context list
doptctl context set contextName
```

### Inspection
```shell
doptctl list agents
doptctl list regions

doptctl describe agent agentId
doptctl describe region regionId

doptctl benchmark status
```

### Simulation Control
```shell
doptctl simulation start simulationFile
doptctl simulation stop
doptctl simulation status
doptctl simulation extractData 
```

### Monitoring & Server
```shell
doptctl watch agent agentId
doptctl watch region regionId

doptctl server restart
```
