package domain

import (
	"errors"
	"github.com/onsi/gomega"
	"testing"
)

func TestCanValidatePort(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	portValidator := NewPortValidator()

	tests := []struct {
		description        string
		port               Port
		expectedViolations []error
	}{
		{
			description: "with a valid port",
			port: NewPort(
				NewPortID("AEFJR"),
				"Al Fujayrah",
				"Al Fujayrah",
				"United Arab Emirates",
				[]string{},
				[]string{},
				NewCoordinates(25.12, 56.33),
				"Ajman",
				"Asia/Dubai",
				[]string{"AEFJR"},
				"",
			),
			expectedViolations: nil,
		},
		{
			description: "with an invalid port",
			port: NewPort(
				NewPortID("AEFJR"),
				"", // name should not be an empty string
				"", // city should not be an empty string
				"United Arab Emirates",
				[]string{},
				[]string{},
				NewCoordinates(25.12, 0),
				"Ajman",
				"Asia/Dubai",
				[]string{"AEFJR"},
				"",
			),
			expectedViolations: []error{
				errors.New("name is required"),
				errors.New("city is required"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actualViolations := portValidator.Validate(tc.port)
			g.Expect(actualViolations).To(gomega.Equal(tc.expectedViolations))
		})
	}
}
