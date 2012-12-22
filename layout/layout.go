package layout

import (
	"github.com/Nightgunner5/gogame-engine/tile"
)

type Layout struct {
	base    map[tile.Coord]tile.MultiTile
	changes map[tile.Coord]tile.MultiTile
}

func (l Layout) Set() {
	
}
