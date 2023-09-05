package domain

type PortID struct {
	value string
}

func NewPortID(value string) PortID {
	return PortID{value: value}
}

func (id PortID) Id() string {
	return id.value
}
