package op

import (
	"github.com/kshmatov/gorth/internal/stack"
)

type operation struct {
	desc string
	f    func(s *stack.Stack) (int, error)
}

func (o operation) Exec(s *stack.Stack) (int, error) {
	return o.f(s)
}

func (o operation) String() string {
	return o.desc
}
