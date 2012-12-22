package atom

import "fmt"

type Atom interface {
	Send(Signal)
	HandleSignal(Signal)
	atom() *atom
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

func (a *atom) HandleSignal(s Signal) {
	// atom doesn't have any signal types on its own.
	panic(fmt.Errorf("Unhandled Signal of type %T", s))
}

func (a *atom) atom() *atom {
	return a
}

func (a *atom) dispatch(top Atom) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("Panic in Atom of type %T: %v", top, r))
		}
	}()

	for s := range a.signals {
		top.HandleSignal(s)
	}
}

func Init(a Atom) {
	go a.atom().dispatch(a)
}
