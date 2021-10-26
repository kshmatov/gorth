package op

import (
	"testing"

	"github.com/kshmatov/gorth/internal/stack"
)

func TestGet2(t *testing.T) {
	s := stack.NewStack()
	_, _, err := get2(s)
	if err == nil {
		t.Errorf("get2 on empty stack must return err")
	}
	s.Push(1)
	_, _, err = get2(s)
	if err == nil {
		t.Errorf("get2 on 1 element stack must return err")
	}

	s.Push(1)
	s.Push(2)
	a, b, err := get2(s)
	if err != nil {
		t.Errorf("get2 on 2 element stack must not return err")
	}
	if a != 2 {
		t.Errorf("a must be 2, %v", a)
	}
	if b != 1 {
		t.Errorf("b must be 1, %v", b)
	}
}

func TestCont(t *testing.T) {
	s := stack.NewStack()
	cond(true, s)
	v, err := s.Pop()
	if err != nil {
		t.Errorf("cond must push somethong to stack, %v", err)
	}
	if v != 1 {
		t.Errorf("true cond must push 1 to stack, %v", v)
	}

	cond(false, s)
	v, err = s.Pop()
	if err != nil {
		t.Errorf("cond must push somethong to stack, %v", err)
	}
	if v != 0 {
		t.Errorf("false cond must push 0 to stack, %v", v)
	}

}
