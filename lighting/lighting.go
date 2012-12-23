package lighting

import (
	"github.com/Nightgunner5/gogame-engine/layout"
	"github.com/Nightgunner5/gogame-engine/tile"
	"sync"
)

type Lighting struct {
	l       *layout.Layout
	bright  map[layout.Coord]uint8
	version uint64
	sync.Mutex
}

func (l *Lighting) At(c layout.Coord) uint8 {
	l.Lock()
	for l.version != l.l.Version() {
		l.bright = make(map[layout.Coord]uint8)
		l.version = l.l.ForAll(func(c layout.Coord, t tile.Multi) {
			if b := t.Light(l.l.TileSet); b != 0 {
				l.spread(c, b)
			}
		})
	}
	b := l.bright[c]
	l.Unlock()
	return b
}

func (l *Lighting) spread(c layout.Coord, b uint8) {
	if 255-l.bright[c] < b {
		l.bright[c] = 255
	} else {
		l.bright[c] += b
	}

	if !l.l.Get(c).See(l.l.TileSet) {
		return
	}

	const loss = 10
	if b < loss {
		return
	}
	b -= loss

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			l.spread(layout.C(c.X+i, c.Y+j), b)
		}
	}
}
