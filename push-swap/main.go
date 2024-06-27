package main

import (
	"cmds"
	"input"
	"output"
)

func main() {
	input.CheckArgs()
	input.GetDigits()
	cmds.Stacks.Pb()
	cmds.Stacks.Pb()
	cmds.Stacks.Pa()
	output.DisplayStacks()
}
