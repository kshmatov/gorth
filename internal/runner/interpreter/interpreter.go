package interpreter

import (
	"fmt"

	"github.com/kshmatov/gorth/internal/lexer"
	"github.com/kshmatov/gorth/internal/op"
	"github.com/kshmatov/gorth/internal/stack"
	"github.com/pkg/errors"
)

func Simulate(prog lexer.Program, debug bool) error {
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
	l := len(prog)
	for {
		if i >= l {
			break
		}
		operation = prog[i]
		if debug {
			fmt.Printf("%v\t%v\n", i, operation)
		}
		next, err := operation.Exec(s)
		if err != nil {
			return errors.Wrapf(err, "%v on <%v> runtime error", operation, i)
		}
		i += next
	}
	if debug {
		fmt.Printf("stack:\n%v\n", s)
	}
	return nil
}
