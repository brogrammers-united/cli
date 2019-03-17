package commands

// Command is the stuct to conatin all information about a sub command of the cli
//
// All built in commands must be of type Command
// NOTE: This struct is undergoing heavy changes and will need to change to add more
// information about the commands. Once the built in commands are done the struct will
// not change and would cause major version change if changed as it will break all commands
type Command struct {
	Action func(chan<- bool, ...string) // Action to be performed when command is invoked by name
	Name   string                       // Name of subcommand for invoking the action
	Help   string                       // The usage string to explain how to use subcommand
}
