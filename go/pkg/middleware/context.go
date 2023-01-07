package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GinContextToContextMiddleware Gin Context Middleware
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GinContextFromContext Function to extract context from middleware
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(contextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
