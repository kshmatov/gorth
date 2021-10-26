package op

import (
	"github.com/kshmatov/gorth/internal/stack"
)

func Eq() Op {
	return &operation{
		desc: "eq",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b == a, s)
			return 1, nil
		},
	}
}

func Neq() Op {
	return &operation{
		desc: "neq",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b != a, s)
			return 1, nil
		},
	}
}

func Gt() Op {
	return &operation{
		desc: "gt",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b > a, s)
			return 1, nil
		},
	}
}

func Gte() Op {
	return &operation{
		desc: "gte",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b >= a, s)
			return 1, nil
		},
	}
}

func Lt() Op {
	return &operation{
		desc: "lt",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b < a, s)
			return 1, nil
		},
	}
}

func Lte() Op {
	return &operation{
		desc: "lte",
		f: func(s *stack.Stack) (int, error) {
			a, b, err := get2(s)
			if err != nil {
				return 1, err
			}
			cond(b <= a, s)
			return 1, nil
		},
	}
}
