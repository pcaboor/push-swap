package input

import (
	"cmds"
	"fmt"
	"os"
)

func CheckDuplicates() {
	for i, dig1 := range cmds.Stacks.A {
		for j, dig2 := range cmds.Stacks.A {
			if j != i && dig1 == dig2 {
				fmt.Println("Error")
				os.Exit(0)
			}
		}
	}
}
