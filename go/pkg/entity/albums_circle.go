package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type AlbumsCircles struct {
	bun.BaseModel `bun:"table:albums_circles,alias:alc"`

	AlbumID   string    `bun:"album_id,pk,nullzero,notnull"`
	CircleID  string    `bun:"circle_id,pk,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
