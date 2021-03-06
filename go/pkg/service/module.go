package service

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	ProductServiceProvider,
	OriginalSongServiceProvider,
	EventSeriesServiceProvider,
	EventServiceProvider,
	SubEventServiceProvider,
	ArtistServiceProvider,
	CircleServiceProvider,
)
