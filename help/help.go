package help

import (
	"fmt"

	commands "github.com/brogrammers-united/cli/commands"
)

// GetCommand must be built for the cli to add the command to list
func GetCommand() commands.Command {
	return helpCommand
}

var action = func(notify chan<- bool, args ...string) {
	fmt.Println("No commands currently configured")
	notify <- true
}

var helpCommand = commands.Command{
	Action: action,
	Name:   "help",
	Help:   "This will doc for you",
}
