package cmds

import (
	"fmt"
)

// Display both stacks.
func DisplayResolve() {
	cA, cB := 0, 0
	linesSize := 0
	if len(Stacks.A) >= len(Stacks.B) {
		linesSize = len(Stacks.A)
	} else {
		linesSize = len(Stacks.B)
	}
	for i := 0; i < linesSize; i++ {
		if len(Stacks.A) < linesSize-i {
			fmt.Print(".")
			cA++
		} else {
			fmt.Print(Stacks.A[i-cA])
		}
		fmt.Print("\t|\t")
		if len(Stacks.B) < linesSize-i {
			fmt.Print(".")
			cB++
		} else {
			fmt.Print(Stacks.B[i-cB])
		}
		fmt.Println()
	}
	fmt.Println("Stack A\t|StackB")
}
