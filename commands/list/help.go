package list

import "fmt"

func Help() {
	fmt.Println(`List simulation entities
	
Shows a resumed table of entities in the simulation, being them agents or regions

Examples:

	# List running agents
	doptctl list agents
	
	# List all agents
	doptctl list agents --all 

	# List all regions
	doptctl list regions
	`)
}
