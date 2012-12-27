// Package atom provides the base for all objects in the Go Game Engine. Atoms
// are self-contained and communicate by sending signals to other atoms.
package atom

import "fmt"

// for testing
var kindTestSignal = NewKind("testSignal")

type testSignal chan Kind

func (testSignal) Kind() Kind {
	return kindTestSignal
}

type Atom interface {
	Send(Signal)
	HandleSignal(Signal)
	atom() *atom
}

type atom struct {
	signals chan Signal
}

func New() Atom {
	return new(atom)
}

func (a *atom) Send(s Signal) {
	if a.signals == nil {
		panic("atom: Send on uninitialized Atom is illegal.")
	}

	defer func() {
		recover()
	}()

	a.signals <- s
}

func (a *atom) HandleSignal(s Signal) {
	// for testing
	if t, ok := s.(testSignal); ok {
		t <- s.Kind()
		return
	}

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
	a.atom().signals = make(chan Signal)
	go a.atom().dispatch(a)
}

func Close(a Atom) {
	close(a.atom().signals)
}
