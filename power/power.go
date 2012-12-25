package power

import (
	"github.com/Nightgunner5/gogame-engine/layout"
	"github.com/Nightgunner5/gogame-engine/tile"
	"sync"
)

type Power struct {
	l       *layout.Layout
	roots   []*graph
	powered map[node]bool
	version uint64
	lock    sync.RWMutex
}

func (p *Power) recomputeAll() {
	p.roots, p.powered = nil, nil
	tiles, v := p.l.InOrder()
	p.version = v

	for _, t := range tiles {
		for _, tt := range t.Multi {
			if d, ok := p.l.TileSet.Get(tt).(Def); ok {
				if d.PowerGenerated(p.l.TileSet) != 0 {
					p.construct(t.Coord, d.PowerGenerated(p.l.TileSet), d.PowerUsage(p.l.TileSet), tt)
				}
			}
		}
	}

	p.compute()
}

func (p *Power) get(x, y int, t tile.Tile) bool {
	return p.powered[node{layout.C(x, y), t}]
}

func (p *Power) Powered(x, y int, t tile.Tile) bool {
	p.lock.RLock()
	if p.version == p.l.Version() && p.powered != nil {
		result := p.get(x, y, t)
		p.lock.RUnlock()
		return result
	}
	p.lock.RUnlock()

	p.lock.Lock()
	if p.version == p.l.Version() && p.powered != nil {
		result := p.get(x, y, t)
		p.lock.Unlock()
		return result
	}

	p.recomputeAll()
	result := p.get(x, y, t)
	p.lock.Unlock()
	return result
}
