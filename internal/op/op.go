package op

type Ops uint32

const (
	OP_PUSH   = Ops(0)
	OP_PLUS   = Ops(iota)
	OP_DUMP   = Ops(iota)
	COUNT_OPS = Ops(iota)
)

type Operation interface {
	Execute() error
}

type Program []Operation
