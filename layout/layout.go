package layout

import (
	"github.com/Nightgunner5/gogame-engine/tile"
	"sort"
	"sync"
)

type Layout struct {
	tset    tile.TileSet
	base    map[Coord]tile.MultiTile
	changes map[Coord]tile.MultiTile
	sync.RWMutex
}

func NewLayout(ts tile.TileSet) *Layout {
	return &Layout{
		tset:    ts,
		base:    make(map[Coord]tile.MultiTile),
		changes: make(map[Coord]tile.MultiTile),
	}
}

func (l *Layout) Get(c Coord) tile.MultiTile {
	l.RLock()
	if t, ok := l.changes[c]; ok {
		l.RUnlock()
		return t
	}
	t := l.base[c]
	l.RUnlock()
	return t
}

func (l *Layout) Set(c Coord, t tile.MultiTile) (old tile.MultiTile) {
	for {
		old := l.Get(c)
		if l.Swap(c, old, t) {
			return old
		}
	}
	panic("unreachable")
}

func (l *Layout) Swap(c Coord, old, t tile.MultiTile) bool {
	l.Lock()
	if t1, ok := l.changes[c]; ok {
		if old.Equal(t1) {
			l.changes[c] = t
			l.Unlock()
			return true
		}
		l.Unlock()
		return false
	}
	t1 := l.base[c]
	if old.Equal(t1) {
		l.changes[c] = t
		l.Unlock()
		return true
	}
	l.Unlock()
	return false
}

type Tile struct {
	Coord
	tile.MultiTile
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

func (l *Layout) InOrder() []Tile {
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

	l.RUnlock()

	sort.Sort(sorted)

	return []Tile(sorted)
}

func (l *Layout) ForAll(f func(Coord, tile.MultiTile)) {
	l.RLock()

	for c, t := range l.base {
		if _, ok := l.changes[c]; !ok && len(t) != 0 {
			l.RUnlock()
			f(c, t)
			l.RLock()
		}
	}

	for c, t := range l.changes {
		if len(t) != 0 {
			l.RUnlock()
			f(c, t)
			l.RLock()
		}
	}

	l.RUnlock()
}
