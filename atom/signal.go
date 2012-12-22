package atom

type Kind *string
func NewKind(name string) Kind {
	return &name
}

type Signal interface {
	Kind() Kind
}
