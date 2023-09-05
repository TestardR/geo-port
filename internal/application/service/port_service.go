package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/TestardR/geo-port/internal/application/command"
	"github.com/TestardR/geo-port/internal/domain"
)

type portFindAddUpdater interface {
	domain.PortFinder
	domain.PortAdder
	domain.PortUpdater
}

type portValidator interface {
	Validate(port domain.Port) []error
}

type service struct {
	portStore     portFindAddUpdater
	portValidator portValidator
}

func NewService(
	portStore portFindAddUpdater,
	portValidator portValidator,
) *service {
	return &service{
		portStore:     portStore,
		portValidator: portValidator,
	}
}

func (s *service) HandleAddOrUpdatePort(
	ctx context.Context,
	command command.AddOrUpdatePort,
) error {
	port, err := s.portStore.Find(ctx, domain.NewPortID(command.PortID()))
	if err != nil {
		if errors.Is(err, domain.ErrPortNotFound) {
			// We could have created a scoped command to add ports for better segregation of concerns
			return s.handleAddPort(ctx, command)
		}

		return err
	}

	// We could have created a scoped command to update ports for better segregation of concerns
	return s.handleUpdatePort(ctx, command, port)
}

func (s *service) handleAddPort(
	ctx context.Context,
	command command.AddOrUpdatePort,
) error {
	portID := domain.NewPortID(command.PortID())
	port := domain.NewPort(
		portID,
		command.Name(),
		command.City(),
		command.Country(),
		command.Aliases(),
		command.Regions(),
		domain.NewCoordinates(command.Latitude(), command.Longitude()),
		command.Province(),
		command.Timezone(),
		command.Unlocs(),
		command.Code(),
	)

	errs := s.portValidator.Validate(port)
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	err := s.portStore.Add(ctx, port)
	if err != nil {
		return fmt.Errorf("failed to add port with id %s: %v", portID.Id(), err)
	}

	return nil
}

func (s *service) handleUpdatePort(
	ctx context.Context,
	command command.AddOrUpdatePort,
	port domain.Port,
) error {
	portID := domain.NewPortID(command.PortID())
	updatePortChange := port.UpdatePortChange(
		portID,
		command.Name(),
		command.City(),
		command.Country(),
		command.Aliases(),
		command.Regions(),
		domain.NewCoordinates(command.Latitude(), command.Longitude()),
		command.Province(),
		command.Timezone(),
		command.Unlocs(),
		command.Code(),
	)

	errs := s.portValidator.Validate(updatePortChange)
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	err := s.portStore.Update(ctx, updatePortChange)
	if err != nil {
		return fmt.Errorf("failed to update port with id %s: %v", portID.Id(), err)
	}

	return nil
}
