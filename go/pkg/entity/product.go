package entity

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID           string            `bun:",pk"`
	Name         string            `bun:"name,nullzero,notnull"`
	ShortName    string            `bun:"short_name,nullzero,notnull"`
	ProductType  model.ProductType `bun:"product_type,nullzero,notnull"`
	SeriesNumber float64           `bun:"series_number,nullzero,notnull"`
	CreatedAt    time.Time         `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt    time.Time         `bun:"updated_at,notnull,default:current_timestamp"`
}

// ToGraphQL Convert to GraphQL Schema
func (e *Product) ToGraphQL() *model.Product {
	return &model.Product{
		ID:           e.ID,
		Name:         e.Name,
		ShortName:    e.ShortName,
		ProductType:  e.ProductType,
		SeriesNumber: e.SeriesNumber,
	}
}

// Products Method Injection
type Products []*Product

func (arr Products) ToGraphQLs() []*model.Product {
	res := make([]*model.Product, len(arr))
	for i, p := range arr {
		res[i] = p.ToGraphQL()
	}
	return res
}
