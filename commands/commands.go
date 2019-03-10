package commands

// FOR A PACKAGE TO PROVIDE A COMMAND
//
// By convention all packages adding a command need to have a method getCommand
// that returns an instance of the command the package implements

// Command is a subcommand that can be run from the main application
type Command interface {
	Executable
	Describable
}

// Executable interface specifies a command that can be executed by the cli
//
// Main cli will block until notified through the channel. This allows the command to
// either block the main thread or spin up background processes and return execution to the main
// thread immediately.
//
// execute will be called with all args passed to the program after the first
// ex. cli <command name> <arg1> <arg2> arg 1 - n will be passed in the execute params
type Executable interface {
	Execute(notify chan<- bool, args ...string)
}

// Describable is the interface used to retrieve information about each command
//
// Should print information in given format with no extra leading or ending lines
// extra lines will screw the internal styleing of menus.
type Describable interface {
	Describe(printLong bool)
}
