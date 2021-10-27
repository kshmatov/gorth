package lexer

import (
	"testing"
)

func TestTokenizeEmpty(t *testing.T) {
	str := " \t\n"
	f, c, i := tokenize(1, str, nil)
	if f != nil {
		t.Errorf("First shoud be nil, %#v", f)
	}
	if c != nil {
		t.Errorf("Last shoud be nil, %#v", c)
	}
	if i != 0 {
		t.Errorf("Count should be 0, but %v", i)
	}

	str = "  a \tNN    \t qqqq\n"
	f, c, i = tokenize(1, str, nil)
	if i != 3 {
		t.Errorf("Count should be 3, but %v", i)
	}

	if f == nil {
		t.Errorf("First should be 'a', but nil")
	}
	if f.v != "a" {
		t.Errorf("First should be 'a', but '%v'", f.v)
	}
	if f.i != 1 {
		t.Errorf("First should be on line 1, but %v", f.i)
	}
	n := f.next
	if n == nil {
		t.Errorf("next should be 'NN', but nil")
	}
	if n.v != "NN" {
		t.Errorf("next should be 'NN', but '%v'", n.v)
	}
	if n.i != 1 {
		t.Errorf("next should be on line 1, but %v", n.i)
	}
	if n.next != c {
		t.Errorf("Third should be equal to last, but '%#v' and '%#v'", n.next, c)
	}
	if c == nil {
		t.Errorf("last should be 'qqqq', but nil")
	}
	if c.v != "qqqq" {
		t.Errorf("last should be 'qqqq', but '%v'", c.v)
	}
	if c.i != 1 {
		t.Errorf("last should be on line 1, but %v", c.i)
	}
}

func TestLexer(t *testing.T) {
	str := " \t\n"
	f, cnt := string2Tokens(str)
	if f != nil {
		t.Errorf("First shoud be nil, %#v", f)
	}
	if cnt != 0 {
		t.Errorf("Count should be 0, but %v", cnt)
	}

	str = "  a \tNN   \n \t qqqq\n"
	ctrl := []string{"a", "NN", "qqqq"}
	cline := []int{1, 1, 2}
	f, cnt = string2Tokens(str)
	c := f
	if cnt != 3 {
		t.Errorf("Count should be 3, but %v", cnt)
	}
	for i := 0; i < 3; i++ {
		if c == nil {
			t.Errorf("Item should be value, but nil")
		}
		if c.v != ctrl[i] {
			t.Errorf("Value should be '%v' but '%v'", ctrl[i], c.v)
		}
		if c.i != cline[i] {
			t.Errorf("Line should be '%v' but '%v'", cline[i], c.i)
		}
		c = c.next
	}
}

func TestToken2Op(t *testing.T) {
	s := "\n"
	op, err := getOperation(s)
	if err == nil {
		t.Errorf("Should be err, but op is %#v %T", op, op)
	}

	s = "+"
	op, err = getOperation(s)
	if err != nil {
		t.Errorf("Error should be nil, but op is %#v", err)
	}
	if op == nil {
		t.Errorf("op should be Op, but is nill")
	}
}

func TestTokens2Program(t *testing.T) {
	str := " 12\n 25 fg + 4\n"
	p, err := Text2Program([]byte(str))
	if err == nil {
		t.Errorf("Error should be not nil, but prog: <%#v>", p)
	}

	str = "15 7\n+\n45\t1 - +"
	p, err = Text2Program([]byte(str))
	if err != nil {
		t.Errorf("Error should be nil, but %v", err)
	}
	if len(p) != 7 {
		t.Errorf("Program should have 7 operands, but has %v", len(p))
	}
}
