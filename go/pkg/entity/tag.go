package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Tag struct {
	bun.BaseModel `bun:"table:tag"`

	ID        string    `bun:",pk"`
	Name      string    `bun:"name,nullzero,notnull"`
	TagType   string    `bun:"tag_type,nullzero,notnull,default:'unknown'::tag_type"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
