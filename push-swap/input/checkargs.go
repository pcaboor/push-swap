package input

import (
	"fmt"
	"os"
)

// Handle errors if no digit or non-digit characters are found.
func CheckArgs() {
	if len(os.Args) == 1 {
		fmt.Println()
		os.Exit(0)
	}
	for _, digit := range os.Args[1] {
		if !(digit >= '0' && digit <= '9') &&
			digit != ' ' {
			fmt.Println("Error")
			os.Exit(0)
		}
	}
}
