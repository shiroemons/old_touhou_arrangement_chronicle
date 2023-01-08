package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ProductRepository struct {
	db *bun.DB
}

func ProductRepositoryProvider(db *bun.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) All(ctx context.Context) ([]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	err := r.db.NewSelect().Model(&products).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetMapInIDs(ctx context.Context, ids []string) (map[string]*entity.Product, error) {
	products := make([]*entity.Product, 0)
	err := r.db.NewSelect().Model(&products).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, err
	}

	productById := map[string]*entity.Product{}
	for _, product := range products {
		p := product
		productById[p.ID] = p
	}
	return productById, nil
}
