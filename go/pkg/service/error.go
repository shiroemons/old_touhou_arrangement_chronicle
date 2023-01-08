package service

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func SrvErr(ctx context.Context, msg string) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: msg,
	}
}
