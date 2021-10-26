package op

import (
	"testing"

	"github.com/kshmatov/gorth/internal/stack"
)

func TestEq(t *testing.T) {
	s := stack.NewStack()
	eq := Eq()

	s.Push(10)
	s.Push(10)
	_, err := eq.Exec(s)
	if err != nil {
		t.Errorf("Eq error should be nil, %v", err)
	}
	i, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Eq should be 1, %v", i)
	}

	s.Push(10)
	s.Push(100)
	eq.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Eq should be 0, %v", i)
	}
}

func TestNeq(t *testing.T) {
	s := stack.NewStack()
	neq := Neq()

	s.Push(10)
	s.Push(10)
	_, err := neq.Exec(s)
	if err != nil {
		t.Errorf("Neq error should be nil, %v", err)
	}
	i, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Neq should be 0, %v", i)
	}

	s.Push(10)
	s.Push(100)
	neq.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Neq should be 1, %v", i)
	}
}

func TestLt(t *testing.T) {
	s := stack.NewStack()
	lt := Lt()

	s.Push(10)
	s.Push(100)
	_, err := lt.Exec(s)
	if err != nil {
		t.Errorf("Lt error should be nil, %v", err)
	}
	i, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Lt should be 1, %v", i)
	}

	s.Push(100)
	s.Push(10)
	lt.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Lt should be 0, %v", i)
	}

	s.Push(10)
	s.Push(10)
	lt.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Lt should be 0, %v", i)
	}
}

func TestGte(t *testing.T) {
	s := stack.NewStack()
	gte := Gte()

	s.Push(100)
	s.Push(10)
	_, err := gte.Exec(s)
	if err != nil {
		t.Errorf("Gte error should be nil, %v", err)
	}
	i, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Gte should be 1, %v", i)
	}

	s.Push(10)
	s.Push(100)
	gte.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Gte should be 0, %v", i)
	}

	s.Push(10)
	s.Push(10)
	gte.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Gte should be 1, %v", i)
	}
}

func TestLte(t *testing.T) {
	s := stack.NewStack()
	lte := Lte()

	s.Push(10)
	s.Push(100)
	_, err := lte.Exec(s)
	if err != nil {
		t.Errorf("Lte error should be nil, %v", err)
	}
	i, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Lte should be 1, %v", i)
	}

	s.Push(100)
	s.Push(10)
	lte.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 0 {
		t.Errorf("Lte should be 0, %v", i)
	}

	s.Push(10)
	s.Push(10)
	lte.Exec(s)
	i, err = s.Pop()
	if err != nil {
		t.Errorf("Pop error should be nil, %v", err)
	}
	if i != 1 {
		t.Errorf("Lte should be 0, %v", i)
	}
}
