package layout

import (
	"github.com/Nightgunner5/gogame-engine/tile"
	"sort"
	"sync"
)

type Layout struct {
	version uint64
	TileSet tile.Set
	base    map[Coord]tile.Multi
	changes map[Coord]tile.Multi
	sync.RWMutex
}

func NewLayout(ts tile.Set) *Layout {
	return &Layout{
		version: 1, // Start version at 1 to make uninitialized vs first revision easier to handle.
		TileSet: ts,
		base:    make(map[Coord]tile.Multi),
		changes: make(map[Coord]tile.Multi),
	}
}

func (l *Layout) Get(c Coord) tile.Multi {
	l.RLock()
	if t, ok := l.changes[c]; ok {
		l.RUnlock()
		return t
	}
	t := l.base[c]
	l.RUnlock()
	return t
}

func (l *Layout) Set(c Coord, t tile.Multi) (old tile.Multi) {
	for {
		old := l.Get(c)
		if l.Swap(c, old, t) {
			return old
		}
	}
	panic("unreachable")
}

func (l *Layout) Swap(c Coord, old, t tile.Multi) bool {
	l.Lock()
	if t1, ok := l.changes[c]; ok {
		if old.Equal(t1) {
			l.changes[c] = t
			l.version++
			l.Unlock()
			return true
		}
		l.Unlock()
		return false
	}
	t1 := l.base[c]
	if old.Equal(t1) {
		l.changes[c] = t
		l.version++
		l.Unlock()
		return true
	}
	l.Unlock()
	return false
}

type Tile struct {
	Coord
	tile.Multi
}

type sortTile []Tile

func (s sortTile) Len() int {
	return len(s)
}

func (s sortTile) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortTile) Less(i, j int) bool {
	return s[i].X < s[j].X || (s[i].X == s[j].X && s[i].Y < s[j].Y)
}

// Returns the current contents of this Layout in order of coordinate (first
// sorted by X ascending, then sub-sorted by Y ascending) and the current
// version of the Layout.
func (l *Layout) InOrder() ([]Tile, uint64) {
	var sorted sortTile

	l.RLock()

	for c, t := range l.base {
		if _, ok := l.changes[c]; !ok && len(t) != 0 {
			sorted = append(sorted, Tile{c, t})
		}
	}

	for c, t := range l.changes {
		if len(t) != 0 {
			sorted = append(sorted, Tile{c, t})
		}
	}

	v := l.version

	l.RUnlock()

	sort.Sort(sorted)

	return []Tile(sorted), v
}

// Calls f for each defined coordinate. If the Layout changes during iteration,
// this method stops iterating immediately and returns 0. If a non-zero value is
// returned, every tile has been visited and the return value is the current
// version of the Layout. The ordering of calls to f is not defined.
func (l *Layout) ForAll(f func(Coord, tile.Multi)) uint64 {
	l.RLock()
	v := l.version

	for c, t := range l.base {
		if _, ok := l.changes[c]; !ok && len(t) != 0 {
			if v != l.version {
				l.RUnlock()
				return 0
			}
			l.RUnlock()
			f(c, t)
			l.RLock()
		}
	}

	for c, t := range l.changes {
		if len(t) != 0 {
			if v != l.version {
				l.RUnlock()
				return 0
			}
			l.RUnlock()
			f(c, t)
			l.RLock()
		}
	}

	if v != l.version {
		l.RUnlock()
		return 0
	}
	l.RUnlock()
	return v
}

func (l *Layout) Version() uint64 {
	l.RLock()
	v := l.version
	l.RUnlock()
	return v
}
