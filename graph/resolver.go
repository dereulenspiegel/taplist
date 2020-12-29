package graph

import (
	"github.com/dereulenspiegel/go-brewchild"
	"github.com/dereulenspiegel/taplist/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Store interface {
	Taps() ([]*model.Tap, error)
	Kegerator() (*model.Kegerator, error)
	Tap(id string) (*model.Tap, error)
	UpdateTap(tap *model.Tap) error
}

type Resolver struct {
	store            Store
	brewfatherClient *brewchild.Client
	serverVersion    string
}

func NewResolver(str Store, brewchildClient *brewchild.Client, serveVersion string) *Resolver {
	return &Resolver{
		store:            str,
		brewfatherClient: brewchildClient,
		serverVersion:    serveVersion,
	}
}
