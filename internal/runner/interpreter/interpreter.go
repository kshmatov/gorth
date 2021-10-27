package interpreter

import (
	"fmt"
	"os"

	"github.com/kshmatov/gorth/internal/lexer"
	"github.com/kshmatov/gorth/internal/op"
	"github.com/kshmatov/gorth/internal/stack"
	"github.com/pkg/errors"
)

func Simulate(prog lexer.Program, debug bool) (err error) {
	var i int
	var operation op.Op
	var s *stack.Stack

	defer func() {
		if e := recover(); e != nil {
			err = errors.Wrapf(e.(error), "line [%v], %v", i, operation)
		}
	}()

	s = stack.NewStack()
	l := len(prog)
	next := 0
	for {
		if i >= l {
			break
		}
		operation = prog[i]
		if debug {
			debugOut(fmt.Sprintf("%04x\t%v\t%v\n", i, operation, s))
		}
		next, err = operation.Exec(s)
		if err != nil {
			return errors.Wrapf(err, "%v on <%v> runtime error", operation, i)
		}
		i += next
	}
	if debug {
		debugOut(fmt.Sprintf("stack:%v\n", s))
	}
	return nil
}

func debugOut(s string) {
	_, err := fmt.Fprint(os.Stderr, s)
	if err != nil {
		fmt.Print(s)
	}
}
