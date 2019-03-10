package help

import (
	"fmt"

	"github.com/brogrammers-united/cli/commands"
)

// GetCommand must be built for the cli to add the command to list
func GetCommand() commands.Command {
	return helpCommand
}

// Help is a sub command to print documentation for the built in commands
type Help struct{}

var helpCommand = Help{}

// Execute will print the help documentation for all commands
// TODO: Implement the printing of all commands in the local env
func (h Help) Execute(notify chan<- bool, args ...string) {
	fmt.Println("No commands currently configured")
	notify <- true
}

// Describe will describe what the command does
// TODO: Add description of the command once the Execute func is built
func (h Help) Describe(printLong bool) {
	fmt.Println("No current implementation of helping documentation")
}
