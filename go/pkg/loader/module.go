package loader

import (
	"go.uber.org/fx"
)

var InitModule = fx.Provide(
	ProductLoaderProvider,
	EventSeriesLoaderProvider,
)

var Module = fx.Provide(
	LoadersProvider,
)
