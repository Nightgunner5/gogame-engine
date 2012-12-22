package tile

type TileSet map[Tile]TileDef

func (ts TileSet) Get(t Tile) TileDef {
	if d, ok := ts[t]; ok {
		return d
	}
	return noopTileDef{}
}
