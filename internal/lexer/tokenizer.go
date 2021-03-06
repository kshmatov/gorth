package lexer

import (
	"fmt"
	"strings"
)

type token struct {
	i    int    // line number where token was placed
	v    string // token
	next *token // next token
}

// String for debug purpose
func (t token) String() string {
	return fmt.Sprintf("%v: %v -> %v", t.i, t.v, t.next)
}

// tokenize split one line to tokens
func tokenize(i int, s string, first *token) (*token, *token, int) {
	t := strings.Fields(s)
	var cur *token
	items := 0
	for _, v := range t {
		if v != "" {
			if first == nil {
				first = &token{
					i: i, v: v,
				}
			} else {
				if cur == nil {
					cur = first
				}
				cur.next = &token{
					i: i, v: v,
				}
				cur = cur.next
			}
			items++
		}
	}
	return first, cur, items
}

// string2Tokens split incoming programm to lines and form list of tokens
func string2Tokens(str string) (*token, int) {
	lines := strings.Split(str, "\n")
	var start, cur *token
	cnt := 0
	for i, s := range lines {
		f, c, ls := tokenize(i+1, s, cur)
		if start == nil {
			start = f
		}
		if c == nil {
			c = f
		}
		cur = c
		cnt += ls
	}
	return start, cnt
}
