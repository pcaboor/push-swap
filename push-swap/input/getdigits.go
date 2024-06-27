package input

import (
	"fmt"
	"os"
	"strconv"
)

type StacksObject struct {
	A, B []int
}

var Stacks StacksObject

// Fill stack "a" with digits from first argument.
func GetDigits() {
	toConvert := ""
	for _, digit := range os.Args[1] {
		if digit >= '0' && digit <= '9' {
			toConvert += string(digit)
		} else {
			convertedDigits, err := strconv.Atoi(toConvert)
			if err != nil {
				fmt.Println(err)
			}
			Stacks.A = append(Stacks.A, convertedDigits)
			toConvert = ""
		}
	}
	if toConvert != "" {
		convertedDigits, err := strconv.Atoi(toConvert)
		if err != nil {
			fmt.Println(err)
		}
		Stacks.A = append(Stacks.A, convertedDigits)
	}
}
