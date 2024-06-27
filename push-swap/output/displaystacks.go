package output

import (
	"fmt"
	"input"
)

// Display both stacks.
func DisplayStacks() {
	cA, cB := 0, 0
	linesSize := 0
	if len(input.Stacks.A) >= len(input.Stacks.B) {
		linesSize = len(input.Stacks.A)
	} else {
		linesSize = len(input.Stacks.B)
	}
	for i := 0; i < linesSize; i++ {
		if len(input.Stacks.A) < linesSize-i {
			fmt.Print(".")
			cA++
		} else {
			fmt.Print(input.Stacks.A[i-cA])
		}
		fmt.Print("\t|\t")
		if len(input.Stacks.B) < linesSize-i {
			fmt.Print(".")
			cB++
		} else {
			fmt.Print(input.Stacks.B[i-cB])
		}
		fmt.Println()
	}
	fmt.Println("Stack A\t|StackB")
}
