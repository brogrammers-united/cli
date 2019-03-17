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

func TestNoCommand(t *testing.T) {
	expected := "Sub command needs to be specified"
	cmd := exec.Command("cli")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		t.Fatalf("Failed to throw error when no sub command specified")
	}

	printout := out.String()
	if !strings.Contains(printout, expected) {
		t.Fatalf("Incorrect print out expected %v but got %v", expected, printout)
	}
}
