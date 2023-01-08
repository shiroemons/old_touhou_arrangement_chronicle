package entity

import (
	"time"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
	"github.com/uptrace/bun"
)

type OriginalSong struct {
	bun.BaseModel `bun:"table:original_songs,alias:os"`

	ID          string    `bun:",pk"`
	ProductID   string    `bun:"product_id,nullzero,notnull"`
	Name        string    `bun:"name,nullzero,notnull"`
	Composer    string    `bun:"composer,nullzero,notnull,default:''"`
	Arranger    string    `bun:"arranger,nullzero,notnull,default:''"`
	TrackNumber int       `bun:"track_number,nullzero,notnull"`
	Original    bool      `bun:"is_original,notnull"`
	SourceID    string    `bun:"source_id,nullzero,notnull,default:''"`
	CreatedAt   time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

// ToGraphQL Convert to GraphQL Schema
func (e *OriginalSong) ToGraphQL() *model.OriginalSong {
	return &model.OriginalSong{
		ID:          e.ID,
		Product:     &model.Product{ID: e.ProductID},
		Name:        e.Name,
		Composer:    e.Composer,
		Arranger:    e.Arranger,
		TrackNumber: e.TrackNumber,
		Original:    e.Original,
		SourceID:    e.SourceID,
	}
}

// OriginalSongs Method Injection
type OriginalSongs []*OriginalSong

// ToGraphQLs Convert all to GraphQL Schema
func (arr OriginalSongs) ToGraphQLs() []*model.OriginalSong {
	res := make([]*model.OriginalSong, len(arr))
	for i, os := range arr {
		res[i] = os.ToGraphQL()
	}
	return res
}
