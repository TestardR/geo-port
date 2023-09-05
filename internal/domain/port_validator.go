package domain

import "errors"

type portValidator struct{}

func NewPortValidator() portValidator {
	return portValidator{}
}

// Validate is made up of fields that could be required from ports.json.
// It is an example of how validation could be handled in this service.
func (v portValidator) Validate(port port) []error {
	var violations []error

	if port.Name() == "" {
		violations = append(violations, errors.New("name is required"))
	}

	if port.Coordinates().Longitude() == 0 || port.Coordinates().Latitude() == 0 {
		violations = append(violations, errors.New("coordinates are required"))
	}

	return violations
}
