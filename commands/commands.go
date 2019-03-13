package commands

// FOR A PACKAGE TO PROVIDE A COMMAND
//
// By convention all packages adding a command need to have a method getCommand
// that returns an instance of the command the package implements

// Command is the stuct to conatin all information about a sub command of the cli
//
// NOTE: This struct is undergoing heavy changes and will need to change to add more
// information about the commands. Once the built in commands are done the struct will
// not change and would cause major version change if changed as it will break all commands
type Command struct {
	Action func(chan<- bool, ...string) // Action to be performed when command is invoked by name
	Name   string                       // Name of subcommand for invoking the action
	Help   string                       // The usage string to explain how to use subcommand
}
