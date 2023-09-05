package domain

import "errors"

type portValidator struct{}

func NewPortValidator() portValidator {
	return portValidator{}
}

// Validate is made up of fields that could be required from ports.json.
// It is an example of how validation could be handled in this service.
func (v portValidator) Validate(port Port) []error {
	var violations []error

	if port.Name() == "" {
		violations = append(violations, errors.New("name is required"))
	}

	if port.City() == "" {
		violations = append(violations, errors.New("city is required"))
	}
	return violations
}
