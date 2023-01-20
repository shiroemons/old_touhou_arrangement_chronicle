package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/ctxkey"
	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/entity"
)

type ArtistRepository struct {
	db *bun.DB
}

func ArtistRepositoryProvider(db *bun.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) Create(ctx context.Context, artist *entity.Artist) error {
	tx, ok := ctx.Value(ctxkey.TxCtxKey).(*bun.Tx)
	if ok {
		if _, err := tx.NewInsert().Model(artist).Exec(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err := r.db.NewInsert().Model(artist).Exec(ctx); err != nil {
		return err
	}
	return nil
}
