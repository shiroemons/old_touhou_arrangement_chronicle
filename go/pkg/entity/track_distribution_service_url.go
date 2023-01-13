package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TrackDistributionServiceURL struct {
	bun.BaseModel `bun:"table:track_distribution_service_urls,alias:tdsu"`

	ID        string    `bun:",pk"`
	TrackID   string    `bun:"track_id,nullzero,notnull"`
	Service   string    `bun:"service,nullzero,notnull"`
	URL       string    `bun:"url,nullzero,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
