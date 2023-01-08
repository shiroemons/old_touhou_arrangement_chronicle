package loader

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

func Middleware(loaders *Loaders) gin.HandlerFunc {
	loaders.pLoader.ClearAll()
	// return a middleware that injects the loader to the request context
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
