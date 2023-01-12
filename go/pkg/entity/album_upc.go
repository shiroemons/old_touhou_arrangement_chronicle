package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type AlbumUPC struct {
	bun.BaseModel `bun:"table:album_upcs,alias:alu"`

	ID        string    `bun:",pk"`
	AlbumID   string    `bun:"album_id,nullzero,notnull"`
	UPC       string    `bun:"upc,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
