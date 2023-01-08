package resolver

import (
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/service"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Params struct {
	fx.In

	Logger          *zap.Logger
	ProductSrv      *service.ProductService
	OriginalSongSrv *service.OriginalSongService
}

type Resolver struct {
	logger *zap.Logger
	pSrv   *service.ProductService
	osSrv  *service.OriginalSongService
}

// NewResolver Resolver Constructor
func NewResolver(p Params) *Resolver {
	return &Resolver{
		logger: p.Logger,
		pSrv:   p.ProductSrv,
		osSrv:  p.OriginalSongSrv,
	}
}

// Provider Fx Provider
func Provider(p Params) generated.Config {
	return generated.Config{
		Resolvers: NewResolver(p),
	}
}

var Module = fx.Provide(
	Provider,
)
