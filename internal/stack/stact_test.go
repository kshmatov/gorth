package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := NewStack()
	if s.pos != -1 {
		t.Errorf("Position should be -1, but %v", s.pos)
	}

	for i := 0; i < 10; i++ {
		v := int64(i + 10)
		s.Push(v)
		if s.pos != i {
			t.Errorf("Stack pos should be %v, but %v", i, s.pos)
		}
		if s.s[s.pos] != v {
			t.Errorf("Stack[%v] should be %v, but %v", i, v, s.s[s.pos])
		}
	}

	s.Push(100)
	if len(s.s) != 20 {
		t.Errorf("Stack should grouw, but it is %v", len(s.s))
	}
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	if err != ErrStackIsOut {
		t.Errorf("Shoud return %#v error, return %#v", ErrStackIsOut, err)
	}
	s.Push(1)
	v, err := s.Pop()
	if v != 1 {
		t.Errorf("Shoud return %v, return %v", 1, err)
	}
	if err != nil {
		t.Errorf("Error should be nil, but %#v", err)
	}
	if s.pos != -1 {
		t.Errorf("Position shoud be -1, but %v", s.pos)
	}

	v, err = s.Pop()
	if err != ErrStackIsOut {
		t.Errorf("Shoud return %#v error, return %#v", ErrStackIsOut, err)
	}
	if v != 0 {
		t.Errorf("On empty stack value should be 0, returns %v", v)
	}
	if s.pos != -1 {
		t.Errorf("Empty position shoud be -1, but %v", s.pos)
	}
}
