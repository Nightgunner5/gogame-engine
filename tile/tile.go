package tile

// Each Tile has a game-specific meaning.
type Tile uint16

type MultiTile []Tile
var _ TileDef = MultiTile{}
