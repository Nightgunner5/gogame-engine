package layout

import "image"

type Coord image.Point

// Shortcut for layout.Coord{x, y}
func C(x, y int) Coord {
	return Coord{x, y}
}
