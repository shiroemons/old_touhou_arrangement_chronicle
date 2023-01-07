package server

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	// Gin App
	AppProvider,
	// Handlers
	AppHandlersProvider,
)

var InvokeModule = fx.Invoke(
	// Gin Middleware and Endpoints Invoker
	InvokeMiddleware,
	InvokeHandler,
)
