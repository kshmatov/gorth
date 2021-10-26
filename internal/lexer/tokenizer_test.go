package lexer

import (
	"testing"
)

func TestTokenizeEmpty(t *testing.T) {
	str := " \t\n"
	f, c := tokenize(1, str, nil)
	if f != nil {
		t.Errorf("First shoud be nil, %#v", f)
	}
	if c != nil {
		t.Errorf("Last shoud be nil, %#v", c)
	}

}
