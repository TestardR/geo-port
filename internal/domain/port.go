package domain

type Port struct {
	id          portID
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
	id portID,
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

func (p Port) Id() portID {
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
