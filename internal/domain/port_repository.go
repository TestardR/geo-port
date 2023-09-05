package domain

import (
	"context"
	"errors"
)

var (
	ErrPortNotFound      = errors.New("port not found")
	ErrPortAlreadyExists = errors.New("port already exists")
)

type PortFinder interface {
	Find(ctx context.Context, id PortID) (Port, error)
}

type PortAdder interface {
	Add(ctx context.Context, port Port) error
}

type PortUpdater interface {
	Update(ctx context.Context, port Port) error
}
