package op

import (
	"fmt"

	"github.com/kshmatov/gorth/internal/stack"
	"github.com/pkg/errors"
)

func Add() Op {
	return &op{
		f: func(s *stack.Stack) error {
			a, err := s.Pop()
			if err != nil {
				return errors.Wrap(err, "on first argument")
			}
			b, err := s.Pop()
			if err != nil {
				return errors.Wrap(err, "on second argument")
			}
			s.Push(a + b)
			return nil
		},
	}
}

func Sub() Op {
	return &op{
		f: func(s *stack.Stack) error {
			a, err := s.Pop()
			if err != nil {
				return errors.Wrap(err, "on first argument")
			}
			b, err := s.Pop()
			if err != nil {
				return errors.Wrap(err, "on second argument")
			}
			s.Push(b - a)
			return nil
		},
	}
}

func Push(x int64) Op {
	return &op{
		f: func(s *stack.Stack) error {
			s.Push(x)
			return nil
		},
	}
}

func Dump() Op {
	return &op{
		f: func(s *stack.Stack) error {
			v, err := s.Pop()
			if err != nil {
				return err
			}
			fmt.Printf("%v", v)
			return nil
		},
	}
}
