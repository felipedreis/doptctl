package context

import "fmt"

func Help() {
	fmt.Println(`Manage client context. 
	
Client context stores the connection data to d-optimas instance. You must have at least one connection configured
and set as current connection. To create a new context one may call "doptctl context configure".
You can see all configure contexts with "doptctl context list". To change the current context call 
"doptctl context set".

Examples:
	# Prompts user to context data
	doptctl context configure

	# Prompts user to context data and set current context
	doptctl context configure --set-current

	# Creates new context with the given data
	doptctl context configure --name=local --host=localhost --port=8080

	# List context data
	doptctl context list
	
	# Change current context
	doptctl context set prod
	`)
}
