package op

import (
	"testing"

	"github.com/kshmatov/gorth/internal/programm/stack"
	"github.com/pkg/errors"
)

func TestPush(t *testing.T) {
	s := stack.NewStack()
	val := int64(-12345)
	op := Push(val)
	_, err := op.Exec(s)
	if err != nil {
		t.Errorf("Push should be succesfull, but %v", err)
	}
	v, err := s.Pop()
	if err != nil {
		t.Errorf("Stack.pop should be succesfull, but %v", err)
	}
	if v != val {
		t.Errorf("%v was pushed but got %v", val, v)
	}
}

func TestDub(t *testing.T) {
	s := stack.NewStack()
	op := Dub()
	_, err := op.Exec(s)
	if !errors.Is(err, stack.ErrStackIsOut) {
		t.Errorf("Empty stack can't dub, got %v", err)
	}

	val := int64(-9876)
	s.Push(val)
	_, err = op.Exec(s)
	if err != nil {
		t.Errorf("Stack has value and can dub, but %v", err)
	}
	v1, err := s.Pop()
	if err != nil {
		t.Errorf("Stack must have values, but %v", err)
	}
	if v1 != val {
		t.Errorf("%v was dubbed, got %v", val, v1)
	}
	v2, err := s.Pop()
	if err != nil {
		t.Errorf("Stack must have second value, but %v", err)
	}
	if v2 != val {
		t.Errorf("%v was dubbed, got %v", val, v2)
	}

}
