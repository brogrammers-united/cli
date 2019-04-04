package main

import (
	"fmt"
	"os"
	"strings"

	commandLib "github.com/brogrammers-united/cli/commands"
	help "github.com/brogrammers-united/cli/help"
)

// For ease of extending the internal commands simply import and add to array
//
// this is to decouple the name and action form this package and is fully
//  configured in the package for each subcommand the name will be pulled
// from the struct and add to the map
var builtInCommands = []commandLib.Command{help.HelpCommand}

// Returns a map of all comfigured commands
//
// For now this funciton is hard coded with the build in commands
// this can be exteded in later iterations to include custom execution
func createCommandMap() (commands map[string]commandLib.Command) {
	commands = make(map[string]commandLib.Command)
	for _, command := range builtInCommands {
		commands[command.Name] = command
	}
	return commands
}

func main() {
	GetAccessToken()

	if len(os.Args) <= 1 {
		fmt.Println("Sub command needs to be specified")
		os.Exit(1)
	}

	commands := createCommandMap()
	command := commands[strings.ToLower(os.Args[1])]

	// Since comparing stuct need something that can be compared
	// The main part of the stuct that matters is the action if not
	// nil then we may continue
	if command.Action == nil {
		fmt.Println("Command not found")
		os.Exit(1)
	}

	var args []string
	if len(os.Args) > 1 {
		args = os.Args[2:]
	}

	waiting := make(chan bool)
	go command.Action(waiting, args...)
	<-waiting
}
