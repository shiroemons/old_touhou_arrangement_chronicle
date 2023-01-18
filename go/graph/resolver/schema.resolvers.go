package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/generated"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/loader"
)

// CreateEventSeries is the resolver for the createEventSeries field.
func (r *mutationResolver) CreateEventSeries(ctx context.Context, input model.NewEventSeries) (*model.EventSeries, error) {
	series, err := r.esSrv.New(ctx, input)
	if err != nil {
		return nil, err
	}
	return series.ToGraphQL(), nil
}

// Product is the resolver for the product field.
func (r *originalSongResolver) Product(ctx context.Context, obj *model.OriginalSong) (*model.Product, error) {
	product, err := loader.LoadProduct(ctx, obj.Product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.pSrv.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return products.ToGraphQLs(), nil
}

// OriginalSongs is the resolver for the originalSongs field.
func (r *queryResolver) OriginalSongs(ctx context.Context) ([]*model.OriginalSong, error) {
	os, err := r.osSrv.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return os.ToGraphQLs(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// OriginalSong returns generated.OriginalSongResolver implementation.
func (r *Resolver) OriginalSong() generated.OriginalSongResolver { return &originalSongResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type originalSongResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
