package visibility

import (
	"github.com/Nightgunner5/gogame-engine/layout"
	"sync"
)

type Visibility struct {
	l       *layout.Layout
	version uint64
	cache   map[[2]layout.Coord]bool
	sync.RWMutex
}

func NewVisibility(l *layout.Layout) *Visibility {
	return &Visibility{l: l}
}

func (v *Visibility) Visible(x, y layout.Coord) bool {
	v.RLock()
	if v.l.Version() != v.version {
		v.RUnlock()
		v.Lock()
		defer v.Unlock()

		if v.l.Version() == v.version {
			return v.Visible(x, y)
		}

		v.version = v.l.Version()
		v.cache = make(map[[2]layout.Coord]bool)
		return v.Visible(x, y)
	}

	if visible, ok := v.cache[[2]layout.Coord{x, y}]; ok {
		v.RUnlock()
		return visible
	}
	v.RUnlock()

	v.Lock()
	defer v.Unlock()
	if visible, ok := v.cache[[2]layout.Coord{x, y}]; ok {
		return visible
	}

	visible := v.compute(x, y)
	v.cache[[2]layout.Coord{x, y}] = visible

	return visible
}

func (v *Visibility) compute(x, y layout.Coord) bool {
	if x == y {
		return true
	}
	if !v.l.Get(x).See(v.l.TileSet) {
		return false
	}
	dx, dy := y.X-x.X, y.Y-x.Y
	if dx < 0 {
		if dy < 0 {
			if -dx < -dy {
				return v.compute(layout.C(x.X, x.Y-1), y)
			}
			return v.compute(layout.C(x.X-1, x.Y), y)
		}
		if -dx < dy {
			return v.compute(layout.C(x.X, x.Y+1), y)
		}
		return v.compute(layout.C(x.X-1, x.Y), y)
	}
	if dy < 0 {
		if dx < -dy {
			return v.compute(layout.C(x.X, x.Y-1), y)
		}
		return v.compute(layout.C(x.X+1, x.Y), y)
	}
	if dx < dy {
		return v.compute(layout.C(x.X, x.Y+1), y)
	}
	return v.compute(layout.C(x.X+1, x.Y), y)
}
