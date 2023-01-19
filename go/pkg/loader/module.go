package loader

import (
	"go.uber.org/fx"
)

var InitModule = fx.Provide(
	ProductLoaderProvider,
	EventSeriesLoaderProvider,
	EventLoaderProvider,
)

var Module = fx.Provide(
	LoadersProvider,
)
