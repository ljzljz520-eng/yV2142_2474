package main

import "strings"

func commandName(args []string) string {
	if len(args) < 2 {
		return "show"
	}
	return strings.ToLower(args[1])
}
func usage() string                { return "chain [session] [idiom]" }
func splitInput(v string) []string { return strings.Fields(v) }
