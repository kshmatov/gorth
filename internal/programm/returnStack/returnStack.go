package returnStack

type Adress struct {
	Subroutine string
	Line       int
	next       *Adress
}

type ReturnStack struct {
	s *Adress
}

func New() *ReturnStack {
	return &ReturnStack{
		s: nil,
	}
}

func (r *ReturnStack) Add(name string, line int) {
	s := &Adress{
		Subroutine: name,
		Line:       line,
		next:       r.s,
	}
	r.s = s
}

func (r *ReturnStack) Get() *Adress {
	if r.s == nil {
		return nil
	}
	s := r.s
	r.s = s.next
	return s
}
