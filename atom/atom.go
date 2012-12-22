package atom

type Atom interface {
	Send(Signal)
}

type atom struct {
	signals chan Signal
}

func (a *atom) initialize() {
	a.signals = make(chan Signal)
}

func NewAtom(a Atom) Atom {
	if a == nil {
		a_ := new(atom)
		a_.initialize()
		a = a_
	}
	return a
}

func (a *atom) Send(s Signal) {
	defer func() {
		recover()
	}()

	a.signals <- s
}
