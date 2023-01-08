package loader

import (
	"github.com/graph-gophers/dataloader"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	ProductLoader *ProductLoader
}

type Loaders struct {
	pLoader *dataloader.Loader
}

func LoadersProvider(p Params) *Loaders {
	return &Loaders{
		pLoader: dataloader.NewBatchedLoader(p.ProductLoader.BatchGetProducts),
	}
}
