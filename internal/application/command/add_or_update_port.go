package command

type AddOrUpdatePort struct {
	portID    string
	name      string
	city      string
	country   string
	aliases   []string
	regions   []string
	latitude  float64
	longitude float64
	province  string
	timezone  string
	unlocs    []string
	code      string
}

func NewAddOrUpdatePort(
	portID string,
	name string,
	city string,
	country string,
	aliases []string,
	regions []string,
	latitude float64,
	longitude float64,
	province string,
	timezone string,
	unlocs []string,
	code string,
) AddOrUpdatePort {
	return AddOrUpdatePort{
		portID:    portID,
		name:      name,
		city:      city,
		country:   country,
		aliases:   aliases,
		regions:   regions,
		latitude:  latitude,
		longitude: longitude,
		province:  province,
		timezone:  timezone,
		unlocs:    unlocs,
		code:      code,
	}
}

func (c AddOrUpdatePort) PortID() string {
	return c.portID
}

func (c AddOrUpdatePort) Name() string {
	return c.name
}

func (c AddOrUpdatePort) City() string {
	return c.city
}

func (c AddOrUpdatePort) Country() string {
	return c.country
}

func (c AddOrUpdatePort) Aliases() []string {
	return c.aliases
}

func (c AddOrUpdatePort) Regions() []string {
	return c.regions
}

func (c AddOrUpdatePort) Latitude() float64 {
	return c.latitude
}

func (c AddOrUpdatePort) Longitude() float64 {
	return c.longitude
}

func (c AddOrUpdatePort) Province() string {
	return c.province
}

func (c AddOrUpdatePort) Timezone() string {
	return c.timezone
}

func (c AddOrUpdatePort) Unlocs() []string {
	return c.unlocs
}

func (c AddOrUpdatePort) Code() string {
	return c.code
}
