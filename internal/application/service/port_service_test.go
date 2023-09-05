package service

import (
	"context"
	"github.com/TestardR/geo-port/internal/application/command"
	"github.com/TestardR/geo-port/internal/domain"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockedPortStore struct {
	mock.Mock
}

func (s *mockedPortStore) Find(ctx context.Context, id domain.PortID) (domain.Port, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(domain.Port), args.Error(1)
}

func (s *mockedPortStore) Add(ctx context.Context, port domain.Port) error {
	args := s.Called(ctx, port)
	return args.Error(0)
}

func (s *mockedPortStore) Update(ctx context.Context, port domain.Port) error {
	args := s.Called(ctx, port)
	return args.Error(0)
}

type mockedPortValidator struct {
	mock.Mock
}

func (v *mockedPortValidator) Validate(port domain.Port) []error {
	args := v.Called(port)
	return args.Get(0).([]error)
}

func TestPortService(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	portStore := &mockedPortStore{}
	portValidator := &mockedPortValidator{}

	portService := NewPortService(
		portStore,
		portValidator,
	)

	t.Run("can add new port if port does not already exist", func(t *testing.T) {
		ctx := context.Background()

		port := domain.NewPort(
			domain.NewPortID("AEFJR"),
			"Al Fujayrah",
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{},
			[]string{},
			domain.NewCoordinates(25.12, 56.33),
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"",
		)

		portStore.On("Find", ctx, domain.NewPortID("AEFJR")).Return(
			domain.Port{},
			domain.ErrPortNotFound,
		)
		portStore.On("Add", ctx, port).Return(nil)

		portValidator.On("Validate", port).Return([]error{})

		addOrUpdatePortCommand := command.NewAddOrUpdatePort(
			"AEFJR",
			"Al Fujayrah",
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{},
			[]string{},
			25.12,
			56.33,
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"",
		)

		err := portService.HandleAddOrUpdatePort(ctx, addOrUpdatePortCommand)
		g.Expect(err).To(gomega.BeNil())
	})

	t.Run("can update an existing port", func(t *testing.T) {
		ctx := context.Background()

		existingPort := domain.NewPort(
			domain.NewPortID("AEFJR-1"),
			"Al Fujayrah",
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{},
			[]string{},
			domain.NewCoordinates(25.12, 56.33),
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"",
		)

		updatedPort := domain.NewPort(
			domain.NewPortID("AEFJR-1"),
			"Al Fujayrah",
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{"test-alias-1"},
			[]string{"test-region-1"},
			domain.NewCoordinates(25.12, 56.33),
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"52000",
		)

		portStore.On("Find", ctx, domain.NewPortID("AEFJR-1")).Return(
			existingPort,
			nil,
		)
		portStore.On("Update", ctx, updatedPort).Return(nil)

		portValidator.On("Validate", updatedPort).Return([]error{})

		addOrUpdatePortCommand := command.NewAddOrUpdatePort(
			"AEFJR-1",
			"Al Fujayrah",
			"Al Fujayrah",
			"United Arab Emirates",
			[]string{"test-alias-1"},
			[]string{"test-region-1"},
			25.12,
			56.33,
			"Ajman",
			"Asia/Dubai",
			[]string{"AEFJR"},
			"52000",
		)

		err := portService.HandleAddOrUpdatePort(ctx, addOrUpdatePortCommand)
		g.Expect(err).To(gomega.BeNil())
	})

	t.Skip("TODO: handle error/unhappy paths")
}
