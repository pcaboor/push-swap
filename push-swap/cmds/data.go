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
