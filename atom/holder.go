package atom

import (
	"sync"
)

type Holder interface {
	Hold(Atom) bool
	Drop(Atom) bool

	EachHeld(func(Atom))
}

type holder struct {
	held []Atom
	sync.Mutex
}

func (h *holder) initialize() {
}

func (h *holder) Hold(a Atom) bool {
	if a == nil {
		return false
	}

	h.Lock()
	for _, b := range h.held {
		if a == b {
			h.Unlock()
			return false
		}
	}
	h.held = append(h.held, a)
	h.Unlock()
	return true
}

func (h *holder) Drop(a Atom) bool {
	if a == nil {
		return false
	}

	h.Lock()
	for i, b := range h.held {
		if a == b {
			h.held = append(h.held[:i], h.held[i+1:]...)
			h.Unlock()
			return true
		}
	}
	h.Unlock()
	return false
}

func (h *holder) EachHeld(f func(Atom)) {
	h.Lock()
	for _, a := range h.held {
		f(a)
	}
	h.Unlock()
}

func NewHolder(h Holder) Holder {
	if h == nil {
		h_ := new(holder)
		h_.initialize()
		h = h_
	}
	return h
}
