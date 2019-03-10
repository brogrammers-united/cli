package main

import (
	"fmt"
	"os"
	"strings"

	commandLib "github.com/brogrammers-united/cli/commands"
	help "github.com/brogrammers-united/cli/help"
)

// Creates a map of commands access map by command name to receive
// the struct satisfying the Command interface with the functionality
// of the given command
func createCommandMap() (commands map[string]commandLib.Command) {
	commands = make(map[string]commandLib.Command)
	commands["help"] = help.GetCommand()
	return commands
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Sub command needs to be specified")
	}

	commands := createCommandMap()
	command := commands[strings.ToLower(os.Args[1])]

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
