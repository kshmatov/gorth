package op

import (
	"github.com/kshmatov/gorth/internal/stack"
)

// Op is interface for operators in our program
type Op interface {
	// Exec returns offset of next step in program and error
	// Default value for offset is 1, but IFs, ELSEs and WHILEs will need other values
	Exec(stack *stack.Stack) (int, error)
}
