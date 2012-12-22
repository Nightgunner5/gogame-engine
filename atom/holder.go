package atom

type Holder interface {
}

type holder struct {
}

func (h *holder) initialize() {
}

func NewHolder(h Holder) Holder {
	if h == nil {
		h_ := new(holder)
		h_.initialize()
		h = h_
	}
	return h
}
