package commands

import "fmt"

func Help() {
	fmt.Println(`doptctl is a command-line tool for managing and inspecting d-optimas simulation. If the simulation is running 
in benchmark-mode we can also inspect the problem-set data as well.

The available sub-commands are:
	context 			Managing client connections with different d-optimas instances
	benchmark			Inspect benchmark problem-set data, function evaluation
	describe			Inspect agents and regions data
	list				List running agents and regions 
	simulation			Configure, start, stop and extract simulation data
	help				Show this help or subcommand help
	`)
}
