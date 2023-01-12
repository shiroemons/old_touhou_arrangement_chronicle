package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type AlbumDetail struct {
	bun.BaseModel `bun:"table:album_details,alias:ald"`

	AlbumID      string              `bun:"album_id,pk,nullzero,notnull"`
	AlbumNumber  string              `bun:"album_number,nullzero,notnull,default:''"`
	EventPrice   decimal.NullDecimal `bun:"event_price"`
	Currency     string              `bun:"currency,nullzero,notnull,default:'JPY'"`
	Credit       string              `bun:"credit,nullzero,notnull,default:''"`
	Introduction string              `bun:"introduction,nullzero,notnull,default:''"`
	URL          string              `bun:"url,nullzero,notnull,default:''"`
	CreatedAt    time.Time           `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt    time.Time           `bun:"updated_at,notnull,default:current_timestamp"`
}
