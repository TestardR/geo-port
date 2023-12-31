package in_memory_store

import (
	"context"
	"sync"

	"github.com/TestardR/geo-port/internal/domain"
)

// portStore should have dedicated data models
// and not store domain models. Domain models and data models should evolve separately
type portStore struct {
	mx sync.Mutex
	db map[domain.PortID]domain.Port
}

func NewPortStore(
	db map[domain.PortID]domain.Port,
) *portStore {
	return &portStore{
		db: db,
	}
}

func (s *portStore) Find(ctx context.Context, id domain.PortID) (domain.Port, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	port, isFound := s.db[id]
	if !isFound {
		return domain.Port{}, domain.ErrPortNotFound
	}

	return port, nil
}

func (s *portStore) Add(ctx context.Context, port domain.Port) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	_, isFound := s.db[port.ID()]
	if isFound {
		return domain.ErrPortAlreadyExists
	}

	s.db[port.ID()] = port

	return nil
}

func (s *portStore) Update(ctx context.Context, port domain.Port) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	_, isFound := s.db[port.ID()]
	if !isFound {
		return domain.ErrPortNotFound
	}

	s.db[port.ID()] = port

	return nil
}
