package atom

import (
	"testing"
)

func TestHolder(t *testing.T) {
	toAdd := NewAtom(nil)
	// No init as toAdd is never sent any signals

	h := NewHolder(nil)

	count := 0
	h.EachHeld(func(a Atom) {
		t.Errorf("unexpected held atom: %T %v", a, a)
		count++
	})
	if count != 0 {
		t.Errorf("empty Holder had %d held atoms", count)
	}

	count = 0
	if h.Hold(nil) {
		t.Error("nil was held")
	}
	h.EachHeld(func(a Atom) {
		t.Errorf("unexpected held atom: %T %v", a, a)
		count++
	})
	if count != 0 {
		t.Errorf("empty Holder (+nil) had %d held atoms", count)
	}

	count = 0
	if !h.Hold(toAdd) {
		t.Error("toAdd was not held")
	}
	if h.Hold(toAdd) {
		t.Error("toAdd was held twice")
	}
	h.EachHeld(func(a Atom) {
		if a != toAdd {
			t.Errorf("unexpected held atom: %T %v", a, a)
		}
		count++
	})
	if count != 1 {
		t.Errorf("Holder (one atom) had %d held atoms", count)
	}

	count = 0
	if !h.Drop(toAdd) {
		t.Error("toAdd was not dropped")
	}
	if h.Drop(nil) {
		t.Error("nil was dropped")
	}
	if h.Drop(toAdd) {
		t.Error("toAdd was dropped twice")
	}
	h.EachHeld(func(a Atom) {
		t.Errorf("unexpected held atom: %T %v", a, a)
		count++
	})
	if count != 0 {
		t.Errorf("empty Holder (recent drop) had %d held atoms", count)
	}
}
