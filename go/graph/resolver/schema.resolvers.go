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

// EventSeries is the resolver for the eventSeries field.
func (r *eventResolver) EventSeries(ctx context.Context, obj *model.Event) (*model.EventSeries, error) {
	eventSeries, err := loader.LoadEventSeries(ctx, obj.EventSeries.ID)
	if err != nil {
		return nil, err
	}
	return eventSeries, nil
}

// CreateEventSeries is the resolver for the createEventSeries field.
func (r *mutationResolver) CreateEventSeries(ctx context.Context, input model.NewEventSeries) (*model.EventSeries, error) {
	series, err := r.esSrv.New(ctx, input)
	if err != nil {
		return nil, err
	}
	return series.ToGraphQL(), nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	event, err := r.eSrv.New(ctx, input)
	if err != nil {
		return nil, err
	}
	return event.ToGraphQL(), nil
}

// CreateSubEvent is the resolver for the createSubEvent field.
func (r *mutationResolver) CreateSubEvent(ctx context.Context, input model.NewSubEvent) (*model.SubEvent, error) {
	sub, err := r.seSrv.New(ctx, input)
	if err != nil {
		return nil, err
	}
	return sub.ToGraphQL(), nil
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

// Event is the resolver for the event field.
func (r *subEventResolver) Event(ctx context.Context, obj *model.SubEvent) (*model.Event, error) {
	event, err := loader.LoadEvent(ctx, obj.Event.ID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// OriginalSong returns generated.OriginalSongResolver implementation.
func (r *Resolver) OriginalSong() generated.OriginalSongResolver { return &originalSongResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// SubEvent returns generated.SubEventResolver implementation.
func (r *Resolver) SubEvent() generated.SubEventResolver { return &subEventResolver{r} }

type eventResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type originalSongResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subEventResolver struct{ *Resolver }
