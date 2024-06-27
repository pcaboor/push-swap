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
	for !isSorted() && c < 10 {
		// Show evolution of ints in stacks.
		DisplayResolve()
		// Check if we can use Sa.
		if len(Stacks.A) > 1 && Stacks.A[0] > Stacks.A[1] {
			Stacks.Sa()
			fmt.Println(string(c+'0') + " - sa")
			c++
			continue
		}
		//Check if we can use Pb.
		for i := 0; i+1 < len(Stacks.A); i++ {
			if Stacks.A[i] > Stacks.A[i+1] {
				for j := 0; j < i; j++ {
					Stacks.Pb()
					fmt.Println(string(c+'0') + " - pb")
					c++
				}
				break
			}
		}
		//Check if we can use Pa.
		if len(Stacks.B) > 0 {
			if areBothSorted() {
				for i := 0; len(Stacks.B) > 0; i++ {
					Stacks.Pa()
					fmt.Println(string(c+'0') + " - pa")
					c++
				}
			}
		}
	}
}
