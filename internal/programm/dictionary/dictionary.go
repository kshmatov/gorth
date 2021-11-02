package dictionary

import (
	"github.com/kshmatov/gorth/internal/programm/op"
)

var ops = map[string]func() op.Op{
	"+":    op.Add,
	"-":    op.Sub,
	"*":    op.Mul,
	"/":    op.Div,
	".":    op.Dump,
	"add":  op.Add,
	"sub":  op.Sub,
	"mul":  op.Mul,
	"div":  op.Div,
	"dump": op.Dump,
	"eq":   op.Eq,
	"=":    op.Eq,
	"gt":   op.Gt,
	">":    op.Gt,
	"gte":  op.Gte,
	">=":   op.Gte,
	"lt":   op.Lt,
	"<":    op.Lt,
	"lte":  op.Lte,
	"<=":   op.Lte,
	"dub":  op.Dub,
}
