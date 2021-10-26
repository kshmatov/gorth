package stack

import (
	"errors"
	"fmt"
	"strings"
)

var ErrStackIsOut = errors.New("stack is out of data")

type Stack struct {
	s   []int64
	pos int
}

func NewStack() *Stack {
	return &Stack{
		s:   nil,
		pos: -1,
	}
}

func (s *Stack) Push(i int64) {
	s.pos++
	if cap(s.s) <= s.pos {
		ns := make([]int64, s.pos+10)
		copy(ns, s.s)
		s.s = ns
	}
	s.s[s.pos] = i
}

func (s *Stack) Pop() (int64, error) {
	if s.pos < 0 {
		return 0, ErrStackIsOut
	}
	v := s.s[s.pos]
	s.pos--
	return v, nil
}

func (s Stack) String() string {
	m := make([]string, s.pos+1)
	for i := 0; i <= s.pos; i++ {
		m[i] = fmt.Sprintf("%04x\t%d", i, s.s[s.pos-i])
	}
	return strings.Join(m, "\n")
}
