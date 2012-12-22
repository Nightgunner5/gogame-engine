package tile

type MultiTile []Tile
var _ TileDef = MultiTile{}

func (m MultiTile) Pass(ts TileSet) bool {
	for _, t := range m {
		if !t.Pass(ts) {
			return false
		}
	}
	return true
}

func (m MultiTile) See(ts TileSet) bool {
	for _, t := range m {
		if !t.See(ts) {
			return false
		}
	}
	return true
}

func (m MultiTile) Light(ts TileSet) uint8 {
	var light uint8
	for _, t := range m {
		l := t.Light(ts)
		if 255 - l > light {
			light = 255
		} else {
			light += l
		}
	}
	return light
}
