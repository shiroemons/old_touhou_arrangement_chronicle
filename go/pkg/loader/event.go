package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type EventLoader struct {
	eRepo domain.EventRepository
}

func EventLoaderProvider(eRepo domain.EventRepository) *EventLoader {
	return &EventLoader{eRepo: eRepo}
}

func (e *EventLoader) BatchGetEvent(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	eIDs := make([]string, len(keys))
	for ix, key := range keys {
		eIDs[ix] = key.String()
	}

	eventSeriesByID, err := e.eRepo.GetMapInIDs(ctx, eIDs)
	if err != nil {
		err = fmt.Errorf("fail get events, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index, eventKey := range keys {
		event, ok := eventSeriesByID[eventKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: event, Error: nil}
		} else {
			err = fmt.Errorf("event not found %s", eventKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func LoadEvent(ctx context.Context, eID string) (*model.Event, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.pLoader.Load(ctx, dataloader.StringKey(eID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	event := result.(*entity.Event)
	return event.ToGraphQL(), nil
}
