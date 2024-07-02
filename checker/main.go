package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StacksObject struct {
	A, B []int
}

func (s *StacksObject) Pa() {
	if len(s.B) > 0 {
		s.A = append([]int{s.B[0]}, s.A...)
		s.B = s.B[1:]
	}
}

func (s *StacksObject) Pb() {
	if len(s.A) > 0 {
		s.B = append([]int{s.A[0]}, s.B...)
		s.A = s.A[1:]
	}
}

func (s *StacksObject) Sa() {
	if len(s.A) > 1 {
		s.A[0], s.A[1] = s.A[1], s.A[0]
	}
}

func (s *StacksObject) Sb() {
	if len(s.B) > 1 {
		s.B[0], s.B[1] = s.B[1], s.B[0]
	}
}

func (s *StacksObject) Ss() {
	s.Sa()
	s.Sb()
}

func (s *StacksObject) Ra() {
	if len(s.A) > 1 {
		s.A = append(s.A[1:], s.A[0])
	}
}

func (s *StacksObject) Rb() {
	if len(s.B) > 1 {
		s.B = append(s.B[1:], s.B[0])
	}
}

func (s *StacksObject) Rr() {
	s.Ra()
	s.Rb()
}

func (s *StacksObject) Rra() {
	if len(s.A) > 1 {
		s.A = append([]int{s.A[len(s.A)-1]}, s.A[:len(s.A)-1]...)
	}
}

func (s *StacksObject) Rrb() {
	if len(s.B) > 1 {
		s.B = append([]int{s.B[len(s.B)-1]}, s.B[:len(s.B)-1]...)
	}
}

func (s *StacksObject) Rrr() {
	s.Rra()
	s.Rrb()
}

var Stacks StacksObject

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./checker \"<Int list>\"")
		return
	}

	err := checkDuplicates(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}

		instruction := strings.TrimSpace(scanner.Text())
		if instruction == "" {
			break
		}

		fmt.Printf("Executing instruction: %s\n", instruction)
		fmt.Printf("Before - A: %v, B: %v\n", Stacks.A, Stacks.B)

		if err := resolve(instruction); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			fmt.Println("KO")
			return
		}

		fmt.Fprintf(os.Stderr, "After - A: %v, B: %v\n", Stacks.A, Stacks.B)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		fmt.Println("KO")
		return
	}

	if isSorted(Stacks) && len(Stacks.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func checkDuplicates(args string) error {
	parts := strings.Fields(args)
	seen := make(map[int]bool)

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("invalid argument: %s", part)
		}
		if seen[num] {
			return fmt.Errorf("duplicate found: %d", num)
		}
		seen[num] = true
		Stacks.A = append(Stacks.A, num)
	}

	return nil
}

func resolve(instruction string) error {
	switch strings.ToLower(instruction) {
	case "sa":
		Stacks.Sa()
	case "sb":
		Stacks.Sb()
	case "ss":
		Stacks.Ss()
	case "pa":
		Stacks.Pa()
	case "pb":
		Stacks.Pb()
	case "ra":
		Stacks.Ra()
	case "rb":
		Stacks.Rb()
	case "rr":
		Stacks.Rr()
	case "rra":
		Stacks.Rra()
	case "rrb":
		Stacks.Rrb()
	case "rrr":
		Stacks.Rrr()
	default:
		return fmt.Errorf("unknown instruction: %s", instruction)
	}
	return nil
}

/*
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
*/

func isSorted(s StacksObject) bool {
	for i := 1; i < len(s.A); i++ {
		if s.A[i] < s.A[i-1] {
			return false
		}
	}
	return true
}
