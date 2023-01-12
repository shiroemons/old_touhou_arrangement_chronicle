package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type CircleDetail struct {
	bun.BaseModel `bun:"table:circle_details,alias:cd"`

	CircleID          string    `bun:"circle_id,pk,nullzero,notnull"`
	Description       string    `bun:"description,nullzero,notnull,default:''"`
	URL               string    `bun:"url,nullzero,notnull,default:''"`
	BlogURL           string    `bun:"blog_url,nullzero,notnull,default:''"`
	TwitterURL        string    `bun:"twitter_url,nullzero,notnull,default:''"`
	YoutubeChannelURL string    `bun:"youtube_channel_url,nullzero,notnull,default:''"`
	CreatedAt         time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt         time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
