package lexer

import (
	"github.com/pkg/errors"
)

var ErrEmptyProg = errors.New("empty text")

// var ErrSyntax = errors.New("syntax error")

func Text2Program(t []byte) (*token, error) {
	token, _ := string2Tokens(string(t))
	if token == nil {
		return nil, ErrEmptyProg
	}
	return token, nil
}
