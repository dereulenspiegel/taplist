package graph

import "github.com/dereulenspiegel/taplist/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Store interface {
	Taps() ([]*model.Tap, error)
	Kegerator() (*model.Kegerator, error)
}

type Resolver struct {
	store Store
}

func NewResolver(str Store) *Resolver {
	return &Resolver{
		store: str,
	}
}
