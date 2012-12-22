package tile

type Def interface {
	// Atoms can move through this tile
	Pass(Set) bool

	// Atoms can see over/through this tile
	See(Set) bool

	// Amount of light emitted by this tile
	Light(Set) uint8
}

type noopDef struct{}

func (noopDef) Pass(Set) bool {
	return true
}
func (noopDef) See(Set) bool {
	return true
}
func (noopDef) Light(Set) uint8 {
	return 0
}
