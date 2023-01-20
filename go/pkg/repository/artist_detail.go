package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ArtistDetailRepository struct {
	db *bun.DB
}

func ArtistDetailRepositoryProvider(db *bun.DB) *ArtistDetailRepository {
	return &ArtistDetailRepository{db: db}
}

func (r *ArtistDetailRepository) Create(ctx context.Context, artistDetail *entity.ArtistDetail) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(artistDetail).Exec(ctx); err != nil {
			return err
		}
		return nil
	}
	if _, err := r.db.NewInsert().Model(artistDetail).Exec(ctx); err != nil {
		return err
	}
	return nil
}
