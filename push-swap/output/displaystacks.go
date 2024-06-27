package output

import (
	"cmds"
	"fmt"
)

// Display both stacks.
func DisplayStacks() {
	cA, cB := 0, 0
	linesSize := 0
	if len(cmds.Stacks.A) >= len(cmds.Stacks.B) {
		linesSize = len(cmds.Stacks.A)
	} else {
		linesSize = len(cmds.Stacks.B)
	}
	for i := 0; i < linesSize; i++ {
		if len(cmds.Stacks.A) < linesSize-i {
			fmt.Print(".")
			cA++
		} else {
			fmt.Print(cmds.Stacks.A[i-cA])
		}
		fmt.Print("\t|\t")
		if len(cmds.Stacks.B) < linesSize-i {
			fmt.Print(".")
			cB++
		} else {
			fmt.Print(cmds.Stacks.B[i-cB])
		}
		fmt.Println()
	}
	fmt.Println("Stack A\t|StackB")
}
