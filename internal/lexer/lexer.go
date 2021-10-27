package lexer

import (
	"strconv"

	"github.com/kshmatov/gorth/internal/op"
	"github.com/pkg/errors"
)

var ErrEmptyProg = errors.New("empty text")
var ErrSyntax = errors.New("syntax error")

type Program []op.Op

func Text2Program(t []byte) (Program, error) {
	token, lines := string2Tokens(string(t))
	if token == nil {
		return nil, errors.New("empty text")
	}
	prog := make(Program, lines)
	i := 0
	for token != nil {
		op, err := getOperation(token.v)
		if err != nil {
			return nil, errors.Wrapf(ErrSyntax, "on line <%v>, token <%v>", token.i, token.v)
		}
		prog[i] = op
		i++
		token = token.next
	}
	return prog, nil
}

// getOperation searches for a operator to execute or tries to parse token to Int64
func getOperation(s string) (op.Op, error) {
	if v, ok := ops[s]; ok {
		return v(), nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return op.Push(v), nil
}
