package tile

type TileDef interface {
	// Atoms can move through this tile
	Pass() bool

	// Atoms can see over/through this tile
	See() bool

	// Amount of light emitted by this tile
	Light() uint8
}

type noopTileDef struct {}
func (noopTileDef) Pass() bool {
	return true
}
func (noopTileDef) See() bool {
	return true
}
func (noopTileDef) Light() uint8 {
	return 0
}
