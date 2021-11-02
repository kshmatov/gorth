package programm

import (
	"github.com/kshmatov/gorth/internal/programm/op"
	"github.com/kshmatov/gorth/internal/programm/returnStack"
	"github.com/kshmatov/gorth/internal/programm/stack"
	"github.com/pkg/errors"
)

type Subroutine []op.Op

type Programm struct {
	body    map[string]Subroutine
	stack   stack.Stack
	returns returnStack.ReturnStack
}

func (p Programm) Size() int {
	i := 0
	for _, v := range p.body {
		i += len(v)
	}
	return i
}

func (p *Programm) AddOperation(token string) error {
	return errors.New("not implemented")
}

func (p *Programm) Run(out chan<- string) error {
	close(out)
	return errors.New("not implemented")
}
