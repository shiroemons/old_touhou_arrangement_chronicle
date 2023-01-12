package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Artist struct {
	bun.BaseModel `bun:"table:artists,alias:a"`

	ID                  string    `bun:",pk"`
	Name                string    `bun:"name,nullzero,notnull"`
	InitialLetterType   string    `bun:"initial_letter_type,nullzero,notnull"`
	InitialLetterDetail string    `bun:"initial_letter_detail,nullzero,notnull,default:''"`
	CreatedAt           time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
