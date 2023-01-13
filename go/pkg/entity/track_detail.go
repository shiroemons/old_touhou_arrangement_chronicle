package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type TrackDetail struct {
	bun.BaseModel `bun:"table:track_details,alias:td"`

	TrackID             string    `bun:"track_id,pk,nullzero,notnull"`
	Length              int       `bun:"length,nullzero"`
	BPM                 int       `bun:"bpm,nullzero"`
	DisplayComposer     string    `bun:"display_composer,nullzero,notnull,default:''"`
	DisplayArranger     string    `bun:"display_arranger,nullzero,notnull,default:''"`
	DisplayRearranger   string    `bun:"display_rearranger,nullzero,notnull,default:''"`
	DisplayLyricist     string    `bun:"display_lyricist,nullzero,notnull,default:''"`
	DisplayVocalist     string    `bun:"display_vocalist,nullzero,notnull,default:''"`
	DisplayOriginalSong string    `bun:"display_original_song,nullzero,notnull,default:''"`
	CreatedAt           time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
