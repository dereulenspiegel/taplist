package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/dereulenspiegel/go-brewchild"
	"github.com/dereulenspiegel/taplist/graph/generated"
	"github.com/dereulenspiegel/taplist/graph/model"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) UpdateTap(ctx context.Context, id string, data model.TapData) (*model.Tap, error) {
	logger := logrus.WithFields(logrus.Fields{
		"tapId": id,
	})
	tap, err := r.store.Tap(id)
	if err != nil {
		logger.WithError(err).Error("Failed to query tap to be updated")
		return nil, err
	}

	tap.Empty = data.Empty
	if data.Name != nil {
		tap.Name = *data.Name
	}
	if data.PercentAvailable != nil {
		tap.PercentAvailable = data.PercentAvailable
	}
	if data.VolumeAvailable != nil {
		tap.VolumeAvailable = data.VolumeAvailable
	}
	if data.Beer != nil {
		if tap.Beer == nil {
			tap.Beer = &model.Beer{}
		}
		tap.Beer.Abv = data.Beer.Abv
		tap.Beer.Name = data.Beer.Name
		if data.Beer.BuGuRatio != nil {
			tap.Beer.BuGuRatio = data.Beer.BuGuRatio
		}
		if data.Beer.ColorEbc != nil {
			tap.Beer.ColorEbc = data.Beer.ColorEbc
		}
		if data.Beer.Fg != nil {
			tap.Beer.Fg = data.Beer.Fg
		}
		if data.Beer.GravityUnit != nil {
			tap.Beer.GravityUnit = data.Beer.GravityUnit
		}
		if data.Beer.Ibu != nil {
			tap.Beer.Ibu = data.Beer.Ibu
		}
		if data.Beer.Og != nil {
			tap.Beer.Og = data.Beer.Og
		}
	}

	if err := r.store.UpdateTap(tap); err != nil {
		logger.WithError(err).Error("Failed to update tap")
		return nil, err
	}
	return tap, nil
}

func (r *mutationResolver) SetBrewfatherBatchOnTap(ctx context.Context, tapID string, brewfatherBatchID string) (*model.Tap, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Kegerator(ctx context.Context) (*model.Kegerator, error) {
	return r.store.Kegerator()
}

func (r *queryResolver) Taps(ctx context.Context) ([]*model.Tap, error) {
	return r.store.Taps()
}

func (r *queryResolver) BrewfatherBatches(ctx context.Context, state *string) ([]*model.BrewfatherBatch, error) {
	if r.brewfatherClient == nil {
		return nil, errors.New("Brewfather is unconfigured")
	}
	var err error
	var batches []*brewchild.Batch
	if state != nil && *state != "" {
		batches, err = r.brewfatherClient.Batches(brewchild.Status(*state), brewchild.Complete(true))
	} else {
		batches, err = r.brewfatherClient.Batches(brewchild.Complete(true))
	}
	if err != nil {
		return nil, err
	}
	tb := make([]*model.BrewfatherBatch, len(batches))
	for i, b := range batches {
		ebc := int(b.EstimatedColor)
		gravityUnit := "SG"
		tb[i] = &model.BrewfatherBatch{
			ID:    fmt.Sprintf("brewfather:%d", b.BatchNumber),
			State: &b.Status,
			Beer: &model.Beer{
				ID:          fmt.Sprintf("brewfather:%d", b.BatchNumber),
				Name:        b.Name,
				Abv:         b.MeasuredABV,
				BuGuRatio:   &b.BuGuRatio,
				Ibu:         &b.IBU,
				ColorEbc:    &ebc,
				Og:          &b.OG,
				Fg:          &b.FG,
				GravityUnit: &gravityUnit,
			},
		}

		if tb[i].Beer.Abv > 0.1 {
			tb[i].Beer.Abv = b.ABV
		}

	}

	return tb, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
