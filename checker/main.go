package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./checker \"<Int list>\"")
		return
	}

	stackA, err := checkDuplicates(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		fmt.Println("KO")
		return
	}

	stackB := make(Stack, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Fprintf(os.Stderr, "Enter instruction (or press [Enter] again to finish): ")
		if !scanner.Scan() {
			break
		}

		instruction := strings.TrimSpace(scanner.Text())
		if instruction == "" {
			break
		}

		fmt.Fprintf(os.Stderr, "Executing instruction: %s\n", instruction)
		fmt.Fprintf(os.Stderr, "Before - A: %v, B: %v\n", stackA, stackB)

		if err := resolve(&stackA, &stackB, instruction); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Println("KO")
			return
		}

		fmt.Fprintf(os.Stderr, "After - A: %v, B: %v\n", stackA, stackB)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		fmt.Println("KO")
		return
	}

	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func checkDuplicates(args string) (Stack, error) {
	parts := strings.Fields(args)
	stack := make(Stack, 0, len(parts))
	seen := make(map[int]bool)

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid argument: %s", part)
		}
		if seen[num] {
			return nil, fmt.Errorf("duplicate found: %d", num)
		}
		seen[num] = true
		stack = append(stack, num)
	}

	return stack, nil
}

func resolve(a, b *Stack, instruction string) error {
	switch strings.ToLower(instruction) {
	case "sa":
		return swapTop(a)
	case "sb":
		return swapTop(b)
	case "ss":
		if err := swapTop(a); err != nil {
			return err
		}
		return swapTop(b)
	case "pa":
		return push(b, a)
	case "pb":
		return push(a, b)
	case "ra":
		return rotate(a)
	case "rb":
		return rotate(b)
	case "rr":
		if err := rotate(a); err != nil {
			return err
		}
		return rotate(b)
	case "rra":
		return reverseRotate(a)
	case "rrb":
		return reverseRotate(b)
	case "rrr":
		if err := reverseRotate(a); err != nil {
			return err
		}
		return reverseRotate(b)
	default:
		return fmt.Errorf("unknown instruction: %s", instruction)
	}
}

func swapTop(s *Stack) error {
	if len(*s) < 2 {
		return fmt.Errorf("not enough elements to swap")
	}
	(*s)[len(*s)-1], (*s)[len(*s)-2] = (*s)[len(*s)-2], (*s)[len(*s)-1]
	return nil
}

func push(from, to *Stack) error {
	if len(*from) == 0 {
		return fmt.Errorf("source stack is empty")
	}
	*to = append(*to, (*from)[len(*from)-1])
	*from = (*from)[:len(*from)-1]
	return nil
}

func rotate(s *Stack) error {
	if len(*s) < 2 {
		return nil
	}
	*s = append((*s)[1:], (*s)[0])
	return nil
}

func reverseRotate(s *Stack) error {
	if len(*s) < 2 {
		return nil
	}
	*s = append(Stack{(*s)[len(*s)-1]}, (*s)[:len(*s)-1]...)
	return nil
}

func isSorted(s Stack) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
