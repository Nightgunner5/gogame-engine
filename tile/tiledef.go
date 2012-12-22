package tile

type TileDef interface {
	// Atoms can move through this tile
	Pass(TileSet) bool

	// Atoms can see over/through this tile
	See(TileSet) bool

	// Amount of light emitted by this tile
	Light(TileSet) uint8
}

type noopTileDef struct {}
func (noopTileDef) Pass(TileSet) bool {
	return true
}
func (noopTileDef) See(TileSet) bool {
	return true
}
func (noopTileDef) Light(TileSet) uint8 {
	return 0
}
