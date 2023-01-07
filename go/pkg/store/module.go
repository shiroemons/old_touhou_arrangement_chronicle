package store

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	// Postgres Database
	PostgresProvider,
)
