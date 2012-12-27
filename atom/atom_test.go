package atom

import (
	"testing"
)

func expectPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Log("Caught panic (expected): ", r)
	} else {
		t.Error("Expected panic, but there was none.")
	}
}

func TestDoubleClose(t *testing.T) {
	a := New()
	Init(a)

	Close(a)

	defer expectPanic(t)
	Close(a)
}

func TestUninitializedClose(t *testing.T) {
	a := New()

	defer expectPanic(t)
	Close(a)
}

func TestUninitializedSend(t *testing.T) {
	a := New()

	defer expectPanic(t)
	a.Send(make(testSignal, 1)) // buffered so if it is actually recieved it doesn't cause a different panic
}

func TestInitializedSend(t *testing.T) {
	a := New()

	Init(a)
	defer Close(a)

	s := make(testSignal)
	a.Send(s)
	if k := <-s; k != kindTestSignal {
		t.Error("unexpected result from testSignal: expected ", kindTestSignal, ", got ", k)
	}
}
