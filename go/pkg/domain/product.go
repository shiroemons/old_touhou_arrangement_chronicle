package domain

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ProductService interface {
	GetAll(ctx context.Context) (entity.Products, error)
}

type ProductRepository interface {
	All(ctx context.Context) ([]*entity.Product, error)
	GetMapInIDs(ctx context.Context, ids []string) (map[string]*entity.Product, error)
}
