package loader

import (
	"context"
	"fmt"
	"log"

	"github.com/graph-gophers/dataloader"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/domain"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ProductLoader struct {
	pRepo domain.ProductRepository
}

func ProductLoaderProvider(pRepo domain.ProductRepository) *ProductLoader {
	return &ProductLoader{pRepo: pRepo}
}

func (p *ProductLoader) BatchGetProducts(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	productIDs := make([]string, len(keys))
	for ix, key := range keys {
		productIDs[ix] = key.String()
	}

	productByID, err := p.pRepo.GetMapInIDs(ctx, productIDs)
	if err != nil {
		err = fmt.Errorf("fail get products, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index, productKey := range keys {
		product, ok := productByID[productKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: product, Error: nil}
		} else {
			err = fmt.Errorf("product not found %s", productKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func LoadProduct(ctx context.Context, productID string) (*model.Product, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.pLoader.Load(ctx, dataloader.StringKey(productID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	product := result.(*entity.Product)
	return product.ToGraphQL(), nil
}
