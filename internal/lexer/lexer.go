package lexer

import (
	"errors"

	"github.com/kshmatov/gorth/internal/op"
)

var ops = map[string]func() op.Op{
	"+":    op.Add,
	"-":    op.Sub,
	".":    op.Dump,
	"add":  op.Add,
	"sub":  op.Sub,
	"dump": op.Dump,
}

func Tokens2Program(t []byte) (op.Op, error) {
	token := liner(string(t))
	if token == nil {
		return nil, errors.New("empty text")
	}
	for {

	}
}
