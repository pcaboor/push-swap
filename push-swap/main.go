package main

import (
	"cmds"
	"input"
	"output"
)

func main() {
	input.GetDigits()
	cmds.Resolve()
	output.DisplayStacks()
}
