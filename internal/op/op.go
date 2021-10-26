package op

import (
	"github.com/kshmatov/gorth/internal/stack"
)

type Ops uint32

const (
	OP_PUSH   = Ops(0)
	OP_PLUS   = Ops(iota)
	OP_DUMP   = Ops(iota)
	OP_MINUS  = Ops(iota)
	COUNT_OPS = Ops(iota)
)

type Program []Op

type Op interface {
	Exec(stack *stack.Stack) error
}
type fun struct {
	arg int64
	f   func(s *stack.Stack, i int64) error
}

func (o fun) Exec(s *stack.Stack) error {
	return o.f(s, o.arg)
}

type op struct {
	f func(s *stack.Stack) error
}

func (o op) Exec(s *stack.Stack) error {
	return o.f(s)
}
