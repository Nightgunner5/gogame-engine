package tile

// Each Tile has a game-specific meaning.
type Tile uint16

var _ Def = Tile(0)

func (t Tile) Pass(ts Set) bool {
	return ts.Get(t).Pass(ts)
}

func (t Tile) See(ts Set) bool {
	return ts.Get(t).See(ts)
}

func (t Tile) Light(ts Set) uint8 {
	return ts.Get(t).Light(ts)
}

func (t Tile) Type(ts Set) Type {
	return ts.Get(t).Type(ts)
}
