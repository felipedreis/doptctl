package benchmark

import "fmt"

func Help() {
	fmt.Println(`Manage benchmark actor.
	
Examples:
	# Prints the benchmark related data, such as current benchmarking problem, number of function evaluations
	doptctl benchmark status
	`)
}
