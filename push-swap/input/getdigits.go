package input

import (
	"cmds"
	"fmt"
	"os"
	"strconv"
)

// Fill stack "a" with digits from first argument.
func GetDigits() {
	CheckArgs()
	toConvert := ""
	for _, digit := range os.Args[1] {
		if digit >= '0' && digit <= '9' {
			toConvert += string(digit)
		} else {
			convertedDigits, err := strconv.Atoi(toConvert)
			if err != nil {
				fmt.Println(err)
			}
			cmds.Stacks.A = append(cmds.Stacks.A, convertedDigits)
			toConvert = ""
		}
	}
	if toConvert != "" {
		convertedDigits, err := strconv.Atoi(toConvert)
		if err != nil {
			fmt.Println(err)
		}
		cmds.Stacks.A = append(cmds.Stacks.A, convertedDigits)
	}
	CheckDuplicates()
}
