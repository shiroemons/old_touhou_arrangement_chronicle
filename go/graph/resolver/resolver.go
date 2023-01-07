package resolver

import (
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

	Logger *zap.Logger
}

type Resolver struct {
	logger *zap.Logger
}

// NewResolver Resolver Constructor
func NewResolver(p Params) *Resolver {
	return &Resolver{
		logger: p.Logger,
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
