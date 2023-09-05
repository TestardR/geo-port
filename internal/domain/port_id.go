package domain

type portID struct {
	value string
}

func NewPortID(value string) portID {
	return portID{value: value}
}

func (id portID) Id() string {
	return id.value
}
