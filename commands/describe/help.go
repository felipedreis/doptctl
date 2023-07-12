package describe

import "fmt"

func Help() {
	fmt.Println(`Describe simulation agents and regions.
	
Describe agent or region configuration, produced or received solutions, current status.

Example:
	# Describe agent
	doptctl describe agent agent-1

	# Describe region
	doptctl describe region region-1
	`)
}
