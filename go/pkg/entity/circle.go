package entity

import (
	"context"
	"time"

	"github.com/rs/xid"
	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/graph/model"
)

type Circle struct {
	bun.BaseModel `bun:"table:circles,alias:c"`

	ID                  string    `bun:",pk"`
	Name                string    `bun:"name,nullzero,notnull"`
	InitialLetterType   string    `bun:"initial_letter_type,type:initial_letter_type,nullzero,notnull"`
	InitialLetterDetail string    `bun:"initial_letter_detail,nullzero,notnull,default:''"`
	CreatedAt           time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt           time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*Circle)(nil)

func (e *Circle) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if e.ID == "" {
			e.ID = xid.New().String()
		}
	}
	return nil
}

func (e *Circle) ToGraphQL() *model.Circle {
	return &model.Circle{
		ID:                  e.ID,
		Name:                e.Name,
		InitialLetterType:   e.InitialLetterType,
		InitialLetterDetail: e.InitialLetterDetail,
	}
}
