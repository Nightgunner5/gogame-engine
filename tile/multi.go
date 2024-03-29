package tile

type Multi []Tile

var _ Def = Multi(nil)

func (m Multi) Pass(ts Set) bool {
	for _, t := range m {
		if !t.Pass(ts) {
			return false
		}
	}
	return true
}

func (m Multi) See(ts Set) bool {
	for _, t := range m {
		if !t.See(ts) {
			return false
		}
	}
	return true
}

func (m Multi) Light(ts Set) uint8 {
	var light uint8
	for _, t := range m {
		l := t.Light(ts)
		if 255-l > light {
			light = 255
		} else {
			light += l
		}
	}
	return light
}

// Returns the Type of the Tile with the highest index in this slice.
func (m Multi) Type(ts Set) Type {
	return m[len(m)-1].Type(ts)
}

func (m Multi) Equal(other Multi) bool {
	if len(m) != len(other) {
		return false
	}
	for i := range m {
		if m[i] != other[i] {
			return false
		}
	}
	return true
}
