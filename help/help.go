package help

import (
	"fmt"

	commands "github.com/brogrammers-united/cli/commands"
)

var action = func(notify chan<- bool, args ...string) {
	fmt.Println("No commands currently configured")
	notify <- true
}

// Help command to print docs for internal commands
var HelpCommand = commands.Command{
	Action: action,
	Name:   "help",
	Help:   "This will doc for you",
}
