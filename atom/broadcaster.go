package atom

type Broadcaster interface {
}

type broadcaster struct {
}

func (b *broadcaster) initialize() {
}

func NewBroadcaster(b Broadcaster) Broadcaster {
	if b == nil {
		b_ := new(broadcaster)
		b_.initialize()
		b = b_
	}
	return b
}
