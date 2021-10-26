package op

import (
	"github.com/kshmatov/gorth/internal/stack"
)

type Ops uint32

type Program []Op

type Op interface {
	Exec(stack *stack.Stack) error
}

type op struct {
	desc string
	f    func(s *stack.Stack) error
}

func (o op) Exec(s *stack.Stack) error {
	return o.f(s)
}

func (o op) String() string {
	return o.desc
}
