package interpreter

import (
	"fmt"

	"github.com/kshmatov/gorth/internal/op"
	"github.com/kshmatov/gorth/internal/stack"
	"github.com/pkg/errors"
)

func Simulate(prog op.Program, debug bool) error {
	var i int
	var operation op.Op
	var s *stack.Stack

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Exception: %v\n", e)
			if operation != nil {
				fmt.Printf("on [%v]\t%v\n", i, operation)
			}
			if s != nil {
				fmt.Printf("stack:\n%v\n", s)
			}
		}
	}()

	s = stack.NewStack()
	for i, operation = range prog {
		if debug {
			fmt.Printf("%v\t%v\n", i, operation)
		}
		err := operation.Exec(s)
		if err != nil {
			return errors.Wrapf(err, "%v on <%v> runtime error", operation, i)
		}
	}
	if debug {
		fmt.Printf("stack:\n%v\n", s)
	}
	return nil
}
