package domain

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"
)

func TestCanValidatePort(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	t.Run("with a valid port", func(t *testing.T) {
		port := NewPort(
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
		)

		portValidator := NewPortValidator()

		violations := portValidator.Validate(port)
		g.Expect(violations).To(gomega.BeNil())
	})

	t.Run("with an invalid port", func(t *testing.T) {
		port := NewPort(
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
		)

		portValidator := NewPortValidator()

		expectedErrors := []error{
			errors.New("name is required"),
			errors.New("city is required"),
		}

		violations := portValidator.Validate(port)
		g.Expect(violations).To(gomega.Equal(expectedErrors))
	})
}
