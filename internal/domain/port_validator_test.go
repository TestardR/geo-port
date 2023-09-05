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
			NewCoordinates(56.33, 25.12),
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
			"", // name should not be empty string
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{},
			[]string{},
			NewCoordinates(56.33, 0), // latitude should not be 0
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"",
		)

		portValidator := NewPortValidator()

		expectedErrors := []error{
			errors.New("name is required"),
			errors.New("coordinates are required"),
		}

		violations := portValidator.Validate(port)
		g.Expect(violations).To(gomega.Equal(expectedErrors))
	})
}
