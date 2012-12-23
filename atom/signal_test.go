package atom

import (
	"testing"
)

func BenchmarkSignalKindEquality(b *testing.B) {
	a, c := NewKind("foo"), NewKind("bar")
	d, e := a, NewKind("bar")

	for i := 0; i < b.N; i++ {
		_ = a == c
		_ = a == d
		_ = a == e
		_ = c == d
		_ = c == e
		_ = d == e
	}
}

func BenchmarkStringEquality(b *testing.B) {
	a, c := "foo", "bar"
	d, e := a, "bar"

	for i := 0; i < b.N; i++ {
		_ = a == c
		_ = a == d
		_ = a == e
		_ = c == d
		_ = c == e
		_ = d == e
	}
}
