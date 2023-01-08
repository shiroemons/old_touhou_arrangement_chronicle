package service

import (
	"context"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ProductService struct {
	pRepo domain.ProductRepository
}

func ProductServiceProvider(pRepo domain.ProductRepository) *ProductService {
	return &ProductService{pRepo: pRepo}
}

func (srv *ProductService) GetAll(ctx context.Context) (entity.Products, error) {
	products, err := srv.pRepo.All(ctx)
	if err != nil {
		return make([]*entity.Product, 0), SrvErr(ctx, err.Error())
	}
	return products, nil
}
