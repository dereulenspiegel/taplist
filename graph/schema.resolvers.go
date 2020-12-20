package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dereulenspiegel/taplist/graph/generated"
	"github.com/dereulenspiegel/taplist/graph/model"
)

func (r *mutationResolver) UpdateTap(ctx context.Context, id string, data model.TapData) (*model.Tap, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
