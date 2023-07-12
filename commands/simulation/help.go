package simulation

import "fmt"

func Help() {
	fmt.Println(`Manage simulation actor
	
This command can start, stop, show status and extract data from current simulation. Simulation can be ready, running, 
stopped, finished.

Example:
	# Show simulation status
	doptctl simulation status

	# Start new simulation from configuration file
	doptctl simulation start simulation.conf

	# Stop current simulation
	doptctl simulation stop
	
	# Extract data from simulation to dir
	doptctl simulation extractData ~/Doptimas/Simulation/Data
	`)
}
