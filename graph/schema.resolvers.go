package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/situmorangbastian/grip/graph/generated"
	"github.com/situmorangbastian/grip/graph/model"
)

// Examples is the resolver for the examples field.
func (r *queryResolver) Examples(ctx context.Context) ([]*model.Example, error) {
	return []*model.Example{
		{
			ID: "id",
		},
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
