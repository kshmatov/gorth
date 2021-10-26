package lexer

import (
	"strings"
)

type token struct {
	i    int
	v    string
	next *token
}

func tokenize(i int, s string, first *token) (*token, *token) {
	t := strings.Fields(s)
	var cur *token
	for _, v := range t {
		if v != "" {
			if first == nil {
				first = &token{
					i: i, v: v,
				}
				cur = first
			} else {
				cur.next = &token{
					i: i, v: v,
				}
				cur = cur.next
			}
		}
	}
	return first, cur
}

func liner(s string) *token {
	lines := strings.Split(s, "\n")
	var start, cur *token
	for i, s := range lines {
		f, c := tokenize(i, s, cur)
		if start == nil {
			start = f
		}
		cur = c
	}
	return start
}
