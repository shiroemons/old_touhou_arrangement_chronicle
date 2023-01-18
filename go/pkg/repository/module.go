package repository

import (
	"go.uber.org/fx"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
)

var Module = fx.Provide(
	fx.Annotate(ProductRepositoryProvider, fx.As(new(domain.ProductRepository))),
	fx.Annotate(OriginalSongRepositoryProvider, fx.As(new(domain.OriginalSongRepository))),
	fx.Annotate(EventSeriesRepositoryProvider, fx.As(new(domain.EventSeriesRepository))),
)
