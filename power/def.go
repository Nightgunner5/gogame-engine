package power

import "github.com/Nightgunner5/gogame-engine/tile"

type Conductivity struct {
	Dx, Dy     int
	AcceptType tile.Type
}

type Def interface {
	tile.Def

	Conductivity(tile.Set) []Conductivity
	PowerUsage(tile.Set) uint32
	PowerGenerated(tile.Set) uint32
}
