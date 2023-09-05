package domain

type coordinates struct {
	longitude float64
	latitude  float64
}

func NewCoordinates(
	longitude float64,
	latitude float64,
) coordinates {
	return coordinates{
		longitude: longitude,
		latitude:  latitude,
	}
}

func (c coordinates) Longitude() float64 {
	return c.longitude
}

func (c coordinates) Latitude() float64 {
	return c.latitude
}
