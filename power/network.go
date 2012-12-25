package power

import (
	"github.com/Nightgunner5/gogame-engine/layout"
	"github.com/Nightgunner5/gogame-engine/tile"
)

type node struct {
	layout.Coord
	tile.Tile
}

type graph struct {
	node
	emit, used uint32
	neighbors  []*graph
}

func (p *Power) construct(root layout.Coord, emit, used uint32, t tile.Tile) {
	g := &graph{
		node: node{root, t},
		emit: emit,
		used: used,
	}
	p.roots = append(p.roots, g)

	p.visit(g, map[node]bool{g.node: true})
}

func (p *Power) visit(root *graph, visited map[node]bool) {
	for _, c := range p.l.TileSet.Get(root.Tile).(Def).Conductivity(p.l.TileSet) {
		p.next(root, visited, c.Dx, c.Dy, c.AcceptType)
	}
}

func (p *Power) next(root *graph, visited map[node]bool, dx, dy int, t tile.Type) {
	coord := layout.C(root.X+dx, root.Y+dy)
	tiles := p.l.Get(coord)
	for _, tt := range tiles {
		d, ok := p.l.TileSet.Get(tt).(Def)
		if !ok {
			continue
		}
		if t2 := tt.Type(p.l.TileSet); t.Equal(t2) || t2.Equal(t) {
			n := node{coord, tt}
			if visited[n] {
				return
			}
			visited[n] = true

			g := &graph{
				node: n,
				emit: d.PowerGenerated(p.l.TileSet),
				used: d.PowerUsage(p.l.TileSet),
			}
			root.neighbors = append(root.neighbors, g)
			p.visit(g, visited)
			return
		}
	}
}

func (p *Power) compute() {
	p.powered = make(map[node]bool)

	for _, root := range p.roots {
		p.walk(root, make(map[*graph]bool), root.emit)
	}

	// Feed it to the GC
	p.roots = nil
}

func (p *Power) walk(g *graph, visited map[*graph]bool, remaining uint32) {
	if visited[g] {
		return
	}
	visited[g] = true

	if !p.powered[g.node] && remaining >= g.used {
		p.powered[g.node] = true
		remaining -= g.used

		if remaining+g.emit < remaining {
			panic("overflow")
		}

		remaining += g.emit
	}

	for _, g2 := range g.neighbors {
		p.walk(g2, visited, remaining)
	}
}
