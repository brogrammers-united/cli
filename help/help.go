package help

import "fmt"

// Help Command for showing docs
type Help struct{}

// Execute will print the help documentation for all commands
func (h Help) Execute(notify chan<- bool, args ...string) {
	fmt.Println("No commands currently configured")
	notify <- true
}
