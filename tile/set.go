package tile

type Set map[Tile]Def

func (ts Set) Get(t Tile) Def {
	if d, ok := ts[t]; ok {
		return d
	}
	return noopDef{}
}
