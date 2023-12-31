package domain

type PortID struct {
	value string
}

func NewPortID(value string) PortID {
	return PortID{value: value}
}

func (id PortID) ID() string {
	return id.value
}
