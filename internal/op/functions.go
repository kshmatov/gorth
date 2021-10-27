package op

import (
	"fmt"
	"strconv"

	"github.com/kshmatov/gorth/internal/stack"
)

func Push(x int64) Op {
	return &operation{
		desc: strconv.FormatInt(x, 10),
		f: func(s *stack.Stack) (int, error) {
			s.Push(x)
			return 1, nil
		},
	}
}

func Dump() Op {
	return &operation{
		desc: "dump",
		f: func(s *stack.Stack) (int, error) {
			v, err := s.Pop()
			if err != nil {
				return 1, err
			}
			fmt.Printf("%v\n", v)
			return 1, nil
		},
	}
}

func Dub() Op {
	return &operation{
		desc: "dub",
		f: func(s *stack.Stack) (int, error) {
			v, err := s.Pop()
			if err != nil {
				return 1, err
			}
			s.Push(v)
			s.Push(v)
			return 1, nil
		},
	}
}
