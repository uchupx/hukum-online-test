package repo

import (
	"context"

	"github.com/uchupx/geek-garden-test/internal/repo/model"
)

type GatheringRepo interface {
	GetGatheringByID(ctx context.Context, id int64) (*model.Gathering, error)
	GetGatherings(ctx context.Context) ([]model.Gathering, error)
	InsertGathering(ctx context.Context, data model.Gathering) (*int64, error)
	UpdateGathering(ctx context.Context, data model.Gathering) (*int64, error)
	DeleteGathering(ctx context.Context, data model.Gathering) (*int64, error)
}
