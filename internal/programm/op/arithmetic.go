package op

import (
	"github.com/kshmatov/gorth/internal/programm/stack"
)

func Add() Op {
	return &operation{
		desc: "add",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			s.Push(a + b)
			return 1, nil
		},
	}
}

func Sub() Op {
	return &operation{
		desc: "sub",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			s.Push(b - a)
			return 1, nil
		},
	}
}

func Mul() Op {
	return &operation{
		desc: "mul",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			s.Push(a * b)
			return 1, nil
		},
	}
}

func Div() Op {
	return &operation{
		desc: "div",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			s.Push(b / a)
			s.Push(b % a)
			return 1, nil
		},
	}
}
