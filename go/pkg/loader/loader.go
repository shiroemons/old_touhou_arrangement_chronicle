package loader

import (
	"github.com/graph-gophers/dataloader"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	ProductLoader             *ProductLoader
	EventSeriesLoaderProvider *EventSeriesLoader
	EventLoaderProvider       *EventLoader
}

type Loaders struct {
	pLoader  *dataloader.Loader
	esLoader *dataloader.Loader
	eLoader  *dataloader.Loader
}

func LoadersProvider(p Params) *Loaders {
	return &Loaders{
		pLoader:  dataloader.NewBatchedLoader(p.ProductLoader.BatchGetProducts),
		esLoader: dataloader.NewBatchedLoader(p.EventSeriesLoaderProvider.BatchGetEventSeries),
		eLoader:  dataloader.NewBatchedLoader(p.EventLoaderProvider.BatchGetEvent),
	}
}
