package main

import (
	"fmt"
	"os"

	"github.com/brogrammers-united/cli/help"
)

// Command interface specifies a command that can be executed by the cli
//
// Main cli will block until notified through the channel. This allows the command to
// either block the main thread or spin up background processes and return execution to the main
// thread immediately.
//
// execute will be called with all args passed to the program after the first
// ex. cli <command name> <arg1> <arg2> arg 1 - n will be passed in the execute params
type Command interface {
	Execute(notify chan<- bool, args ...string)
}

// Creates a map of commands access map by command name to receive
// the struct satisfying the Command interface with the functionality
// of the given command
func createCommandMap() (commands map[string]Command) {
	commands = make(map[string]Command)
	commands["help"] = help.Help{}
	return commands
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Sub command needs to be specified")
	}

	commands := createCommandMap()
	command := commands[os.Args[1]]

	if command == nil {
		fmt.Println("Command not found")
		os.Exit(1)
	}

	var args []string
	if len(os.Args) > 1 {
		args = os.Args[2:]
	}

	waiting := make(chan bool)
	go command.Execute(waiting, args...)
	<-waiting
}
