package loader

import (
	"go.uber.org/fx"
)

var InitModule = fx.Provide(
	ProductLoaderProvider,
)

var Module = fx.Provide(
	LoadersProvider,
)
