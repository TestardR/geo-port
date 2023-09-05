package domain

type port struct {
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
) port {
	return port{
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

func (p port) Id() portID {
	return p.id
}

func (p port) Name() string {
	return p.name
}

func (p port) City() string {
	return p.city
}

func (p port) Country() string {
	return p.country
}

func (p port) Aliases() []string {
	return p.aliases
}

func (p port) Regions() []string {
	return p.regions
}

func (p port) Coordinates() coordinates {
	return p.coordinates
}

func (p port) Province() string {
	return p.province
}

func (p port) Timezone() string {
	return p.timezone
}

func (p port) Unlocs() []string {
	return p.unlocs
}

func (p port) Code() string {
	return p.code
}
