package cmds

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

var Stacks StacksObject
