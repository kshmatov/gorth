package op

import (
	"fmt"
	"strconv"

	"github.com/kshmatov/gorth/internal/stack"
	"github.com/pkg/errors"
)

func Add() Op {
	return &op{
		desc: "add",
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
		desc: "sub",
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
		desc: strconv.FormatInt(x, 10),
		f: func(s *stack.Stack) error {
			s.Push(x)
			return nil
		},
	}
}

func Dump() Op {
	return &op{
		desc: "dump",
		f: func(s *stack.Stack) error {
			v, err := s.Pop()
			if err != nil {
				return err
			}
			fmt.Printf("%v\n", v)
			return nil
		},
	}
}
