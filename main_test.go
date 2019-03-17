package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	expected := "No commands currently configured"

	cmd := exec.Command("cli", "help")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed test help thrown %+v", err)
	}

	printout := out.String()
	if !strings.Contains(printout, expected) {
		t.Fatalf("Incorrect print out expected %v but got %v", expected, printout)
	}
}
