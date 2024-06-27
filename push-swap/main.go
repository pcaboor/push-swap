package main

import (
	"cmds"
	"input"
	"output"
)

func main() {
	input.CheckArgs()
	input.GetDigits()
	// Cmds tests with go run . "1 2 4 5 34 7 4"
	cmds.Stacks.Pb()
	cmds.Stacks.Pb()
	cmds.Stacks.Pb()
	cmds.Stacks.Pb()
	cmds.Stacks.Pa()
	cmds.Stacks.Sa()
	cmds.Stacks.Pb()
	cmds.Stacks.Sb()
	cmds.Stacks.Ss()
	cmds.Stacks.Ra()
	cmds.Stacks.Rb()
	cmds.Stacks.Rr()
	cmds.Stacks.Rra()
	cmds.Stacks.Rrb()
	cmds.Stacks.Rrr()
	// End.
	output.DisplayStacks()
}
