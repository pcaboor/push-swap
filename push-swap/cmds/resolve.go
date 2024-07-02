package cmds

import (
	"fmt"
	"reflect"
	"sort"
)

var Ordered []int

func OrderReference() {
	Ordered = make([]int, len(Stacks.A))
	copy(Ordered, Stacks.A)
	sort.Ints(Ordered)
}

func isSorted() bool {
	return reflect.DeepEqual(Stacks.A, Ordered)
}

func isBSorted() bool {
	sortedB := make([]int, len(Stacks.B))
	copy(sortedB, Stacks.B)
	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i] > sortedB[j]
	})
	return reflect.DeepEqual(Stacks.B, sortedB)
}

func areBothSorted() bool {
	sortedA, sortedB := make([]int, len(Stacks.A)), make([]int, len(Stacks.B))
	copy(sortedA, Stacks.A)
	copy(sortedB, Stacks.B)
	sort.Ints(sortedA)
	sort.Slice(sortedB, func(i, j int) bool {
		return sortedB[i] > sortedB[j]
	})
	return reflect.DeepEqual(Stacks.A, sortedA) && reflect.DeepEqual(Stacks.B, sortedB)
}

func Resolve() {
	// Create a reference ordered to know the final state we need.
	OrderReference()
	c := 1
	for !isSorted() {
		// Show evolution of ints in stacks.
		//DisplayResolve()
		//Check if we can use Rra
		if len(Stacks.A) > 2 {
			doRra := true
			for i := 0; i < len(Stacks.A); i++ {
				if Stacks.A[len(Stacks.A)-1] > Stacks.A[i] {
					doRra = false
					break
				}
			}
			if doRra {
				Stacks.Rra()
				//fmt.Print(c)
				fmt.Println("rra")
				c++
				continue
			}
		}
		//Check if we can use Ra
		if len(Stacks.A) > 2 {
			doRa := true
			for i := 0; i < len(Stacks.A); i++ {
				if Stacks.A[0] < Stacks.A[i] {
					doRa = false
					break
				}
			}
			if doRa {
				Stacks.Ra()
				//fmt.Print(c)
				fmt.Println("ra")
				c++
				continue
			}
		}
		// Check if we can use Ss.
		if len(Stacks.A) > 1 && len(Stacks.B) > 1 && Stacks.A[0] > Stacks.A[1] && Stacks.B[0] < Stacks.B[1] {
			Stacks.Ss()
			//fmt.Print(c)
			fmt.Println("ss")
			c++
			continue
		}
		// Check if we can use Sa.
		if len(Stacks.A) > 1 && Stacks.A[0] > Stacks.A[1] {
			Stacks.Sa()
			//fmt.Print(c)
			fmt.Println("sa")
			c++
			continue
		}
		// Check if we can use Sb.
		if len(Stacks.B) > 1 && Stacks.B[0] < Stacks.B[1] {
			Stacks.Sb()
			//fmt.Print(c)
			fmt.Println("sb")
			c++
			continue
		}
		//Check if we can use Pb.
		executed := false
		for i := 0; i+1 < len(Stacks.A); i++ {
			if Stacks.A[i] > Stacks.A[i+1] {
				Stacks.Pb()
				//fmt.Print(c)
				fmt.Println("pb")
				c++
				executed = true
				break
			}
		}

		if executed {
			continue
		}
		//Check if we can use Pa.
		if len(Stacks.B) > 0 {
			if areBothSorted() {
				Stacks.Pa()
				//fmt.Print(c)
				fmt.Println("pa")
				c++
				continue
			}
		}
		//Check if we can use Rb
		if len(Stacks.B) > 2 && !isBSorted() {
			if Stacks.B[len(Stacks.B)-1] < Stacks.B[0] {
				Stacks.Rb()
				//fmt.Print(c)
				fmt.Println("rb")
				c++
				continue
			}
		}
		//Check if we can use Rrb
		if len(Stacks.B) > 2 {
			if Stacks.B[len(Stacks.B)-1] > Stacks.B[0] {
				Stacks.Rrb()
				//fmt.Print(c)
				fmt.Println("rrb")
				c++
				continue
			}
		}
	}
}
