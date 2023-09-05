package domain

// Coordinates Value Object. Be careful latitude should come first:
// https://support.google.com/maps/answer/18539
type coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinates(
	latitude float64,
	longitude float64,
) coordinates {
	return coordinates{
		latitude:  latitude,
		longitude: longitude,
	}
}

func (c coordinates) Latitude() float64 {
	return c.latitude
}

func (c coordinates) Longitude() float64 {
	return c.longitude
}
