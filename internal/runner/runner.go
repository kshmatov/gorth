package runner

import (
	"os"

	"github.com/kshmatov/gorth/internal/lexer"
	"github.com/kshmatov/gorth/internal/runner/interpreter"
	"github.com/pkg/errors"
)

func Run(fn string, debug bool) error {
	code, err := os.ReadFile(fn)
	if err != nil {
		return errors.Wrapf(err, "open %v", fn)
	}
	prog, err := lexer.Text2Program(code)
	if err != nil {
		return err
	}
	return interpreter.Simulate(prog, debug)
}
