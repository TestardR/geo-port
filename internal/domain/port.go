package domain

type Port struct {
	id          PortID
	name        string
	city        string
	country     string
	aliases     []string
	regions     []string
	coordinates coordinates
	province    string
	timezone    string
	unlocs      []string
	code        string
}

func NewPort(
	id PortID,
	name string,
	city string,
	country string,
	aliases []string,
	regions []string,
	coordinates coordinates,
	province string,
	timezone string,
	unlocs []string,
	code string,
) Port {
	return Port{
		id:          id,
		name:        name,
		city:        city,
		country:     country,
		aliases:     aliases,
		regions:     regions,
		coordinates: coordinates,
		province:    province,
		timezone:    timezone,
		unlocs:      unlocs,
		code:        code,
	}
}

func (p Port) Id() PortID {
	return p.id
}

func (p Port) Name() string {
	return p.name
}

func (p Port) City() string {
	return p.city
}

func (p Port) Country() string {
	return p.country
}

func (p Port) Aliases() []string {
	return p.aliases
}

func (p Port) Regions() []string {
	return p.regions
}

func (p Port) Coordinates() coordinates {
	return p.coordinates
}

func (p Port) Province() string {
	return p.province
}

func (p Port) Timezone() string {
	return p.timezone
}

func (p Port) Unlocs() []string {
	return p.unlocs
}

func (p Port) Code() string {
	return p.code
}

// UpdatePortChange encapsulate the business logic to update a port
// Method could receive dedicated validator to update port.
// It could return an updatePortChange struct which would better reflect domain intention.
func (p Port) UpdatePortChange(
	id PortID,
	name string,
	city string,
	country string,
	aliases []string,
	regions []string,
	coordinates coordinates,
	province string,
	timezone string,
	unlocs []string,
	code string,
) Port {
	return Port{
		id:          id,
		name:        name,
		city:        city,
		country:     country,
		aliases:     aliases,
		regions:     regions,
		coordinates: coordinates,
		province:    province,
		timezone:    timezone,
		unlocs:      unlocs,
		code:        code,
	}
}
