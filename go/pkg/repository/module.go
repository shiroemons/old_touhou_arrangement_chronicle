package repository

import (
	"go.uber.org/fx"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
)

var Module = fx.Provide(
	fx.Annotate(TxRepositoryProvider, fx.As(new(domain.TxRepository))),
	fx.Annotate(ProductRepositoryProvider, fx.As(new(domain.ProductRepository))),
	fx.Annotate(OriginalSongRepositoryProvider, fx.As(new(domain.OriginalSongRepository))),
	fx.Annotate(EventSeriesRepositoryProvider, fx.As(new(domain.EventSeriesRepository))),
	fx.Annotate(EventRepositoryProvider, fx.As(new(domain.EventRepository))),
	fx.Annotate(EventDetailRepositoryProvider, fx.As(new(domain.EventDetailRepository))),
)
