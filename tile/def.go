package tile

type Type interface {
	// If a.Equal(b) OR b.Equal(a), the Types are considered equivelent.
	// This allows fuzzy types to accept their subtypes.
	Equal(Type) bool
}

type Def interface {
	// Atoms can move through this tile
	Pass(Set) bool

	// Atoms can see over/through this tile
	See(Set) bool

	// Amount of light emitted by this tile
	Light(Set) uint8

	// Used in various packages to differentiate, for example, door tiles
	// from wall tiles.
	Type(Set) Type
}

type noopDef struct{}

var _ Def = noopDef{}

func (noopDef) Pass(Set) bool {
	return true
}
func (noopDef) See(Set) bool {
	return true
}
func (noopDef) Light(Set) uint8 {
	return 0
}
func (noopDef) Type(Set) Type {
	return noopType{}
}

type noopType struct{}

var _ Type = noopType{}

func (noopType) Equal(other Type) bool {
	_, ok := other.(noopType)
	return ok
}
