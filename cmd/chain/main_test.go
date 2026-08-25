package main

import "testing"

func TestHelpers(t *testing.T) {
	if commandName(nil) != "show" {
		t.Fatal()
	}
	if len(splitInput("a b")) != 2 {
		t.Fatal()
	}
}
