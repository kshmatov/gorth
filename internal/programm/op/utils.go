package op

import (
	"github.com/kshmatov/gorth/internal/programm/stack"
	"github.com/pkg/errors"
)

func get2(s *stack.Stack) (int64, int64, error) {
	a, err := s.Pop()
	if err != nil {
		return 0, 0, errors.Wrap(err, "on first argument")
	}
	b, err := s.Pop()
	if err != nil {
		return 0, 0, errors.Wrap(err, "on second argument")
	}
	return a, b, nil
}

func cond(b bool, s *stack.Stack) {
	if b {
		s.Push(True)
	} else {
		s.Push(False)
	}
}
