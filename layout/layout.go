package layout

import (
	"github.com/Nightgunner5/gogame-engine/tile"
)

type Layout struct {
	base    map[Coord]tile.MultiTile
	changes map[Coord]tile.MultiTile
}

func (l Layout) Set() {

}
