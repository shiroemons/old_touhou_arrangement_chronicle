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

type EventSeriesLoader struct {
	esRepo domain.EventSeriesRepository
}

func EventSeriesLoaderProvider(esRepo domain.EventSeriesRepository) *EventSeriesLoader {
	return &EventSeriesLoader{esRepo: esRepo}
}

func (es *EventSeriesLoader) BatchGetEventSeries(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	esIDs := make([]string, len(keys))
	for ix, key := range keys {
		esIDs[ix] = key.String()
	}

	eventSeriesByID, err := es.esRepo.GetMapInIDs(ctx, esIDs)
	if err != nil {
		err = fmt.Errorf("fail get event series, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index, eventSeriesKey := range keys {
		eventSeries, ok := eventSeriesByID[eventSeriesKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: eventSeries, Error: nil}
		} else {
			err = fmt.Errorf("event series not found %s", eventSeriesKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func LoadEventSeries(ctx context.Context, esID string) (*model.EventSeries, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.pLoader.Load(ctx, dataloader.StringKey(esID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	eventSeries := result.(*entity.EventSeries)
	return eventSeries.ToGraphQL(), nil
}
