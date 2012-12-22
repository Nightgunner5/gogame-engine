package tile

import "image"

type Coord image.Point

func C(x, y int) Coord {
	return Coord{x, y}
}
