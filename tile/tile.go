package tile

// Each Tile has a game-specific meaning.
type Tile uint16

var _ TileDef = Tile(0)

func (t Tile) Pass(ts TileSet) bool {
	return ts.Get(t).Pass(ts)
}

func (t Tile) See(ts TileSet) bool {
	return ts.Get(t).See(ts)
}

func (t Tile) Light(ts TileSet) uint8 {
	return ts.Get(t).Light(ts)
}
